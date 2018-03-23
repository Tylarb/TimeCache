/*
Package timeCache is a simple implementation of an interface and two structures
which define aa in-memory cache where entries have a timeout


Released under MIT license, copyright 2018 Tyler Ramer
*/
package timeCache

import "time"

// TimeCache interface defines necessary methods
type TimeCache interface {
	Contains(string) bool
	Push(string)
	Pop()
}

// SliceCache is an implementation of the time cache which uses two equal sized
// slices - one tracking keys, the other the time. kQueue[i] -> tQueue[i]
type SliceCache struct {
	kQueue  []string    // unique key in the cache
	tQueue  []time.Time // time in seconds since epoch when the entry was added to cache
	timeout time.Duration
	count   int // "low" mark - oldest entires in cache are here
}

// DictCache is an implementation of the time cache which uses a dict
type DictCache struct {
	entries map[string]time.Time // A map from a unique key to a time
	timeout time.Duration
	count   int
}

func (c *SliceCache) Contains(key string) bool {
	c.Pop()
	for _, k := range c.kQueue {
		if k == key {
			return true
		}
	}
	c.Push(key)
	return false
}

func (c *SliceCache) Push(key string) {
	if key == "" {
		return
	}
	c.kQueue = append(c.kQueue, key)
	c.tQueue = append(c.tQueue, time.Now())
	c.count++
}

func (c *SliceCache) Pop() {
	f := 0 // front of the stack for trimming
	for i, t := range c.tQueue {
		if time.Since(t) > c.timeout {
			f = i + 1
		}
	}
	c.kQueue = c.kQueue[f:]
	c.tQueue = c.tQueue[f:]
	c.count = c.count - f
}

func (c *DictCache) Contains(key string) bool {
	c.Pop()
	for k, _ := range c.entries {
		if k == key {
			return true
		}
	}
	c.Push(key)
	return false

}

func (c *DictCache) Push(key string) {
	if key == "" {
		return
	}
	c.entries[key] = time.Now()
	c.count++
}

func (c *DictCache) Pop() {
	for k, t := range c.entries {
		if time.Since(t) > c.timeout {
			delete(c.entries, k)
		}
	}
}
