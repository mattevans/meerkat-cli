package board_test

import (
	"testing"

	"github.com/mattevans/meerkat-cli/board"
)

func TestNewBoard(t *testing.T) {

	cases := []struct {
		label string
		x     int
		y     int
	}{
		{
			label: "Creating new board without with invalid bounds",
			x:     -2,
			y:     -3,
		},
	}

	for _, c := range cases {

		target := board.NewBoard(c.x, c.y)

		if target.X != board.DEFAULT_BOARD_MAX_X {
			t.Errorf("%s: Expected new board with X `%v` but got `%v`", c.label, board.DEFAULT_BOARD_MAX_X, target.X)
		}
		if target.X != board.DEFAULT_BOARD_MAX_Y {
			t.Errorf("%s: Expected new board with Y `%v` but got `%v`", c.label, board.DEFAULT_BOARD_MAX_Y, target.X)
		}

	}
}

func TestBoardInBounds(t *testing.T) {

	cases := []struct {
		label string
		pass  bool
		x     int
		y     int
	}{
		{
			label: "Checking coordinates in bounds fails",
			x:     5,
			y:     6,
			pass:  false,
		},
		{
			label: "Checking coordinates in bounds succeeds",
			x:     2,
			y:     3,
			pass:  true,
		},
	}

	for _, c := range cases {

		target := board.NewBoard(board.DEFAULT_BOARD_MAX_X, board.DEFAULT_BOARD_MAX_Y)

		if !c.pass {
			if target.InBounds(c.x, c.y) {
				t.Errorf("%s: Expected coordinates `%v,%v` to be out of board bounds", c.label, c.x, c.y)
			}
		}

		if c.pass {
			if !target.InBounds(c.x, c.y) {
				t.Errorf("%s: Expected coordinates `%v,%v` to be within board bounds", c.label, c.x, c.y)
			}
		}
	}
}
