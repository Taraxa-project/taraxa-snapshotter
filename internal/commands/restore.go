package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Restore restores a remote snapshot from IPFS
func Restore(ctx *cli.Context) error {
	fmt.Println("Restore")
	return nil
}
