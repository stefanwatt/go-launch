// TODO:add "apps" like calculator/translator

// TODO:allow to give list of arbitrary commands
package main

import (
	"context"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

var (
	desktopEntries          []*Entry
	mruDesktopEntries       []*Entry
	LAUNCH_TERMINAL_APP_CMD = "wezterm -e "
)

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	initMru()
	desktopEntries = initDesktopEntries()
	mruDesktopEntries = getMruDesktopEntries()
}
