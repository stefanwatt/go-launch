package desktop_entries

import (
	"fmt"
	Config "go-launch/backend/config"
	Icon "go-launch/backend/icon"
	Log "go-launch/backend/log"
	Utils "go-launch/backend/utils"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"code.rocketnine.space/tslocum/desktop"
	"github.com/fsnotify/fsnotify"
)

var (
	DesktopEntries    []*Entry
	MruDesktopEntries []*Entry
	COUNT             = Config.COLS * Config.ROWS
)

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
	Terminal  bool
}

func onWrite(event fsnotify.Event) {
	DesktopEntries = Utils.Filter(DesktopEntries, func(entry *Entry) bool {
		return entry.Path != event.Name
	})
	newEntries, err := getDesktopEntries(event.Name)
	if err != nil {
		Log.Print("error getting desktop entries")
		return
	}
	DesktopEntries = append(DesktopEntries, newEntries...)
}

func onDelete(event fsnotify.Event) {
	path := event.Name
	DesktopEntries = Utils.Filter(DesktopEntries, func(entry *Entry) bool {
		return entry.Path != path
	})
	RemoveMruEntry(path)
}

func InitDesktopEntries() []*Entry {
	dataDirs := desktop.DataDirs()
	var wg sync.WaitGroup
	wg.Add(len(dataDirs))
	for _, dir := range dataDirs {
		go func(directory string) {
			defer wg.Done()
			ObserveDirectory(directory, onWrite, onDelete)
		}(dir)
	}
	desktopEntries := getDesktopEntriesOfDirs(dataDirs)
	for _, entry := range desktopEntries {
		zafiroIcon, err := Icon.MapZafiroIcon(entry.Icon)
		if err == nil {
			src := Icon.BASE_PATH + "/" + *zafiroIcon
			Icon.CopyIcon(src, entry.Icon)
		} else {
			Log.Print("could not copy icon for " + entry.Name)
		}
	}
	return desktopEntries
}

func FillUpDesktopEntries(currentEntries []*Entry) []*Entry {
	if len(currentEntries) >= COUNT {
		return currentEntries
	}

	filler := Utils.Filter(DesktopEntries, func(a *Entry) bool {
		found, _ := Utils.Find(currentEntries, func(b *Entry) bool {
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
		Log.Print("insufficient amount of filler desktop entries")
	}

	return updatedEntries
}

func RemoveDuplicateEntries(searchResultEntries []*Entry) []*Entry {
	filtered := []*Entry{}
	for i := range searchResultEntries {
		found, _ := Utils.Find(filtered, func(entry *Entry) bool {
			if entry == nil {
				return false
			}
			if searchResultEntries[i] == nil {
				searchResultEntries[i] = entry
				return false
			}
			return isSameEntry(entry, searchResultEntries[i])
		})
		if found == nil { // Append if not found, meaning no duplicate
			filtered = append(filtered, searchResultEntries[i])
		}
	}
	return filtered
}

func getDesktopEntriesOfDirs(desktopEntryDirs []string) []*Entry {
	desktopEntries2d := Utils.MapArray(desktopEntryDirs, func(dir string) []*Entry {
		return getDesktopEntriesOfDir(dir)
	})
	return Utils.Flatten(desktopEntries2d)
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
	entries := Utils.MapArray(parts, func(part string) *Entry {
		lines := strings.Split(part, "\n")
		return getDesktopEntry(lines, path)
	})

	entries = Utils.Filter(entries, func(entry *Entry) bool {
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
