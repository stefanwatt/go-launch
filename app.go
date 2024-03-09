// TODO:add "apps" like calculator/translator

// TODO:allow to give list of arbitrary commands
package main

import (
	"context"
	Desktop "go-launch/backend/desktop"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	Desktop.InitMru()
	Desktop.DesktopEntries = Desktop.InitDesktopEntries()
	Desktop.MruDesktopEntries = Desktop.GetMruDesktopEntries()
}
