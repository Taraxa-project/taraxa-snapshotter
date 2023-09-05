package main

import (
	"os"

	"github.com/Taraxa-project/taraxa-snapshotter/internal/commands"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var (
	ipfsUrlFlag = cli.StringFlag{
		Name:    "ipfs-url",
		Value:   "https://ipfs.infura.io:5001",
		Usage:   "URL of the IPFS node",
		Aliases: []string{"t"},
		EnvVars: []string{"IPFS_URL"},
	}
)

func main() {
	app := &cli.App{
		Name:  "taraxa-snapshotter",
		Usage: "Taraxa Snapshotter uploads new snapshots to IPFS and publishes them to the Taraxa network.",
		Flags: []cli.Flag{
			&ipfsUrlFlag,
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

// NewListCommand is the factory method for the list command
func NewListCommand() *cli.Command {
	return &cli.Command{
		Name:   "list",
		Action: commands.List,
	}
}

// NewListRemote is the factory method for the list-remote command
func NewListRemote() *cli.Command {
	return &cli.Command{
		Name:   "list-remote",
		Action: commands.ListRemote,
	}
}

// NewPublishCommand is the factory method for the publish command
func NewPublishCommand() *cli.Command {
	return &cli.Command{
		Name:   "publish",
		Action: commands.Publish,
	}
}

// NewRestoreCommand is the factory method for the restore command
func NewRestoreCommand() *cli.Command {
	return &cli.Command{
		Name:   "restore",
		Action: commands.Restore,
	}
}

// NewListenCommand is the factory method for the listen command
func NewListenCommand() *cli.Command {
	return &cli.Command{
		Name:   "listen",
		Action: commands.Listen,
	}
}
