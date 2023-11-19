package sync

import "sync"

type Counter struct{
	mx sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.value++
}

func (c *Counter) Value() int{
	return c.value
}