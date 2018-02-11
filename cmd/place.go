package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/mattevans/meerkat-cli/board"
	"github.com/mattevans/meerkat-cli/output"
	"github.com/mattevans/meerkat-cli/types"
)

func init() {
	register(CMD_PLACE, ExecutePlaceCmd)
}

// ExecutePlaceCmd is called via a PLACE command, validates the given command and arguments, runs a
// bounds check for our board and then outputs confirmation of placement.
func ExecutePlaceCmd(handles *types.RobotHandles, board *board.Board, instance *types.RobotCommand) {
	// Ensure arguments provided with PLACE command are valid.
	err := validatePlaceArguments(instance)
	if err != nil {
		output.Error(handles.Err, err)
		return
	}

	// Check the PLACE arguments provided are within bounds.
	if !board.InBounds(instance.X, instance.Y) {
		output.Error(handles.Err, fmt.Errorf("Cannot place robot at %v,%v - that's out of bounds", instance.X, instance.Y))
		return
	}

	// Ensure we flag our board as 'in use'.
	board.InUse = true

	// Output confirmation.
	output.Success(handles.Writer, fmt.Sprintf("%s %v,%v,%v", CMD_PLACE, instance.X, instance.Y, instance.Facing))
}

// validatePlaceArguments ensures arguments passed with the PLACE command
// are valid in number, type and value.
func validatePlaceArguments(instance *types.RobotCommand) error {
	var err error

	// Pop off our primary command, so we're left with coordinate arguments.
	coordinates := strings.TrimSpace(
		strings.Replace(instance.RawCmd, instance.PrimaryCmd, "", -1),
	)

	// Check we have the right amount of coordinate arguments.
	coordinatesSlice := strings.Split(coordinates, ",")
	if len(coordinatesSlice) != 3 {
		return fmt.Errorf("Invalid number of arguments for `%s` command, please use `%s X,Y,FACING`", instance.PrimaryCmd, CMD_PLACE)
	}

	// Ensure X coordinate is of a valid type.
	coodX, err := strconv.Atoi(coordinatesSlice[0])
	if err != nil {
		return errors.New("Coordinate `x` provided must be a number")
	}
	instance.X = coodX

	// Ensure Y coordinate is of a valid type.
	coodY, err := strconv.Atoi(coordinatesSlice[1])
	if err != nil {
		return errors.New("Coordinate `y` provided must be a number")
	}
	instance.Y = coodY

	// Ensure FACING argument provided is valid.
	if !directionsMap[strings.ToUpper(coordinatesSlice[2])] {
		return fmt.Errorf("Invalid `facing` value provided, please use %s", strings.Join(DirectionsWhitelist, ", "))
	}
	instance.Facing = strings.ToUpper(coordinatesSlice[2])

	return err
}
