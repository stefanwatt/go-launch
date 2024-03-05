package main

import (
	"os/exec"
	"strconv"
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

func (a *App) LaunchApp(Id string) {
	desktopEntry, err := find(desktopEntries, func(entry *Entry) bool {
		return entry.Id == Id
	})
	if err != nil {
		print("entry not found for id " + Id)
		return
	}
	command, args := parseCommand(desktopEntry.Exec)
	print("launching app with exec=" + desktopEntry.Exec + "; executing cmd: " + command)
	cmd := exec.Command(command, args...)
	cmderr := cmd.Start()
	if cmderr != nil {
		print("error launching app " + cmderr.Error())
	}
	hideLauncher()
	updateMruEntry(desktopEntry)
}

func (a *App) GetDesktopEntries() []*Entry {
	return desktopEntries
}

func (a *App) FuzzyFindDesktopEntry(searchTerm string) [][]*Entry {
	print("searchTerm = " + searchTerm)
	var searchResultEntries []*Entry
	if searchTerm == "" {
		searchResultEntries = mruDesktopEntries
	} else {
		searchResultEntries = fuzzyFindObj(searchTerm, desktopEntries, []string{"Name", "Exec"})
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
