package commands

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/Taraxa-project/taraxa-snapshotter/internal/upload/ipfs"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func Publish(ctx *cli.Context) error {
	fmt.Println("Publish")

	u, err := url.Parse(ctx.String("ipfs-url"))
	if err != nil {
		log.WithError(err).Fatal("Could not parse IPFS URL")
	}

	ipfsHost := u.Scheme + "://" + u.Host
	ipfsUsername := u.User.Username()
	ipfsPassword, _ := u.User.Password()

	log.WithFields(log.Fields{
		"u":        ipfsUsername,
		"p":        ipfsPassword,
		"cleanUrl": ipfsHost,
	}).Info("Starting IPFS client")

	ipfsClient := ipfs.NewIPFSClient(ipfsHost, ipfsUsername, ipfsPassword)

	cid, err := ipfsClient.Add(strings.NewReader("Infura IPFS - Getting started demo."))
	if err != nil {
		log.WithError(err).Fatal("Failed to upload file to IPFS")
	}

	log.Info("IPFS client started: ", cid)

	return nil
}
