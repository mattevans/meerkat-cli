package cmd_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/mattevans/meerkat-cli/board"
	"github.com/mattevans/meerkat-cli/cmd"
	"github.com/mattevans/meerkat-cli/cmd/internal"
	"github.com/mattevans/meerkat-cli/types"
)

func TestExecuteMoveCmd(t *testing.T) {

	cases := []struct {
		label  string
		board  *board.Board
		input  *types.RobotCommand
		output string
	}{
		{
			label: "Executing MOVE before PLACE",
			board: internal.GetInActiveBoard(),
			input: &types.RobotCommand{
				RawCmd:     "MOVE",
				PrimaryCmd: "MOVE",
			},
			output: "✗ Error: Please PLACE your robot before attempting any other commands",
		},
		{
			label: "Executing MOVE our of bounds",
			board: internal.GetActiveBoard(0, 0, "SOUTH"),
			input: &types.RobotCommand{
				RawCmd:     "MOVE",
				PrimaryCmd: "MOVE",
			},
			output: "✗ Error: Ignoring MOVE command. Robot would become out-of-bounds if executed",
		},
		{
			label: "Executing MOVE while facing NORTH",
			board: internal.GetActiveBoard(0, 0, "NORTH"),
			input: &types.RobotCommand{
				RawCmd:     "MOVE",
				PrimaryCmd: "MOVE",
			},
			output: "✔ Success: MOVE",
		},
		{
			label: "Executing MOVE while facing EAST",
			board: internal.GetActiveBoard(0, 0, "EAST"),
			input: &types.RobotCommand{
				RawCmd:     "MOVE",
				PrimaryCmd: "MOVE",
			},
			output: "✔ Success: MOVE",
		},
		{
			label: "Executing MOVE while facing SOUTH",
			board: internal.GetActiveBoard(0, 1, "SOUTH"),
			input: &types.RobotCommand{
				RawCmd:     "MOVE",
				PrimaryCmd: "MOVE",
			},
			output: "✔ Success: MOVE",
		},
		{
			label: "Executing MOVE while facing WEST",
			board: internal.GetActiveBoard(1, 0, "WEST"),
			input: &types.RobotCommand{
				RawCmd:     "MOVE",
				PrimaryCmd: "MOVE",
			},
			output: "✔ Success: MOVE",
		},
	}

	for _, c := range cases {

		var buf bytes.Buffer
		handles := &types.RobotHandles{
			Writer: &buf,
			Err:    &buf,
		}

		cmd.ExecuteMoveCmd(handles, c.board, c.input)

		if strings.TrimSpace(buf.String()) != c.output {
			t.Errorf("%s: Expected `%s` but got `%s`", c.label, c.output, buf.String())
		}

	}
}
