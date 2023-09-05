package commands

import (
	"os"
	"path"
	"path/filepath"
	"regexp"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

// List lists local snapshot numbers (snapshot number is block number after which snapshot was taken)
func List(ctx *cli.Context) error {
	log.Info("Listening for new snapshots...")

	matcher := regexp.MustCompile(`^db\d+$`)

	var files []string
	err := filepath.WalkDir(ctx.String("base-dir"), func(dir string, info os.DirEntry, err error) error {
		if !info.IsDir() {
			return nil
		}

		baseName := path.Base(dir)
		if !matcher.Match([]byte(baseName)) {
			return nil
		}

		files = append(files, dir)
		return nil
	})
	if err != nil {
		log.WithError(err).Fatal("Failed to walk directory")
	}

	log.Info("Found snapshots:")
	for _, file := range files {
		log.Info(file)
	}

	return nil
}
