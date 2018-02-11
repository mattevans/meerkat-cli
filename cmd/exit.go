package cmd

import (
	"os"

	"github.com/mattevans/meerkat-cli/board"
	"github.com/mattevans/meerkat-cli/output"
	"github.com/mattevans/meerkat-cli/types"
)

func init() {
	register(CMD_EXIT, ExecuteExitCmd)
}

// ExecuteExitCmd is called upon EXIT command and simply exits the process for us.
func ExecuteExitCmd(handles *types.RobotHandles, board *board.Board, instance *types.RobotCommand) {
	// Output confirmation.
	output.Info(handles.Writer, "Come back soon!")
	os.Exit(0)
}
