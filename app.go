package main

import (
	"code.rocketnine.space/tslocum/desktop"
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
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
	cmd := exec.Command("i3-msg", "[class=\"Nvim-wails\"]", "scratchpad", "show")
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

func findMatchingIcon(entryName string) string {
	iconsDir := "/home/stefan/Projects/wails-test/frontend/static/app-icons/"
	// Normalize entry name to lower case for case-insensitive comparison
	entryName = strings.ToLower(entryName)

	var partialMatches []string
	exactMatch := ""

	filepath.Walk(iconsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fileName := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
			fileNameLower := strings.ToLower(fileName)

			// Check for an exact match
			if fileNameLower == entryName || fileNameLower == entryName+"-icon" {
				exactMatch = path
				return filepath.SkipDir // Found an exact match, no need to continue walking the directory
			}

			// Check if the file name contains the entry name
			if strings.Contains(fileNameLower, entryName) {
				partialMatches = append(partialMatches, path)
			}
		}
		return nil
	})

	if exactMatch != "" {
		return exactMatch
	}

	if len(partialMatches) > 0 {
		// Here you can implement a strategy to choose the best match from partialMatches
		// For simplicity, we return the first partial match
		return partialMatches[0]
	}

	// No match found, return an empty string or a default icon path
	return ""
}

func (a *App) GetDesktopEntries() []*desktop.Entry {
	dataDirs := desktop.DataDirs()
	desktopEntries2d, _ := desktop.Scan(dataDirs)
	flatEntries := flatten(desktopEntries2d)
	for _, entry := range flatEntries {
		if entry.Icon != "" && !strings.HasPrefix(entry.Icon, "/") {
			iconPath := findMatchingIcon(entry.Name)
			if iconPath != "" {
				entry.Icon = iconPath
			}
		}
	}
	return flatEntries
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
