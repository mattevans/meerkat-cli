package cmd

import (
	"fmt"
	"strings"

	"github.com/mattevans/meerkat-cli/board"
	"github.com/mattevans/meerkat-cli/types"
)

type robotCmdFn func(*types.RobotHandles, *board.Board, *types.RobotCommand)

var cmdRegistry = map[string]robotCmdFn{}

func register(name string, fn robotCmdFn) {
	cmdRegistry[name] = fn
}

var directionsMap = map[string]bool{}
var commandsMap = map[string]bool{}

func init() {
	// Map our whitelisted commands.
	for _, cmd := range CommandsWhitelist {
		commandsMap[cmd] = true
	}

	// Map our whitelisted directions.
	for _, dir := range DirectionsWhitelist {
		directionsMap[dir] = true
	}
}

// Delegate validates the user input, determines and retrieves the command
// to be executed from our registry, and then executes it.
func Delegate(handles *types.RobotHandles, board *board.Board, input string) (*types.RobotCommand, error) {
	// Is the command something we recognise?
	command := getCommand(input)
	if command == nil {
		return nil, fmt.Errorf("Invalid command, please use %s", strings.Join(CommandsWhitelist, ", "))
	}

	// Grab command from our registry.
	fn := cmdRegistry[strings.ToUpper(*command)]

	// Build a Robot instance.
	instance := &types.RobotCommand{
		RawCmd:     input,
		PrimaryCmd: *command,
	}

	// Call the respective command.
	fn(handles, board, instance)

	return instance, nil
}

// getCommand attempts to match the given command to our list of supported commands.
func getCommand(command string) *string {
	// Explode by space delimeter. Target length will never be < 0  due to
	// previous checks run in Load().
	target := strings.Split(command, " ")

	if commandsMap[strings.ToUpper(target[0])] {
		return &target[0]
	}

	return nil
}
