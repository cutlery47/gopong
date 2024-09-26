package queue

import (
	"gopong/server/internal/game/session"
	"log"
	"sync"
	"time"

	"github.com/cutlery47/gopong/common/conn"
)

// queue for incoming websocket connections
type Queue struct {
	// the queue itself
	queue connQueue
	mu    *sync.Mutex

	// channel for receiving connections
	connPipe <-chan conn.Connection
	// exit channel for Close() calls listener
	exitListen chan byte
}

func New(pipe <-chan conn.Connection) *Queue {
	return &Queue{
		queue:      []conn.Connection{},
		mu:         &sync.Mutex{},
		connPipe:   pipe,
		exitListen: make(chan byte),
	}
}

// adding incoming connections to the queue
func (q *Queue) Accept() {
	for {
		conn := <-q.connPipe
		go q.listenForRemoval(conn)
		q.queue = append(q.queue, conn)
	}
}

// creating a new gaming session for each pair of connections in the queue
func (q *Queue) Run() {
	for {
		// log.Println("Current connections:", q.queue)
		if len(q.queue) < 2 {
			// waiting for connections
			time.Sleep(200 * time.Millisecond)
		} else {
			// telling listener to stop handling
			q.exitListen <- 1
			// creating a gaming session
			q.mu.Lock()
			go session.Init(q.queue[0], q.queue[1])
			q.queue = q.queue[2:]
			q.mu.Unlock()
		}
	}
}

// removing connections if any such signal received
func (q *Queue) listenForRemoval(conn conn.Connection) {
	select {
	case <-conn.RemoveConnPipe():
		q.mu.Lock()
		q.queue.findAndRemove(conn)
		q.mu.Unlock()
	case <-q.exitListen:
	}
}

type connQueue []conn.Connection

func (cq *connQueue) findAndRemove(conn conn.Connection) {
	for i, c := range *cq {
		if c.RemoteAddr() == conn.RemoteAddr() {
			*cq = append((*cq)[:i], (*cq)[i+1:]...)
		}
	}
	log.Println("connection not found")
}
