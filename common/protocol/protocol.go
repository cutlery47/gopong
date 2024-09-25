package protocol

type GameConfig struct {
	Side                  PlayerSide
	WindowWidth           int
	WindowHeight          int
	BallSize              int
	PlatformWidth         int
	PlatformHeight        int
	BallPosition          Vector
	LeftPlatformPosition  Vector
	RightPlatformPosition Vector
}

type PlayerSide string

var LeftSide PlayerSide = "left"
var RightSide PlayerSide = "right"

// Server Message
type ServerPacket struct {
	Status PlayerStatus
	State  ServerState
}

type PlayerStatus string

var SearchingStatus PlayerStatus = "searching"
var FoundStatus PlayerStatus = "found"
var PlayingStatus PlayerStatus = "playing"

// Client Message
type ClientPacket struct {
	Position Vector
}

var ClientAck string = "ACK"

type ServerState struct {
	LeftPosition  Vector
	RightPosition Vector
	BallPosition  Vector
}

// // position on the screen
type Vector struct {
	X float64
	Y float64
}
