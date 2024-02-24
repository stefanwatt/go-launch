package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"

	"code.rocketnine.space/tslocum/desktop"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

var ZAFIRO_ICONS_PATH = "/home/stefan/Projects/Zafiro-icons/Dark"
var ICONS_BASE_PATH = ZAFIRO_ICONS_PATH + "/apps/scalable"
var GENERIC_ICON_PATH = ZAFIRO_ICONS_PATH + "/categories/22-Dark/applications-utilities.svg"

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}
func findIcon(appName string) string {
	dirEntries, error := os.ReadDir(ICONS_BASE_PATH)
	if error != nil {
		return ""
	}
	for _, entry := range dirEntries {
		// try to find exact match first
		if strings.Contains(entry.Name(), appName) {
			return entry.Name()
		}
	}
	return ""
}

var destkopEntries []*desktop.Entry

func getDesktopEntries() []*desktop.Entry {
	dataDirs := desktop.DataDirs()
	desktopEntries2d, _ := desktop.Scan(dataDirs)
	flatEntries := flatten[*desktop.Entry](desktopEntries2d)
	for _, entry := range flatEntries {
		if entry.Icon == "" || strings.HasPrefix(entry.Icon, "/") {
			fmt.Println("Icon not found for ", entry.Name)
			src := GENERIC_ICON_PATH
			iconName := "default.svg"
			dest := "/home/stefan/Projects/go-launch/frontend/static/app-icons/" + iconName
			cmd := exec.Command("cp", src, dest)
			cmd.Start()
			entry.Icon = iconName
			continue
		}
		iconPath := findIcon(entry.Icon)

		if iconPath != "" {
			entry.Icon = iconPath
			if !strings.HasSuffix(entry.Icon, "svg") {
				entry.Icon += ".svg"
			}
			src := ICONS_BASE_PATH + "/" + iconPath
			destFileName := entry.Icon
			if !strings.HasSuffix(destFileName, "svg") {
				destFileName += ".svg"
			}
			dest := "/home/stefan/Projects/go-launch/frontend/static/app-icons/" + destFileName
			fmt.Println(dest)
			cmd := exec.Command("cp", src, dest)
			err := cmd.Start()
			fmt.Println(err)
		} else {

			src := GENERIC_ICON_PATH
			iconName := "default.svg"
			dest := "/home/stefan/Projects/go-launch/frontend/static/app-icons/" + iconName
			cmd := exec.Command("cp", src, dest)
			cmd.Start()
			entry.Icon = iconName
		}
	}
	return flatEntries

}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	destkopEntries = getDesktopEntries()
}

func parseCommand(commandWithArgs string) (string, []string) {
	// Split the command and its arguments by spaces.
	parts := strings.Fields(commandWithArgs)

	// The first part is the command.
	command := parts[0]

	// The rest are the arguments.
	args := parts[1:]

	return command, args
}
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
	searchResults := make([][]*desktop.Entry, 4) // Replace YourType with the actual type you want, e.g., int, string, etc.
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

func find[T any](arr []T, f func(T) bool) (T, error) {
	var zero T
	for _, value := range arr {
		if f(value) {
			return value, nil
		}
	}
	return zero, fmt.Errorf("no match found")
}

func mapArray[T any, U any](arr []T, f func(T) U) []U {
	var result []U
	for _, value := range arr {
		result = append(result, f(value))
	}
	return result
}

func flatten[T any](slice [][]T) []T {
	var flatSlice []T
	for _, innerSlice := range slice {
		for _, value := range innerSlice {
			flatSlice = append(flatSlice, value)
		}
	}
	return flatSlice
}
