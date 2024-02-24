package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"code.rocketnine.space/tslocum/desktop"
)

var ZAFIRO_ICONS_PATH = "/home/stefan/Projects/Zafiro-icons/Dark"
var ICONS_BASE_PATH = ZAFIRO_ICONS_PATH + "/apps/scalable"
var GENERIC_ICON_PATH = ZAFIRO_ICONS_PATH + "/categories/22-Dark/applications-utilities.svg"

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
