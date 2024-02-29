package main

import (
	"errors"
	"os"
	"os/exec"
	"path"
	"strings"
)

var DEFAULT_ICON = "default.svg"

func mapZafiroIcon(appName string) (string, error) {
	dirEntries, err := os.ReadDir(ICONS_BASE_PATH)
	if err != nil {
		return "", err
	}
	for _, dir := range dirEntries {
		// try to find exact match first
		if strings.Contains(dir.Name(), appName) {
			return dir.Name(), nil
		}
	}
	return "", errors.New("icon not found")
}

func mapIconPath(zafiroIconPath string, entryIcon string) string {
	if zafiroIconPath == "" {
		return DEFAULT_ICON
	}
	updatedEntryIcon := zafiroIconPath
	if !strings.HasSuffix(entryIcon, "svg") {
		updatedEntryIcon += ".svg"
	}
	destFileName := updatedEntryIcon
	if !strings.HasSuffix(destFileName, "svg") {
		destFileName += ".svg"
	}
	return destFileName
}

func copyIcon(src string, filename string) {
	dest := path.Join(ICONS_OUTPUT_BASE_PATH, filename)
	// check if file exist at dest
	if _, err := os.Stat(dest); err == nil {
		return
	}
	print("copying icon from " + src + " to " + dest)
	cmd := exec.Command("cp", src, dest)
	cmd.Start()
}
