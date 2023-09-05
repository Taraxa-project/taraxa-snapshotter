package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Listen(ctx *cli.Context) error {
	fmt.Println("Listen")
	return nil
}
