package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/mitchellh/hashstructure"
)

var fallbackId = 0

func getDesktopEntry(lines []string) *Entry {
	entry := &Entry{}
	entry.Name = getAttribute(lines, "Name")
	entry.Icon = getIcon(lines)
	entry.Path = getAttribute(lines, "Path")
	entry.Exec = getExec(lines)
	entry.NoDisplay = getAttribute(lines, "NoDisplay") == "true"
	entry.Type = getType(lines)
	hash, err := hashstructure.Hash(entry, nil)
	if err != nil {
		print("unable to generate id for desktop entry")
		print(err.Error())
		entry.Id = strconv.Itoa(fallbackId)
		fallbackId++
	}
	entry.Id = fmt.Sprint(hash)
	return entry
}

func getAttribute(lines []string, attribute string) string {
	found, err := find(lines, func(line string) bool {
		return strings.HasPrefix(line, attribute+"=")
	})

	if err != nil {
		return "not found"
	} else {
		return strings.Split(found, "=")[1]
	}
}

func getIcon(lines []string) string {
	icon, err := find(lines, func(line string) bool {
		return strings.HasPrefix(line, "Icon=")
	})
	if err != nil {
		if strings.Contains(strings.Join(lines, "\n"), "nitrogen") {
			print(strings.Join(lines, "\n"))
		}
		return DEFAULT_ICON
	}
	value := strings.Split(icon, "=")[1]
	zafiroIcon, err := mapZafiroIcon(value)
	if err != nil {
		print("zafiro icon not found for icon name: " + value)
		return DEFAULT_ICON
	}
	ppath := mapIconPath(*zafiroIcon)
	// print("icon will be at " + ppath)
	return ppath
}

func getExec(lines []string) string {
	exec, err := find(lines, func(line string) bool {
		return strings.HasPrefix(line, "Exec=")
	})
	if err != nil {
		return "not found"
	}
	regex := regexp.MustCompile(`=`)
	value := regex.Split(exec, 2)[1]
	return trimExec(value)
}

func getType(lines []string) EntryType {
	value := getAttribute(lines, "Type")
	switch value {
	case "Application":
		return Application
	case "Link":
		return Link
	case "Directory":
		return Directory
	}
	return Unknown
}
