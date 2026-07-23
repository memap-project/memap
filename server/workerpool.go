package server

import "net"

type WorkerPool struct {
	tasks chan net.Conn
}

func NewWorkerPool(size int) *WorkerPool {
	pool := &WorkerPool{tasks: make(chan net.Conn, 100)}
	for i := 0; i < size; i++ {
		go pool.worker()
	}
	return pool
}

func (p *WorkerPool) worker() {
	for conn := range p.tasks {
		handleConnection(conn)
	}
}
