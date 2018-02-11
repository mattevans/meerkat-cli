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

func TestExecuteLeftCmd(t *testing.T) {

	cases := []struct {
		label  string
		board  *board.Board
		input  *types.RobotCommand
		output string
	}{
		{
			label: "Executing LEFT before PLACE",
			board: internal.GetInActiveBoard(),
			input: &types.RobotCommand{
				RawCmd:     "LEFT",
				PrimaryCmd: "LEFT",
			},
			output: "✗ Error: Please PLACE your robot before attempting any other commands",
		},
		{
			label: "Executing LEFT while facing NORTH",
			board: internal.GetActiveBoard(0, 0, "NORTH"),
			input: &types.RobotCommand{
				RawCmd:     "PLACE 0,0,NORTH",
				PrimaryCmd: "PLACE",
			},
			output: "✔ Success: LEFT",
		},
		{
			label: "Executing LEFT while facing EAST",
			board: internal.GetActiveBoard(0, 0, "EAST"),
			input: &types.RobotCommand{
				RawCmd:     "PLACE 0,0,EAST",
				PrimaryCmd: "PLACE",
			},
			output: "✔ Success: LEFT",
		},
		{
			label: "Executing LEFT while facing SOUTH",
			board: internal.GetActiveBoard(0, 0, "SOUTH"),
			input: &types.RobotCommand{
				RawCmd:     "PLACE 0,0,SOUTH",
				PrimaryCmd: "PLACE",
			},
			output: "✔ Success: LEFT",
		},
		{
			label: "Executing LEFT while facing WEST",
			board: internal.GetActiveBoard(0, 0, "WEST"),
			input: &types.RobotCommand{
				RawCmd:     "PLACE 0,0,WEST",
				PrimaryCmd: "PLACE",
			},
			output: "✔ Success: LEFT",
		},
	}

	for _, c := range cases {

		var buf bytes.Buffer
		handles := &types.RobotHandles{
			Writer: &buf,
			Err:    &buf,
		}

		cmd.ExecuteLeftCmd(handles, c.board, c.input)

		if strings.TrimSpace(buf.String()) != c.output {
			t.Errorf("%s: Expected `%s` but got `%s`", c.label, c.output, buf.String())
		}

	}
}
