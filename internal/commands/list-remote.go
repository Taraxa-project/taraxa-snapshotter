package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// ListRemote lists published snapshot hashes
func ListRemote(ctx *cli.Context) error {
	fmt.Println("ListRemote")
	return nil
}
