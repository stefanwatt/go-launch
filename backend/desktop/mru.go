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
	Id      string `json:"id"`
	Count   int    `json:"count"`
	Deleted bool   `json:"deleted"`
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

func IncrementMruEntry(desktopEntry *Entry) {
	UpdateMruEntry(desktopEntry, 1, false)
}

func UpdateMruEntry(desktopEntry *Entry, countAdd int, deleted bool) {
	mruEntries := getMruEntries()
	_, err := Utils.Find(mruEntries, func(entry MruEntry) bool {
		return entry.Id == desktopEntry.Id
	})
	if err != nil {
		mruEntries = append(mruEntries, MruEntry{Id: desktopEntry.Id, Count: 1})
	} else {
		mruEntries = Utils.MapArray(mruEntries, func(entry MruEntry) MruEntry {
			if entry.Id == desktopEntry.Id {
				return MruEntry{Id: desktopEntry.Id, Count: entry.Count + countAdd, Deleted: deleted}
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

func RemoveMruEntry(path string) error {
	index := -1
	for i, entry := range MruDesktopEntries {
		if entry.Path == path {
			index = i
			break
		}
	}
	if index == -1 {
		return fmt.Errorf("MRU entry with path %s not found", path)
	}
	entry := MruDesktopEntries[index]
	UpdateMruEntry(entry, 0, true)
	MruDesktopEntries = append(MruDesktopEntries[:index], MruDesktopEntries[index+1:]...)

	return nil
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
