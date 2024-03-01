package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"sort"
)

var configFilePath = path.Join(homeDir, ".config/go-launch/mru.json")

type MruEntry struct {
	Id    string `json:"id"`
	Count int    `json:"count"`
}

func initMru() {
	_, err := os.Open(configFilePath)
	if err != nil {
		os.WriteFile(configFilePath, []byte("[]"), 0644)
	}
}

func mapToDesktopEntry(mruEntry MruEntry) (*Entry, error) {
	matchingEntry, err := find(desktopEntries, func(entry *Entry) bool {
		return entry.Id == mruEntry.Id
	})
	if err != nil {
		return nil, err
	}
	return matchingEntry, nil
}

func getMruEntries() []MruEntry {
	file, err := os.ReadFile(configFilePath)
	if err != nil {
		return []MruEntry{}
	}
	var mru []MruEntry
	json.Unmarshal(file, &mru)
	sort.Slice(mru, func(i, j int) bool {
		return mru[i].Count > mru[j].Count
	})
	return mru
}

func updateMruEntry(desktopEntry *Entry) {
	mruEntries := getMruEntries()
	_, err := find(mruEntries, func(entry MruEntry) bool {
		return entry.Id == desktopEntry.Id
	})
	if err != nil {
		mruEntries = append(mruEntries, MruEntry{Id: desktopEntry.Id, Count: 1})
	} else {
		mruEntries = mapArray(mruEntries, func(entry MruEntry) MruEntry {
			if entry.Id == desktopEntry.Id {
				return MruEntry{Id: desktopEntry.Id, Count: entry.Count + 1}
			}
			return entry
		})
	}
	bytes, _ := json.Marshal(mruEntries)
	err = os.WriteFile(configFilePath, bytes, 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		updateMruEntries()
	}
}

func updateMruEntries() {
	mruEntries := getMruEntries()
	mapped := mapArray(mruEntries, func(entry MruEntry) *Entry {
		mapped, err := mapToDesktopEntry(entry)
		if err != nil {
			return nil
		}
		return mapped
	})
	mapped = filter(mapped, func(e *Entry) bool {
		return e != nil
	})
	i := 0
	for len(mapped) < COUNT {
		if desktopEntries[i] != nil && desktopEntries[i].Id != "" {
			mapped = append(mapped, desktopEntries[i])
		}
		i = i + 1
	}
	mruDesktopEntries = mapped
}
