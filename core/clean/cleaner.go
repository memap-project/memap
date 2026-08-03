package clean

import (
	"context"
	"time"
)

type Cleaner struct {
	interval  uint64 // in seconds
	cleanFunc func()
	ctx       context.Context
	cancel    context.CancelFunc
}

func NewCleaner(parentCtx context.Context, interval uint64, cleanFunc func()) *Cleaner {
	if parentCtx == nil {
		parentCtx = context.Background()
	}
	ctx, cancel := context.WithCancel(parentCtx)
	return &Cleaner{
		interval:  interval,
		cleanFunc: cleanFunc,
		ctx:       ctx,
		cancel:    cancel,
	}
}

func (c *Cleaner) Start() {
	go c.run()
}

func (c *Cleaner) Stop() {
	c.cancel()
}

func (c *Cleaner) run() {
	ticker := time.NewTicker(time.Duration(c.interval) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.cleanFunc()
		case <-c.ctx.Done():
			return
		}
	}
}
