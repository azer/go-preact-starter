package preact

import (
	"github.com/fsnotify/fsnotify"
	"strings"
	"time"
)

const CHANGE_THRESHOLD_MS = 1000

type Watcher struct {
	Files   []string
	Watcher *fsnotify.Watcher
}

func (watcher *Watcher) Add() error {
	for _, file := range watcher.Files {
		if err := watcher.Watcher.Add(file); err != nil {
			return err
		}
	}

	return nil
}

func (watcher *Watcher) Listen() error {
	var err error
	watcher.Watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	defer watcher.Watcher.Close()

	var lastchange = 0

	done := make(chan bool)
	go func() {
		for {

			select {
			case event := <-watcher.Watcher.Events:
				now := int(time.Now().UnixNano() / 1000000)

				// TODO: better to have debouncing here
				if event.Op&fsnotify.Write == fsnotify.Write && now-lastchange > CHANGE_THRESHOLD_MS {
					lastchange = now
					go watcher.OnWrite(event)
				}
			case err := <-watcher.Watcher.Errors:
				log.Error("File watching error: %v", err)
			}
		}
	}()

	err = watcher.Add()
	if err != nil {
		return err
	}

	<-done

	return nil
}

func (watcher *Watcher) OnWrite(event fsnotify.Event) {
	if strings.Contains(event.Name, "dist.js") {
		runtime.CleanCache()
		return
	}
}

func WatchCodeChanges() {
	watcher := Watcher{
		Files: []string{
			"./public/dist.js",
		},
	}

	if err := watcher.Listen(); err != nil {
		log.Error("Can not listen file changes: %v", err)
	}
}
