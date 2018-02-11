package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mattevans/meerkat-cli/board"
	"github.com/mattevans/meerkat-cli/output"
	"github.com/mattevans/meerkat-cli/types"
)

func init() {
	register(CMD_MOVE, ExecuteMoveCmd)
}

// ExecuteMoveCmd is called via a MOVE command, determines the robots direction and
// increments/decrements the robot position accordingly before outputting confirmation.
func ExecuteMoveCmd(handles *types.RobotHandles, board *board.Board, instance *types.RobotCommand) {

	if !board.InUse {
		output.Error(handles.Err, errors.New("Please PLACE your robot before attempting any other commands"))
		return
	}

	var x, y int

	// Rotate our robot appropriately.
	switch strings.ToUpper(board.Current.Facing) {
	case DIR_NORTH:
		x = board.Current.X
		y = board.Current.Y + 1
		break
	case DIR_EAST:
		x = board.Current.X + 1
		y = board.Current.Y
		break
	case DIR_SOUTH:
		x = board.Current.X
		y = board.Current.Y - 1
		break
	case DIR_WEST:
		x = board.Current.X - 1
		y = board.Current.Y
		break
	}

	// Check to see the move isn't going to push our robot out of bounds.
	if !board.InBounds(x, y) {
		// Set out instance to use existing coordinate arguments.
		instance.X = board.Current.X
		instance.Y = board.Current.Y
		instance.Facing = board.Current.Facing

		// Throw an error explaining we're ignoring their command, as it would
		// push the robot off our board.
		output.Error(handles.Err, fmt.Errorf("Ignoring %s command. Robot would become out-of-bounds if executed", CMD_MOVE))
		return
	}

	instance.X = x
	instance.Y = y
	instance.Facing = board.Current.Facing

	// Output confirmation.
	output.Success(handles.Writer, CMD_MOVE)
}
