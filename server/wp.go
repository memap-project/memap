package server

import "net"

type Pool struct {
	tasks   chan net.Conn
	handler func(net.Conn)
}

func NewPool(handler func(net.Conn), capacity int) *Pool {
	pool := &Pool{
		tasks:   make(chan net.Conn, capacity),
		handler: handler,
	}
	return pool
}

func (p *Pool) Start(workerCount int) {
	for i := 0; i < workerCount; i++ {
		go p.worker()
	}
}

func (p *Pool) worker() {
	for task := range p.tasks {
		p.handler(task)
	}
}
