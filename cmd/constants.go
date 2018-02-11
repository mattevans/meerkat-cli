package cmd

const (
	// Define the different directions we support.
	DIR_NORTH = "NORTH"
	DIR_EAST  = "EAST"
	DIR_SOUTH = "SOUTH"
	DIR_WEST  = "WEST"

	// Define the different commands we support.
	CMD_PLACE  = "PLACE"
	CMD_MOVE   = "MOVE"
	CMD_LEFT   = "LEFT"
	CMD_RIGHT  = "RIGHT"
	CMD_REPORT = "REPORT"
	CMD_EXIT   = "EXIT"
)

// CommandsWhitelist is a slice containing all commands supported.
var CommandsWhitelist = []string{
	CMD_PLACE,
	CMD_MOVE,
	CMD_LEFT,
	CMD_RIGHT,
	CMD_REPORT,
	CMD_EXIT,
}

// DirectionsWhitelist is a slice containing all directions supported.
var DirectionsWhitelist = []string{
	DIR_NORTH,
	DIR_EAST,
	DIR_SOUTH,
	DIR_WEST,
}

// DirectionArrows is a map of directional arrows keyed by direction.
// Used when printing our results via a REPORT command.
var DirectionArrows = map[string]string{
	DIR_NORTH: "↑",
	DIR_EAST:  "→",
	DIR_SOUTH: "↓",
	DIR_WEST:  "←",
}
