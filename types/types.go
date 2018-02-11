package types

import (
	"io"
)

// RobotCommand represents a command issued by the user.
type RobotCommand struct {
	X, Y       int
	Facing     string
	RawCmd     string
	PrimaryCmd string
}

// RobotHandles holds our reader/writers for the session.
type RobotHandles struct {
	Reader io.Reader
	Writer io.Writer
	Err    io.Writer
}
