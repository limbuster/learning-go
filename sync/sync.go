package main

import "sync"

// Counter to store the count
type Counter struct {
	count int
	mu    sync.Mutex
}

// Inc increments the count by one
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// Value returns the current count
func (c *Counter) Value() int {
	return c.count
}
