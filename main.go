package main

import (
	"os"

	"github.com/Taraxa-project/taraxa-snapshotter/internal/commands"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var verbosityCount int

var (
	ipfsUrlFlag = &cli.StringFlag{
		Name:    "ipfs-url",
		Value:   "https://ipfs.infura.io:5001",
		Usage:   "URL of the IPFS node",
		EnvVars: []string{"IPFS_URL"},
	}
	ipfsGatewayFlag = &cli.StringFlag{
		Name:    "ipfs-gateway",
		Value:   "https://gateway.infura.io",
		Usage:   "URL of the IPFS gateway",
		EnvVars: []string{"IPFS_GATEWAY"},
	}
	snapshotBaseDirFlag = &cli.StringFlag{
		Name:    "base-dir",
		Value:   "/root/.taraxa/db",
		Usage:   "Base directory where snapshots are located",
		EnvVars: []string{"BASE_DIR"},
	}
	verboseFlag = &cli.BoolFlag{
		Name:    "verbose",
		Aliases: []string{"v"},
		Usage:   "Enable verbose logging",
		Count:   &verbosityCount,
	}
)

// NewListCommand is the factory method for the list command
func NewListCommand() *cli.Command {
	return &cli.Command{
		Name:   "list",
		Action: commands.List,
		Usage:  "Lists snapshots available locally",
	}
}

// NewListRemote is the factory method for the list-remote command
func NewListRemote() *cli.Command {
	return &cli.Command{
		Name:   "list-remote",
		Action: commands.ListRemote,
		Usage:  "Lists snapshots available in the smart contract",
	}
}

// NewPublishCommand is the factory method for the publish command
func NewPublishCommand() *cli.Command {
	return &cli.Command{
		Name:   "publish",
		Action: commands.Publish,
		Usage:  "Publishes a snapshot to IPFS",
	}
}

// NewRestoreCommand is the factory method for the restore command
func NewRestoreCommand() *cli.Command {
	return &cli.Command{
		Name:   "restore",
		Action: commands.Restore,
		Usage:  "Restores a snapshot from IPFS",
	}
}

// NewListenCommand is the factory method for the listen command
func NewListenCommand() *cli.Command {
	return &cli.Command{
		Name:   "listen",
		Action: commands.Listen,
		Usage:  "Listens for new snapshots and uploads when found",
	}
}

func main() {
	app := &cli.App{
		Name:  "taraxa-snapshotter",
		Usage: "Taraxa Snapshotter uploads new snapshots to IPFS and publishes them to the Taraxa network.",
		Flags: []cli.Flag{
			ipfsUrlFlag,
			ipfsGatewayFlag,
			snapshotBaseDirFlag,
			verboseFlag,
		},
		Commands: []*cli.Command{
			NewListCommand(),
			NewListRemote(),
			NewPublishCommand(),
			NewRestoreCommand(),
			NewListenCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.WithError(err).Fatal("Failed to start application")
	}
}
