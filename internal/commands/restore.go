package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Restore(ctx *cli.Context) error {
	fmt.Println("Restore")
	return nil
}
