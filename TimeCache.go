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
	count   int // number of entires in the cache
}

// DictCache is an implementation of the time cache which uses a dict
type DictCache struct {
	entries map[string]time.Time // A map from a unique key to a time
	timeout time.Duration
	count   int // number of entires in the cache
}

// Contains returns boolean true if the cache contains key, as well as time key
// was added. If key is not in cache, returns false and adds entry to cache
func (c *SliceCache) Contains(key string) (bool, time.Time) {
	c.Pop()
	if c.count == 0 {
		c.Push(key)
		return false, time.Now()
	}
	for i, k := range c.keys {
		if k == key {
			return true, c.times[i]
		}
	}
	c.Push(key)
	return false, time.Now()
}

// Push adds an entry to the cache
func (c *SliceCache) Push(key string) {
	if key == "" {
		return
	}
	c.keys = append(c.keys, key)
	c.times = append(c.times, time.Now())
	c.count++
}

// Pop removes all outdated entries from the cache. For the slice
// implmentation, all outdated entries are removed at the end, unlike the dict
// implementation
func (c *SliceCache) Pop() {
	f := 0 // front of the stack for trimming
	for _, t := range c.times {
		if time.Since(t) > c.timeout {
			f++
			continue
		}
		break
	}
	c.keys = c.keys[f:]
	c.times = c.times[f:]
	c.count = c.count - f
}

func NewSliceCache(t int) *SliceCache {
	var s SliceCache
	s.timeout = time.Duration(t) * time.Second
	return &s
}

// Contains returns boolean of if the cache contains key, as well as time key
// was added. If key is not in cache, returns false and adds entry to cache
func (c *DictCache) Contains(key string) (bool, time.Time) {
	c.Pop()
	if c.count == 0 {
		c.Push(key)
		return false, time.Now()
	}
	for k := range c.entries {
		if k == key {
			return true, c.entries[k]
		}
	}
	c.Push(key)
	return false, time.Now()

}

// Push adds an entry to the cache
func (c *DictCache) Push(key string) {
	if key == "" {
		return
	}
	c.entries[key] = time.Now()
	c.count++
}

// Pop removes all outdated entries from the cache. In the dict inpmentation,
// the entire cache is read and any keys which have expired are removed
func (c *DictCache) Pop() {
	for k, t := range c.entries {
		if time.Since(t) > c.timeout {
			delete(c.entries, k)
			c.count--
		}
	}
}

// NewDictCache is a contructor to intialize a new cache - the dictionary must
// be declared with make
func NewDictCache(t int) *DictCache {
	var d DictCache
	d.entries = make(map[string]time.Time)
	d.timeout = time.Duration(t) * time.Second
	return &d
}
