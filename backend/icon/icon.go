package icon

import (
	"errors"
	Config "go-launch/backend/config"
	Log "go-launch/backend/log"
	Utils "go-launch/backend/utils"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

var (
	ZAFIRO_ICONS_PATH = path.Join(Config.HOME, "Projects/Zafiro-icons/Dark")
	BASE_PATH         = ZAFIRO_ICONS_PATH + "/apps/scalable"
	OUTPUT_BASE_PATH  = path.Join(Config.HOME, "Projects/go-launch/frontend/static/app-icons/")
	GENERIC_ICON_PATH = ZAFIRO_ICONS_PATH + "/categories/22-Dark/applications-utilities.svg"
	DEFAULT_ICON      = "default.svg"
	DIR_ENTRIES, _    = os.ReadDir(BASE_PATH)
	ZAFIRO_ICONS      = Utils.MapArray(DIR_ENTRIES, func(entry os.DirEntry) string {
		return entry.Name()
	})
)

func MapZafiroIcon(appIconName string) (*string, error) {
	if appIconName == "" {
		return nil, errors.New("icon not found")
	}
	matches := fuzzy.RankFindNormalizedFold(appIconName, ZAFIRO_ICONS)
	if len(matches) == 0 {
		Log.Print("no matches found for " + appIconName)
		return nil, errors.New("icon not found")
	}
	sort.Sort(matches)
	return &matches[0].Target, nil
}

func MapIconPath(zafiroIconPath string) string {
	if !strings.HasSuffix(zafiroIconPath, "svg") {
		return zafiroIconPath + ".svg"
	}
	return zafiroIconPath
}

func CopyIcon(src string, filename string) {
	dest := path.Join(OUTPUT_BASE_PATH, filename)
	// check if file exist at dest
	if _, err := os.Stat(dest); err == nil {
		return
	}
	Log.Print("copying icon from " + src + " to " + dest)
	cmd := exec.Command("cp", src, dest)
	cmd.Start()
}
