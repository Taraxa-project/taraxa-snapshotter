package commands

import (
	"fmt"
	"path"
	"strconv"

	"github.com/Taraxa-project/taraxa-snapshotter/internal/dir"
	"github.com/Taraxa-project/taraxa-snapshotter/internal/snapshot"
	"github.com/Taraxa-project/taraxa-snapshotter/internal/upload/ipfs"
	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

// Publish publishes a local snapshot to IPFS
func Publish(ctx *cli.Context) error {
	ipfsClient := ipfs.NewIPFSClient(ctx.String("ipfs-url"))

	i, err := strconv.Atoi(ctx.Args().Get(0))
	if err != nil {
		log.Fatal("Incorrect snapshot number")
	}

	snapshotDir := dir.NewDirEntry(path.Join(ctx.String("base-dir"), fmt.Sprintf("db%d", i)))
	snapshot := snapshot.NewSnapshotFromDir(snapshotDir)
	snapshot.Upload(ipfsClient)

	return nil
}
