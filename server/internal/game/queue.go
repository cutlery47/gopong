package game

import (
	"errors"
	"gopong/server/internal/pack"
	"log"
	"sync"
	"time"
)

type connQueue []connection

func (cq connQueue) remove(i int) connQueue {
	return append(cq[:i], cq[i+1:]...)
}

func (cq connQueue) findAndRemove(conn connection) (connQueue, error) {
	for i, c := range cq {
		if c.conn == conn.conn {
			return append(cq[:i], cq[i+1:]...), nil
		}
	}
	return nil, errors.New("connection not found")
}

// queue for incoming websocket connections
type Queue struct {
	queue connQueue
	pipe  outConnPipe
	mu    *sync.Mutex
}

func NewQueue(pipe outConnPipe) *Queue {
	queue := []connection{}
	mu := &sync.Mutex{}

	return &Queue{
		queue: queue,
		pipe:  pipe,
		mu:    mu,
	}
}

// adding incoming connections to the queue
func (q *Queue) Accept() {
	for {
		conn := <-q.pipe
		// notifying client that it has been placed on the queue
		conn.SendStatus(pack.SearchingStatus)
		q.queue = append(q.queue, conn)
		go q.Manage(conn)
	}
}

// creating a new gaming session for each pair of connections in the queue
func (q *Queue) Run() {
	for {
		log.Println("Current connections:", q.queue)
		// waiting for connections
		if len(q.queue) < 2 {
			time.Sleep(200 * time.Millisecond)
		} else {
			q.mu.Lock()
			sesh := Session{
				player1: q.queue[0],
				player2: q.queue[1],
			}

			go sesh.handle()

			q.queue = q.queue[2:]
			q.mu.Unlock()
		}
	}
}

// listening for client Close() calls
func (q *Queue) Manage(conn connection) {
	for {
		// creating a temporary buffer
		pack := pack.ClientPacket{}
		err := conn.Read(&pack)
		if err != nil {
			log.Printf("Queue.Manage() from %v: %v", conn.conn.RemoteAddr(), err)
			// removing the Closed connection from the queue
			q.mu.Lock()
			q.queue, err = q.queue.findAndRemove(conn)
			if err != nil {
				log.Println(err)
			}
			q.mu.Unlock()
		}
		return
	}
}
