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

func TestExecutePLACECmd(t *testing.T) {

	cases := []struct {
		label  string
		board  *board.Board
		input  *types.RobotCommand
		output string
	}{
		{
			label: "Executing PLACE at 0,0",
			board: internal.GetActiveBoard(0, 0, "SOUTH"),
			input: &types.RobotCommand{
				RawCmd:     "PLACE 0,0,NORTH",
				PrimaryCmd: "PLACE",
			},
			output: "✔ Success: PLACE 0,0,NORTH",
		},
		{
			label: "Executing PLACE at 4,4",
			board: internal.GetActiveBoard(0, 0, "SOUTH"),
			input: &types.RobotCommand{
				RawCmd:     "PLACE 4,4,SOUTH",
				PrimaryCmd: "PLACE",
			},
			output: "✔ Success: PLACE 4,4,SOUTH",
		},
		{
			label: "Executing PLACE with invalid bounds",
			board: internal.GetActiveBoard(0, 0, "SOUTH"),
			input: &types.RobotCommand{
				RawCmd:     "PLACE 5,7,SOUTH",
				PrimaryCmd: "PLACE",
			},
			output: "✗ Error: Cannot place robot at 5,7 - that's out of bounds",
		},
		{
			label: "Executing PLACE with invalid number of coordinates",
			board: internal.GetActiveBoard(0, 0, "SOUTH"),
			input: &types.RobotCommand{
				RawCmd:     "PLACE 3,4,3,2,SOUTH",
				PrimaryCmd: "PLACE",
			},
			output: "✗ Error: Invalid number of arguments for `PLACE` command, please use `PLACE X,Y,FACING`",
		},
		{
			label: "Executing PLACE with invalid X coordinate data type",
			board: internal.GetActiveBoard(0, 0, "SOUTH"),
			input: &types.RobotCommand{
				RawCmd:     "PLACE HELLO,3,SOUTH",
				PrimaryCmd: "PLACE",
			},
			output: "✗ Error: Coordinate `x` provided must be a number",
		},
		{
			label: "Executing PLACE with invalid Y coordinate data type",
			board: internal.GetActiveBoard(0, 0, "SOUTH"),
			input: &types.RobotCommand{
				RawCmd:     "PLACE 0,HELLO,SOUTH",
				PrimaryCmd: "PLACE",
			},
			output: "✗ Error: Coordinate `y` provided must be a number",
		},
		{
			label: "Executing PLACE with invalid facing value",
			board: internal.GetActiveBoard(0, 0, "SOUTH"),
			input: &types.RobotCommand{
				RawCmd:     "PLACE 0,0,CENTER",
				PrimaryCmd: "PLACE",
			},
			output: "✗ Error: Invalid `facing` value provided, please use NORTH, EAST, SOUTH, WEST",
		},
	}

	for _, c := range cases {

		var buf bytes.Buffer
		handles := &types.RobotHandles{
			Writer: &buf,
			Err:    &buf,
		}

		cmd.ExecutePlaceCmd(handles, c.board, c.input)

		if strings.TrimSpace(buf.String()) != c.output {
			t.Errorf("%s: Expected `%s` but got `%s`", c.label, c.output, buf.String())
		}

	}
}
