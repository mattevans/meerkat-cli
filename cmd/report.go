package cmd

import (
	"errors"
	"fmt"

	"github.com/mattevans/meerkat-cli/board"
	"github.com/mattevans/meerkat-cli/output"
	"github.com/mattevans/meerkat-cli/types"
)

func init() {
	register(CMD_REPORT, ExecuteReportCmd)
}

// ExecuteReportCmd is called via a REPORT command and will output the robots
// current position/direction on the board as well as a ASCII representation of the board.
func ExecuteReportCmd(handles *types.RobotHandles, board *board.Board, instance *types.RobotCommand) {

	if !board.InUse {
		output.Error(handles.Err, errors.New("Please PLACE your robot before attempting any other commands"))
		return
	}

	instance.X = board.Current.X
	instance.Y = board.Current.Y
	instance.Facing = board.Current.Facing

	// Output confirmation.
	output.Info(handles.Writer, fmt.Sprintf("%s %v,%v,%v", CMD_REPORT, board.Current.X, board.Current.Y, board.Current.Facing))

	// We will also dump an ASCII table to visually show where
	// our robot is sitting on our board.
	output.ResultsTable(handles.Writer, board, instance, DirectionArrows[board.Current.Facing])
}
