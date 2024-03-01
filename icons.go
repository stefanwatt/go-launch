package main

import (
	"errors"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

var (
	DEFAULT_ICON   = "default.svg"
	DIR_ENTRIES, _ = os.ReadDir(ICONS_BASE_PATH)
	ZAFIRO_ICONS   = mapArray(DIR_ENTRIES, func(entry os.DirEntry) string {
		return entry.Name()
	})
)

func mapZafiroIcon(appIconName string) (*string, error) {
	if appIconName == "" {
		return nil, errors.New("icon not found")
	}
	matches := fuzzy.RankFindNormalizedFold(appIconName, ZAFIRO_ICONS)
	if len(matches) == 0 {
		print("no matches found for " + appIconName)
		return nil, errors.New("icon not found")
	}
	sort.Sort(matches)
	print(appIconName+" matched with ", matches[0].Target)
	return &matches[0].Target, nil
}

func mapIconPath(zafiroIconPath string) string {
	if !strings.HasSuffix(zafiroIconPath, "svg") {
		return zafiroIconPath + ".svg"
	}
	return zafiroIconPath
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
