package commands

import (
	"path"
	"regexp"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// Listen listens for new snapshots and upload to ipfs when found
func Listen(ctx *cli.Context) error {
	log.Info("Listening for new snapshots...")

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.WithError(err).Fatal("Failed to create watcher")
	}
	defer watcher.Close()

	matcher := regexp.MustCompile(`^db\d+$`)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Create != fsnotify.Create {
					continue
				}

				baseName := path.Base(event.Name)
				if !matcher.Match([]byte(baseName)) {
					continue
				}

				log.Println("Event:", event, baseName)

			case err := <-watcher.Errors:
				log.WithError(err).Fatal("Failed to watch directory")
			}
		}
	}()

	err = watcher.Add(ctx.String("base-dir"))
	if err != nil {
		log.Fatal(err)
	}

	for {
	}

	return nil
}
