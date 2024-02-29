package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"code.rocketnine.space/tslocum/desktop"
	"github.com/mitchellh/hashstructure"
)

type EntryType int

// All entry types
const (
	Unknown     EntryType = iota // Unspecified or unrecognized
	Application                  // Execute command
	Link                         // Open browser
	Directory                    // Open file manager
)

const sectionHeaderNotFoundError = "section header not found"

func (t EntryType) String() string {
	switch t {
	case Unknown:
		return "Unknown"
	case Application:
		return "Application"
	case Link:
		return "Link"
	case Directory:
		return "Directory"
	}

	return strconv.Itoa(int(t))
}

type Entry struct {
	Name      string
	Icon      string
	Path      string
	Exec      string
	Hash      string
	Type      EntryType
	NoDisplay bool
}

var (
	ZAFIRO_ICONS_PATH = "/home/stefan/Projects/Zafiro-icons/Dark"
	ICONS_BASE_PATH   = ZAFIRO_ICONS_PATH + "/apps/scalable"
	GENERIC_ICON_PATH = ZAFIRO_ICONS_PATH + "/categories/22-Dark/applications-utilities.svg"
	COUNT             = 16
)

func isSameEntry(a *Entry, b *Entry) bool {
	if a == nil || b == nil {
		return true
	}
	return b.Hash == a.Hash ||
		b.Name == a.Name || b.Exec == a.Exec
}

func trimExec(currentEntries []*Entry) []*Entry {
	trimmedEntries := make([]*Entry, len(currentEntries))
	for i, entry := range currentEntries {
		// Make a shallow copy of the entry
		if entry == nil {
			continue
		}
		entryCopy := *entry
		// Modify the Exec field in the copy
		entryCopy.Exec = strings.Fields(entry.Exec)[0]
		// Assign the modified copy to the new slice
		trimmedEntries[i] = &entryCopy
	}
	return trimmedEntries
}

func fillUpDesktopEntries(currentEntries []*Entry) []*Entry {
	if len(currentEntries) >= COUNT {
		return currentEntries
	}

	filler := filter(desktopEntries, func(a *Entry) bool {
		found, _ := find(currentEntries, func(b *Entry) bool {
			return isSameEntry(a, b)
		})
		return found == nil
	})

	updatedEntries := append([]*Entry(nil), currentEntries...)
	for len(updatedEntries) < COUNT && len(filler) > 0 {
		updatedEntries = append(updatedEntries, filler[0])
		filler = filler[1:]
	}

	if len(updatedEntries) < COUNT {
		print("insufficient amount of filler desktop entries")
	}

	return updatedEntries
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

func getDesktopEntryOfDir(dir string) []*Entry {
	fmt.Println("Getting desktop entries of dir:", dir)
	var entries []*Entry

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return entries
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.HasSuffix(file.Name(), ".desktop") {
			filePath := filepath.Join(dir, file.Name())
			entriesOfFile, err := parseDesktopFile(filePath)
			if err != nil {
				fmt.Println("Error parsing desktop file:", filePath, ":", err)
				continue
			}
			entries = append(entries, entriesOfFile...)
		}
	}
	return entries
}

func parseDesktopEntryLines(lines []string) *Entry {
	entry := &Entry{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "Name":
			entry.Name = value
		case "Icon":
			entry.Icon = value
		case "Path":
			entry.Path = value
		case "Exec":
			entry.Exec = value
		case "NoDisplay":
			entry.NoDisplay = value == "true"
		case "Type":
			switch value {
			case "Application":
				entry.Type = Application
			case "Link":
				entry.Type = Link
			case "Directory":
				entry.Type = Directory
			case "Unknown":
				entry.Type = Unknown
			}
		}
	}
	return entry
}

func parseDesktopFile(path string) ([]*Entry, error) {
	print("parsing desktop entry: " + path)
	file, err := os.ReadFile(path) // Use ioutil.ReadFile in Go versions before 1.16
	if err != nil {
		return nil, err
	}

	content := string(file)
	regex := regexp.MustCompile(`\[.*\]`)
	parts := regex.Split(content, -1)
	entries := mapArray(parts, func(part string) *Entry {
		lines := strings.Split(part, "\n")
		return parseDesktopEntryLines(lines)
	})

	entries = filter(entries, func(entry *Entry) bool {
		return !entry.NoDisplay && entry.Type == Application
	})

	return entries, nil
}

func getDesktopEntriesOfDir(desktopEntryDirs []string) []*Entry {
	desktopEntries2d := mapArray(desktopEntryDirs, func(dir string) []*Entry {
		return getDesktopEntryOfDir(dir)
	})
	return flatten(desktopEntries2d)
}

func getDesktopEntries() []*Entry {
	dataDirs := desktop.DataDirs()
	desktopEntries := getDesktopEntriesOfDir(dataDirs)
	for _, entry := range desktopEntries {

		hash, err := hashstructure.Hash(entry, nil)
		if err == nil {
			entry.Hash = fmt.Sprint(hash)
		}

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
	return desktopEntries
}
