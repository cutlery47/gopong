package game

// pipes for passing connections back & forth between Queue and Server
type inConnPipe chan<- connection
type outConnPipe <-chan connection

// the game itself
type Game struct {
	Server *Server
	Queue  *Queue
}

func New() *Game {
	// idk why 1024
	connPipe := make(chan connection, 1024)

	server := NewServer(connPipe)
	queue := NewQueue(connPipe)

	go queue.Accept()
	go queue.Run()

	return &Game{
		Server: server,
		Queue:  queue,
	}
}
