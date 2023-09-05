package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func List(ctx *cli.Context) error {
	fmt.Println("List")
	return nil
}
