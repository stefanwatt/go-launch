package main

import (
	"context"
	"os"

	"code.rocketnine.space/tslocum/desktop"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

var destkopEntries []*desktop.Entry
var mruDesktopEntries []*desktop.Entry
var homeDir, _ = os.UserHomeDir()

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	initMru()
	destkopEntries = getDesktopEntries()
	updateMruEntries()
}
