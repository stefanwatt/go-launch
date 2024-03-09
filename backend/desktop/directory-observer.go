package desktop_entries

import (
	"context"
	"log"

	Log "go-launch/backend/log"

	"github.com/fsnotify/fsnotify"
)

func ObserveDirectory(ctx context.Context,
	directory string,
	onWrite func(fsnotify.Event, context.Context),
	onDelete func(fsnotify.Event, context.Context),
) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		Log.Print(err)
		log.Fatal(err)
	}
	defer watcher.Close()
	err = watcher.Add(directory)
	if err != nil {
		Log.Print("skipping directory " + directory + " because of error")
		Log.Print(err)
		return
	}
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op == fsnotify.Create || event.Op == fsnotify.Write {
				onWrite(event, ctx)
			}
			if event.Op == fsnotify.Remove {
				onDelete(event, ctx)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			Log.Print("Error:", err)
		}
	}
}
