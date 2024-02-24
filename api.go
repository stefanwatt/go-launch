package main

import (
	"fmt"
	"os/exec"
	"sort"

	"code.rocketnine.space/tslocum/desktop"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

func hideLauncher() {
	cmd := exec.Command("i3-msg", "[class=\"Go-launch\"]", "scratchpad", "show")
	cmd.Start()
}

func (a *App) HideLauncher() {
	hideLauncher()
}
func (a *App) LaunchApp(Exec string) {
	command, args := parseCommand(Exec)
	cmd := exec.Command(command, args...)
	cmd.Start()
	hideLauncher()
}

func (a *App) GetDesktopEntries() []*desktop.Entry {
	return destkopEntries
}

func (a *App) FuzzyFindDesktopEntry(searchTerm string) [][]*desktop.Entry {
	desktopEntryNames := mapArray(destkopEntries, func(entry *desktop.Entry) string {
		return entry.Name
	})
	matches := fuzzy.RankFindNormalizedFold(searchTerm, desktopEntryNames)
	sort.Sort(matches)
	fmt.Println("sorted matches")
	fmt.Println(matches)
	searchResultNames := mapArray(matches, func(match fuzzy.Rank) string {
		return match.Target
	})
	searchResultEntries := mapArray(searchResultNames, func(name string) *desktop.Entry {
		entry, _ := find(destkopEntries, func(entry *desktop.Entry) bool {
			return entry.Name == name
		})
		return entry
	})
	searchResults := make([][]*desktop.Entry, 4)
	for i := range searchResults {
		searchResults[i] = make([]*desktop.Entry, 4)
	}
	size := 4
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			index := i*size + j
			if index < len(matches) {
				searchResults[i][j] = searchResultEntries[index]
			}
		}
	}
	return searchResults
}
