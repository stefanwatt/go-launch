package desktop_entries

import (
	"encoding/json"
	"fmt"
	Config "go-launch/backend/config"
	Log "go-launch/backend/log"
	Utils "go-launch/backend/utils"
	"os"
	"sort"
)

type MruEntry struct {
	Id    string `json:"id"`
	Count int    `json:"count"`
}

func GetMruDesktopEntries() []*Entry {
	mruEntries := getMruEntries()
	mapped := Utils.MapArray(mruEntries, func(entry MruEntry) *Entry {
		mapped, err := mapToDesktopEntry(entry)
		if err != nil {
			return nil
		}
		return mapped
	})
	mapped = Utils.Filter(mapped, func(e *Entry) bool {
		return e != nil
	})
	if len(mapped) < COUNT {
		Log.Print("found only " + fmt.Sprint(len(mapped)) + " mru entries ; " + "filling up " + fmt.Sprint(COUNT-len(mapped)) + " mru entries")
		return FillUpDesktopEntries(mapped)
	}
	return mapped
}

func UpdateMruEntry(desktopEntry *Entry) {
	mruEntries := getMruEntries()
	_, err := Utils.Find(mruEntries, func(entry MruEntry) bool {
		return entry.Id == desktopEntry.Id
	})
	if err != nil {
		mruEntries = append(mruEntries, MruEntry{Id: desktopEntry.Id, Count: 1})
	} else {
		mruEntries = Utils.MapArray(mruEntries, func(entry MruEntry) MruEntry {
			if entry.Id == desktopEntry.Id {
				return MruEntry{Id: desktopEntry.Id, Count: entry.Count + 1}
			}
			return entry
		})
	}
	bytes, _ := json.Marshal(mruEntries)
	err = os.WriteFile(Config.CONFIG_FILE_PATH, bytes, 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		MruDesktopEntries = GetMruDesktopEntries()
	}
}

func InitMru() {
	_, err := os.Open(Config.CONFIG_FILE_PATH)
	if err != nil {
		os.WriteFile(Config.CONFIG_FILE_PATH, []byte("[]"), 0644)
	}
}

func mapToDesktopEntry(mruEntry MruEntry) (*Entry, error) {
	matchingEntry, err := Utils.Find(DesktopEntries, func(entry *Entry) bool {
		return entry.Id == mruEntry.Id
	})
	if err != nil {
		return nil, err
	}
	return matchingEntry, nil
}

func getMruEntries() []MruEntry {
	file, err := os.ReadFile(Config.CONFIG_FILE_PATH)
	if err != nil {
		Log.Print("mru config file not found")
		return []MruEntry{}
	}
	var mru []MruEntry
	json.Unmarshal(file, &mru)
	sort.Slice(mru, func(i, j int) bool {
		return mru[i].Count > mru[j].Count
	})
	Log.Print("found " + fmt.Sprint(len(mru)) + " mru entries")
	return mru
}
