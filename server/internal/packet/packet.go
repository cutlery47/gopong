package packet

const (
	StateMatchmaking = "Matchmaking"
	StatePlaying     = "Playing"
)

type Packet struct {
	State string

	Left  VectorPackage
	Right VectorPackage
	Ball  VectorPackage
}

type VectorPackage struct {
	X float64
	Y float64
}
