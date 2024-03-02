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
		print("mru config file not found")
		return []MruEntry{}
	}
	var mru []MruEntry
	json.Unmarshal(file, &mru)
	sort.Slice(mru, func(i, j int) bool {
		return mru[i].Count > mru[j].Count
	})
	print("found " + fmt.Sprint(len(mru)) + " mru entries")
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
		mruDesktopEntries = getMruDesktopEntries()
	}
}

func getMruDesktopEntries() []*Entry {
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
	if len(mapped) < COUNT {
		print("found only " + fmt.Sprint(len(mapped)) + " mru entries ; " + "filling up " + fmt.Sprint(COUNT-len(mapped)) + " mru entries")
		return fillUpDesktopEntries(mapped)
	}
	return mapped
}
