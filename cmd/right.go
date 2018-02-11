package cmd

import (
	"errors"
	"strings"

	"github.com/mattevans/meerkat-cli/board"
	"github.com/mattevans/meerkat-cli/output"
	"github.com/mattevans/meerkat-cli/types"
)

func init() {
	register(CMD_RIGHT, ExecuteRightCmd)
}

// ExecuteRightCmd is called via a RIGHT command, determines the robots current facing
// position and updates it accordingly before outputting confirmation.
func ExecuteRightCmd(handles *types.RobotHandles, board *board.Board, instance *types.RobotCommand) {

	if !board.InUse {
		output.Error(handles.Err, errors.New("Please PLACE your robot before attempting any other commands"))
		return
	}

	// Rotate our robot appropriately.
	switch strings.ToUpper(board.Current.Facing) {
	case DIR_NORTH:
		instance.Facing = DIR_EAST
	case DIR_EAST:
		instance.Facing = DIR_SOUTH
	case DIR_SOUTH:
		instance.Facing = DIR_WEST
	case DIR_WEST:
		instance.Facing = DIR_NORTH
	}

	instance.X = board.Current.X
	instance.Y = board.Current.Y

	// Output confirmation.
	output.Success(handles.Writer, CMD_RIGHT)
}
