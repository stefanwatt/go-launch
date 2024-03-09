package main

import (
	Config "go-launch/backend/config"
	Desktop "go-launch/backend/desktop"
	Log "go-launch/backend/log"
	Utils "go-launch/backend/utils"
	"os/exec"
	"path"
	"strconv"

	Runtime "github.com/wailsapp/wails/v2/pkg/runtime"
)

func hideLauncher() {
	cmd := exec.Command("i3-msg", "[class=\"Go-launch\"]", "scratchpad", "show")
	cmd.Start()
}

func (a *App) HideLauncher() {
	hideLauncher()
}

func (a *App) LaunchApp(Id string) {
	desktopEntry, err := Utils.Find(Desktop.DesktopEntries, func(entry *Desktop.Entry) bool {
		return entry.Id == Id
	})
	if err != nil {
		Log.Print("entry not found for id " + Id)
		return
	}
	Log.Print("launching app with gtk-launch " + desktopEntry.Path)
	// extract just the filename with extension forom the path
	cmd := exec.Command("gtk-launch", path.Base(desktopEntry.Path))
	cmderr := cmd.Start()
	if cmderr != nil {
		Log.Print("error launching app " + cmderr.Error())
	}
	hideLauncher()
	Desktop.IncrementMruEntry(desktopEntry)
	Runtime.EventsEmit(a.ctx, "desktop-entries-changed")
}

func (a *App) GetDesktopEntries() []*Desktop.Entry {
	return Desktop.DesktopEntries
}

func (a *App) FuzzyFindDesktopEntry(searchTerm string) [][]*Desktop.Entry {
	Log.Print("searchTerm = " + searchTerm)
	var searchResultEntries []*Desktop.Entry
	if searchTerm == "" {
		searchResultEntries = Desktop.MruDesktopEntries
	} else {
		searchResultEntries = Utils.FuzzyFindObj(searchTerm, Desktop.DesktopEntries, []string{"Name", "Exec"})
	}
	searchResultEntries = Desktop.RemoveDuplicateEntries(searchResultEntries)

	if searchTerm == "" {
		searchResultEntries = Desktop.FillUpDesktopEntries(searchResultEntries)
	}

	for _, entry := range searchResultEntries {
		if entry == nil {
			Log.Print("nil entry")
			continue
		}
	}

	searchResults := make([][]*Desktop.Entry, Config.ROWS)
	for i := range searchResults {
		searchResults[i] = make([]*Desktop.Entry, Config.COLS)
	}
	for i := 0; i < Config.ROWS; i++ {
		for j := 0; j < Config.COLS; j++ {
			index := i*Config.ROWS + j
			if index < len(searchResultEntries) {
				searchResults[i][j] = searchResultEntries[index]
			}
		}
	}

	Log.Print("returning " + strconv.Itoa(len(searchResultEntries)) + " results")

	return searchResults
}
