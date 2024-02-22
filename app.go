package main

import (
	"code.rocketnine.space/tslocum/desktop"
	"context"
	"os/exec"
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

func (a *App) GetDesktopEntries() []*desktop.Entry {
	dataDirs := desktop.DataDirs()
	desktopEntries2d, _ := desktop.Scan(dataDirs)
	flatEntries := flatten[*desktop.Entry](desktopEntries2d)
	for _, entry := range flatEntries {
		if entry.Icon != "" && !strings.HasPrefix(entry.Icon, "/") {
			iconPath := ""
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
