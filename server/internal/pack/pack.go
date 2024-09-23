package pack

// Server Message
type ServerPacket struct {
	Status PlayerStatus
	Side   PlayerSide
	State  ServerState
}

// Client Message
type ClientPacket struct {
	Position Vector
}

type PlayerStatus string

var SearchingStatus PlayerStatus = "searching"
var FoundStatus PlayerStatus = "found"
var PlayingStatus PlayerStatus = "playing"

type PlayerSide string

var LeftSide PlayerSide = "left"
var RightSide PlayerSide = "right"

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
