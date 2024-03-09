package config

import (
	"os"
	"path"
)

var (
	LAUNCH_TERMINAL_APP_CMD = "wezterm -e "
	HOME, _                 = os.UserHomeDir()
	CONFIG_FILE_PATH        = path.Join(HOME, ".config/go-launch/mru.json")
	COLS                    = 4
	ROWS                    = 4
)
