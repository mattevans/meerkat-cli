package main

import (
	"fmt"
	"os"

	"github.com/mattevans/meerkat-cli/board"
	"github.com/mattevans/meerkat-cli/robot"
	"github.com/mattevans/meerkat-cli/types"
	"github.com/urfave/cli"
)

const (
	CLI_NAME    = "meerkat-cli"
	CLI_VERSION = "0.0.1"
)

func main() {

	// Boot our CLI app.
	meerkat := cli.NewApp()
	meerkat.Name = CLI_NAME
	meerkat.Version = CLI_VERSION
	meerkat.Action = func(ctx *cli.Context) {
		// Initiate a robot.
		robot := &robot.Robot{
			Board: board.NewBoard(
				board.DEFAULT_BOARD_MAX_X,
				board.DEFAULT_BOARD_MAX_Y,
			),
			Handles: &types.RobotHandles{
				Reader: os.Stdin,
				Writer: os.Stdout,
				Err:    os.Stderr,
			},
		}
		robot.Load()
	}

	// Run our app.
	if err := meerkat.Run(os.Args); err != nil {
		fmt.Printf("Error executing %v: %v\n", CLI_NAME, err)
		os.Exit(1)
	}
}
