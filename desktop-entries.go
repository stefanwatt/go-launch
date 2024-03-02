package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"code.rocketnine.space/tslocum/desktop"
)

var COUNT = COLS * ROWS

type EntryType int

const (
	Unknown EntryType = iota
	Application
	Link
	Directory
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
	Id        string
	Name      string
	Icon      string
	Path      string
	Exec      string
	Type      EntryType
	NoDisplay bool
}

func initDesktopEntries() []*Entry {
	dataDirs := desktop.DataDirs()
	desktopEntries := getDesktopEntriesOfDirs(dataDirs)
	for _, entry := range desktopEntries {
		zafiroIcon, err := mapZafiroIcon(entry.Icon)
		if err == nil {
			src := ICONS_BASE_PATH + "/" + *zafiroIcon
			copyIcon(src, entry.Icon)
		} else {
			print("could not copy icon for " + entry.Name)
		}
	}
	return desktopEntries
}

func getDesktopEntriesOfDirs(desktopEntryDirs []string) []*Entry {
	desktopEntries2d := mapArray(desktopEntryDirs, func(dir string) []*Entry {
		return getDesktopEntriesOfDir(dir)
	})
	return flatten(desktopEntries2d)
}

func getDesktopEntriesOfDir(dir string) []*Entry {
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
			entriesOfFile, err := getDesktopEntries(filePath)
			if err != nil {
				fmt.Println("Error parsing desktop file:", filePath, ":", err)
				continue
			}
			entries = append(entries, entriesOfFile...)
		}
	}
	return entries
}

func getDesktopEntries(path string) ([]*Entry, error) {
	file, err := os.ReadFile(path) // Use ioutil.ReadFile in Go versions before 1.16
	if err != nil {
		return nil, err
	}

	content := string(file)
	regex := regexp.MustCompile(`^\[.*\]`)
	parts := regex.Split(content, -1)
	entries := mapArray(parts, func(part string) *Entry {
		lines := strings.Split(part, "\n")
		return getDesktopEntry(lines)
	})

	entries = filter(entries, func(entry *Entry) bool {
		return !entry.NoDisplay && entry.Type == Application
	})

	return entries, nil
}

// ----------------------UTILS-------------------------------
func isSameEntry(a *Entry, b *Entry) bool {
	if a == nil || b == nil {
		return true
	}
	return b.Id == a.Id ||
		b.Name == a.Name || b.Exec == a.Exec
}

func trimExec(exec string) string {
	fields := strings.Fields(exec)
	if len(fields) == 0 {
		return exec
	}
	if strings.HasPrefix(fields[0], "flatpak") {
		// can be like: flatpak run --branch=stable --arch=x86_64 --comman…top --file-forwarding org.signal.Signal @@u %U @@
		// trim off @@u %U @@ or similar endings
		// loop over fields and only when field starts with @@ or %U or similar, stop and return the string
		value := ""
		for _, field := range fields {
			if strings.HasPrefix(field, "@@") || strings.HasPrefix(field, "%") {
				break
			}
			value += field + " "
		}
		return strings.Trim(value, " ")
	}
	return fields[0]
}

func hasNilEntries(entries []*Entry) bool {
	for _, entry := range entries {
		if entry == nil {
			return true
		}
	}
	return false
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
