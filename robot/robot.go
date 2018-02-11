package robot

import (
	"bufio"
	"strings"

	"github.com/mattevans/meerkat-cli/board"
	"github.com/mattevans/meerkat-cli/cmd"
	"github.com/mattevans/meerkat-cli/output"
	"github.com/mattevans/meerkat-cli/types"
)

// Robot represents our robot in play.
type Robot struct {
	// An instance of *Board that our robot is playing on.
	Board *board.Board

	// An instance of our reader/writers for the session.
	Handles *types.RobotHandles
}

// Load will read our user input and runs given commands.
func (robot *Robot) Load() {

	// Initalize our reader.
	reader := bufio.NewReader(robot.Handles.Reader)

	// Dump a welcome message for the user.
	output.Welcome(robot.Handles.Writer)

	for {
		// Parse our buffer by new-line delimeter.
		tail, _ := reader.ReadString('\n')

		// Trim new-line from value, giving us out command input.
		command := strings.TrimSuffix(tail, "\n")

		// Run the given command.
		err := robot.run(command)
		if err != nil {
			output.Error(robot.Handles.Err, err)
			continue
		}
	}
}

// run will check the given command is valid and then hand-off to our
// command delagtor for excution.
func (robot *Robot) run(command string) error {

	// Call our command delegate which will handle executing the move.
	move, err := cmd.Delegate(robot.Handles, robot.Board, command)
	if err != nil {
		return err
	}

	// Move our existing *RobotCommand into our historical slice
	// before assiging the most recent *RobotCommand to Robot.Current
	if robot.Board.Current != nil {
		robot.Board.History = append(robot.Board.History, robot.Board.Current)
	}
	robot.Board.Current = move

	return err
}
