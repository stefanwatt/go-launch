package desktop_entries

import (
	"log"

	Log "go-launch/backend/log"

	"github.com/fsnotify/fsnotify"
)

func ObserveDirectory(directory string, onWrite func(fsnotify.Event), onDelete func(fsnotify.Event)) {
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
				onWrite(event)
			}
			if event.Op == fsnotify.Remove {
				onDelete(event)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			Log.Print("Error:", err)
		}
	}
}
