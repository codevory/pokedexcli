package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.RWMutex
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
	}

	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) bool {
	if key == " " || key == "" || val == nil {
		return false
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}

	return true
}

func (c *Cache) Get(key string) ([]byte, bool) {
	if key == "" || key == " " {
		return nil, false
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.entries {
		if time.Since(v.createdAt) > interval {
			delete(c.entries, k)
		}
	}
}
