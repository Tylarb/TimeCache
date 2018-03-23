/*
Package timeCache is a simple implementation of an interface and two structures
which define aa in-memory cache where entries have a timeout


Released under MIT license, copyright 2018 Tyler Ramer
*/
package timeCache

import "time"

// TimeCache interface defines necessary methods
type TimeCache interface {
	Contains(string) (bool, time.Time)
	Push(string)
	Pop()
}

// SliceCache is an implementation of the time cache which uses two equal sized
// slices - one tracking keys, the other the time. keys[i] -> times[i]
type SliceCache struct {
	keys    []string    // unique key in the cache
	times   []time.Time // time in seconds since epoch when the entry was added to cache
	timeout time.Duration
	count   int // "low" mark - oldest entires in cache are here
}

// DictCache is an implementation of the time cache which uses a dict
type DictCache struct {
	entries map[string]time.Time // A map from a unique key to a time
	timeout time.Duration
	count   int
}

func (c *SliceCache) Contains(key string) (bool, time.Time) {
	c.Pop()
	for i, k := range c.keys {
		if k == key {
			return true, c.times[i]
		}
	}
	c.Push(key)
	return false, time.Now()
}

func (c *SliceCache) Push(key string) {
	if key == "" {
		return
	}
	c.keys = append(c.keys, key)
	c.times = append(c.times, time.Now())
	c.count++
}

func (c *SliceCache) Pop() {
	f := 0 // front of the stack for trimming
	for i, t := range c.times {
		if time.Since(t) > c.timeout {
			f = i + 1
		}
	}
	c.keys = c.keys[f:]
	c.times = c.times[f:]
	c.count = c.count - f
}

func (c *DictCache) Contains(key string) (bool, time.Time) {
	c.Pop()
	for k := range c.entries {
		a
		if k == key {
			return true, c.entries[k]
		}
	}
	c.Push(key)
	return false, time.Now()

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
