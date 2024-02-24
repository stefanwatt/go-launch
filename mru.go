package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"sort"

	"code.rocketnine.space/tslocum/desktop"
)

var configFilePath = path.Join(homeDir, ".config/go-launch/mru.json")

type MruEntry struct {
	Exec  string `json:"exec"`
	Count int    `json:"count"`
}

func initMru() {
	_, err := os.Open(configFilePath)
	if err != nil {
		os.WriteFile(configFilePath, []byte("[]"), 0644)
	} else {
	}
}

func mapToDesktopEntry(mruEntry MruEntry) (*desktop.Entry, error) {
	matchingEntry, err := find(destkopEntries, func(entry *desktop.Entry) bool {
		return entry.Exec == mruEntry.Exec
	})
	if err != nil {
		return nil, err
	}
	return matchingEntry, nil
}

func getMruExec() []MruEntry {
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

func updateMruEntry(exec string) {
	mruEntries := getMruExec()
	_, err := find(mruEntries, func(entry MruEntry) bool {
		return entry.Exec == exec
	})
	if err != nil {
		mruEntries = append(mruEntries, MruEntry{Exec: exec, Count: 1})
	} else {
		mruEntries = mapArray(mruEntries, func(entry MruEntry) MruEntry {
			if entry.Exec == exec {
				return MruEntry{Exec: exec, Count: entry.Count + 1}
			}
			return entry
		})
	}
	bytes, err := json.Marshal(mruEntries)
	err = os.WriteFile(configFilePath, bytes, 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		updateMruEntries()
	}
}

func updateMruEntries() {
	mruEntries := getMruExec()
	mapped := mapArray(mruEntries, func(entry MruEntry) *desktop.Entry {
		mapped, err := mapToDesktopEntry(entry)
		if err != nil {
			return nil
		}
		return mapped
	})
	length := len(mapped)
	if length < 16 {
		filler := make([]*desktop.Entry, 16-length)
		for i := range filler {
			mapped = append(mapped, destkopEntries[i])
		}
	}
	mruDesktopEntries = mapped
}
