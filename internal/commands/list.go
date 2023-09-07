package commands

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/Taraxa-project/taraxa-snapshotter/internal/dir"
	"github.com/Taraxa-project/taraxa-snapshotter/internal/snapshot"
	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

// List lists local snapshot numbers (snapshot number is block number after which snapshot was taken)
func List(ctx *cli.Context) error {
	log.Info("Snapshots:")

	baseDir := dir.NewDirEntry(ctx.String("base-dir"))
	dirs := baseDir.SnapshotDirs()

	var snapshots []*snapshot.Snapshot
	for _, d := range dirs {
		snapshots = append(snapshots, snapshot.NewSnapshotFromDir(d))
	}

	sort.Slice(snapshots, func(i, j int) bool { return snapshots[i].Number < snapshots[j].Number })

	for _, s := range snapshots {
		json, err := json.Marshal(s)
		if err != nil {
			log.WithError(err).Fatal("Could not marshal json")
		}
		fmt.Println(string(json))
	}

	return nil
}
