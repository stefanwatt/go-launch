package main

import (
	"fmt"
	"os/exec"
	"sort"
	"strconv"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

var (
	COLS = 4
	ROWS = 4
)

func hideLauncher() {
	cmd := exec.Command("i3-msg", "[class=\"Go-launch\"]", "scratchpad", "show")
	cmd.Start()
}

func (a *App) HideLauncher() {
	hideLauncher()
}

func (a *App) LaunchApp(Exec string) {
	print("launching app " + Exec)
	command, args := parseCommand(Exec)
	cmd := exec.Command(command, args...)
	cmd.Start()
	hideLauncher()
	updateMruEntry(Exec)
}

func (a *App) GetDesktopEntries() []*Entry {
	return desktopEntries
}

func getSearchResultEntriesFuzzy(searchTerm string) []*Entry {
	desktopEntryNames := mapArray(desktopEntries, func(entry *Entry) string {
		return entry.Name
	})
	matches := fuzzy.RankFindNormalizedFold(searchTerm, desktopEntryNames)
	sort.Sort(matches)
	fmt.Println(matches)
	searchResultNames := mapArray(matches, func(match fuzzy.Rank) string {
		return match.Target
	})
	result := mapArray(searchResultNames, func(name string) *Entry {
		entry, _ := find(desktopEntries, func(entry *Entry) bool {
			return entry.Name == name
		})
		return entry
	})
	return result
}

func removeDuplicateEntries(searchResultEntries []*Entry) []*Entry {
	filtered := []*Entry{}
	for i := range searchResultEntries {
		found, _ := find(filtered, func(entry *Entry) bool {
			if entry == nil {
				return false
			}
			if searchResultEntries[i] == nil {
				searchResultEntries[i] = entry
				return false
			}
			return isSameEntry(entry, searchResultEntries[i])
		})
		if found == nil { // Append if not found, meaning no duplicate
			filtered = append(filtered, searchResultEntries[i])
		}
	}
	return filtered
}

func (a *App) FuzzyFindDesktopEntry(searchTerm string) [][]*Entry {
	print("searchTerm = " + searchTerm)
	initDesktopEntries()
	var searchResultEntries []*Entry
	if searchTerm == "" {
		searchResultEntries = mruDesktopEntries
	} else {
		searchResultEntries = getSearchResultEntriesFuzzy(searchTerm)
	}
	searchResultEntries = removeDuplicateEntries(searchResultEntries)

	if searchTerm == "" {
		searchResultEntries = fillUpDesktopEntries(searchResultEntries)
	}

	for _, entry := range searchResultEntries {
		if entry == nil {
			print("nil entry")
			continue
		}
		entry.Exec = trimExec(entry.Exec)
	}

	searchResults := make([][]*Entry, ROWS)
	for i := range searchResults {
		searchResults[i] = make([]*Entry, COLS)
	}
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			index := i*ROWS + j
			if index < len(searchResultEntries) {
				searchResults[i][j] = searchResultEntries[index]
			}
		}
	}

	print("returning " + strconv.Itoa(len(searchResultEntries)) + " results")
	return searchResults
}
