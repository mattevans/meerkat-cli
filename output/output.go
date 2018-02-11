package output

import (
	"fmt"
	"io"

	"github.com/fatih/color"
	"github.com/mattevans/meerkat-cli/board"
	"github.com/mattevans/meerkat-cli/types"
	"github.com/olekukonko/tablewriter"
)

var red = color.New(color.FgHiRed).SprintFunc()
var green = color.New(color.FgHiGreen).SprintFunc()
var blue = color.New(color.FgHiBlue).SprintFunc()

// Info will display the input via the given writer in
// our information format.
func Info(writer io.Writer, input string) {
	fmt.Fprintln(writer, fmt.Sprintf("%v %v", blue("▸"), input))
}

// Success will display the input via the given writer in
// our success format.
func Success(writer io.Writer, input string) {
	fmt.Fprintln(writer, fmt.Sprintf("%v %v", green("✔ Success:"), input))
}

// Error will display the input via the given writer in
// our error format.
func Error(writer io.Writer, err error) {
	fmt.Fprintln(writer, fmt.Sprintf("%v %v", red("✗ Error:"), err))
}

// Welcome will display a welcome/boot message to the user.
func Welcome(writer io.Writer) {
	header := `
                 _______
               _/       \_
              / |       | \
             /  |__   __|  \
            |__/((o| |o))\__|
            |      | |      |
            |\     |_|     /|
            | \           / |
             \| /  ___  \ |/
              \ | / _ \ | /
               \_________/
                _|_____|_
           ____|_________|____
          /                   \
           Toy Robot Simulator
`
	fmt.Fprintln(writer, blue(header))
	Info(writer, "Enter your command to proceed:")
}

// ResultsTable will display the current position/direction of our robot
// on the board via an ASCII table. Think of this as a visual representation
// of our robots position.
func ResultsTable(writer io.Writer, board *board.Board, instance *types.RobotCommand, arrow string) {
	data := [][]string{}
	for yi := 0; yi < board.Y; yi++ {
		rows := []string{}
		for xi := 0; xi < board.X; xi++ {
			if instance.X == xi && instance.Y == yi {
				rows = append(rows, green(arrow))
			} else {
				rows = append(rows, "  ")
			}
		}
		data = append(data, rows)
	}

	table := tablewriter.NewWriter(writer)
	table.AppendBulk(reverseTableData(data))
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetRowLine(true)
	table.Render()
}

func reverseTableData(input [][]string) [][]string {
	if len(input) == 0 {
		return input
	}
	return append(reverseTableData(input[1:]), input[0])
}
