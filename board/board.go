package board

import (
	"github.com/mattevans/meerkat-cli/types"
)

const (
	DEFAULT_BOARD_MAX_X = 5
	DEFAULT_BOARD_MAX_Y = 5
)

type Board struct {
	// Integers representing the bounds of our board.
	X, Y int

	// A flag to determine if our board is currently in-use (eg, a robot
	// has been placed).
	InUse bool

	// The last executed command given to our robot.
	Current *types.RobotCommand

	// A slice containing all historical commands given to our robot.
	History []*types.RobotCommand
}

// NewBoard intializes a new *Board using the given cooridnates as it's max
// bounds. If coordinates provided are invalid, it'll default to 5x5.
func NewBoard(x, y int) *Board {
	// Invalid coordinates provided, default to 5x5.
	if x < 0 || y < 0 {
		x = DEFAULT_BOARD_MAX_X
		y = DEFAULT_BOARD_MAX_Y
	}

	// Instantiate the board and return.
	board := &Board{X: x, Y: y}
	return board
}

// InBounds checks to see if the given coordinates fall within the boards bounds.
func (board *Board) InBounds(x int, y int) bool {
	return x < board.X && x >= 0 && y < board.Y && y >= 0
}
