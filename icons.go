package main

import (
	"errors"
	"os"
	"os/exec"
	"path"
	"strings"
)

var (
	DEFAULT_ICON  = "default.svg"
	dirEntries, _ = os.ReadDir(ICONS_BASE_PATH)
)

func mapZafiroIcon(appIconName string) (string, error) {
	if appIconName == "" {
		return "", errors.New("icon not found")
	}
	for _, file := range dirEntries {
		// try to find exact match first
		filename := file.Name()
		if strings.Contains(filename, ".directory") {
			continue
		}
		if strings.Contains(filename, appIconName) {
			return filename, nil
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
