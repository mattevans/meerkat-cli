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

func TestExecuteReportCmd(t *testing.T) {

	cases := []struct {
		label  string
		board  *board.Board
		input  *types.RobotCommand
		output string
	}{
		{
			label: "Executing REPORT before PLACE",
			board: internal.GetInActiveBoard(),
			input: &types.RobotCommand{
				RawCmd:     "REPORT",
				PrimaryCmd: "REPORT",
			},
			output: "✗ Error: Please PLACE your robot before attempting any other commands",
		},
		{
			label: "Executing REPORT on 3,2,NORTH",
			board: internal.GetActiveBoard(3, 2, "NORTH"),
			input: &types.RobotCommand{
				RawCmd:     "REPORT",
				PrimaryCmd: "REPORT",
			},
			output: `▸ REPORT 3,2,NORTH
+----+----+----+----+----+
|    |    |    |    |    |
+----+----+----+----+----+
|    |    |    |    |    |
+----+----+----+----+----+
|    |    |    | ↑  |    |
+----+----+----+----+----+
|    |    |    |    |    |
+----+----+----+----+----+
|    |    |    |    |    |
+----+----+----+----+----+`,
		},
		{
			label: "Executing REPORT on 1,4,SOUTH",
			board: internal.GetActiveBoard(1, 4, "SOUTH"),
			input: &types.RobotCommand{
				RawCmd:     "REPORT",
				PrimaryCmd: "REPORT",
			},
			output: `▸ REPORT 1,4,SOUTH
+----+----+----+----+----+
|    | ↓  |    |    |    |
+----+----+----+----+----+
|    |    |    |    |    |
+----+----+----+----+----+
|    |    |    |    |    |
+----+----+----+----+----+
|    |    |    |    |    |
+----+----+----+----+----+
|    |    |    |    |    |
+----+----+----+----+----+`,
		},
	}

	for _, c := range cases {

		var buf bytes.Buffer
		handles := &types.RobotHandles{
			Writer: &buf,
			Err:    &buf,
		}

		cmd.ExecuteReportCmd(handles, c.board, c.input)

		if strings.TrimSpace(buf.String()) != c.output {
			t.Errorf("%s: Expected `%s` but got `%s`", c.label, c.output, buf.String())
		}

	}
}
