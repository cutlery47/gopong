package packet

// client-side game configuration
type ConfigRequestPacket struct {
	WindowHeight   int
	WindowWidth    int
	PlatformWidth  int
	PlatformHeight int
	BallSize       int
}

const ConfigAccept = "ConfOkay"
const ConfigDecline = "ConfBad"

// server-side configuration acknowledgement
type ConfigResponsePacket struct {
	Response string
	Left     string
	Right    string
}

// client-side game state
type PlayerStatePacket struct {
	Position Vector
}

// server-side game state
type ServerStatePacket struct {
	LeftPosition  Vector
	RightPosition Vector
	BallPosition  Vector
}

// position on the screen
type Vector struct {
	X float64
	Y float64
}
