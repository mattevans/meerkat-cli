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

func TestDelegate(t *testing.T) {

	cases := []struct {
		label  string
		board  *board.Board
		input  string
		output string
	}{
		{
			label:  "Executing Delegate with invalid command",
			board:  internal.GetInActiveBoard(),
			input:  "INVALID COMMAND 1,2,3,4",
			output: "Invalid command, please use PLACE, MOVE, LEFT, RIGHT, REPORT, EXIT",
		},
		{
			label:  "Executing Delegate with valid command",
			board:  internal.GetInActiveBoard(),
			input:  "PLACE 1,3,NORTH",
			output: "Invalid command, please use PLACE, MOVE, LEFT, RIGHT, REPORT, EXIT",
		},
	}

	for _, c := range cases {

		var buf bytes.Buffer
		handles := &types.RobotHandles{
			Writer: &buf,
			Err:    &buf,
		}

		command, err := cmd.Delegate(handles, c.board, c.input)
		if err != nil {
			if strings.TrimSpace(err.Error()) != c.output {
				t.Errorf("%s: Expected `%s` but got `%s`", c.label, c.output, err.Error())
			}
		}

		if command != nil {
			if command.RawCmd != c.input {
				t.Errorf("%s: Expected `%s` but got `%s`", c.label, c.input, command.RawCmd)
			}
		}
	}
}
