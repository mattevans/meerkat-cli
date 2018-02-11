package internal

import (
	"fmt"

	"github.com/mattevans/meerkat-cli/board"
	"github.com/mattevans/meerkat-cli/types"
)

// GetActiveBoard will return an instantiated *board.Board for testing.
func GetActiveBoard(x, y int, facing string) *board.Board {
	return &board.Board{
		X:     5,
		Y:     5,
		InUse: true,
		Current: &types.RobotCommand{
			X:          x,
			Y:          y,
			Facing:     facing,
			RawCmd:     fmt.Sprintf("PLACE %v,%v,%s", x, y, facing),
			PrimaryCmd: "PLACE",
		},
	}
}

// GetInActiveBoard will return a inactive *board.Board for testing.
func GetInActiveBoard() *board.Board {
	return &board.Board{
		X:     5,
		Y:     5,
		InUse: false,
	}
}
