package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
		mu:      &sync.Mutex{},
	}
	go reapLoop(c, interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func reapLoop(c *Cache, interval time.Duration) {
	// log.Printf("Starting reap loop with interval: %v", interval)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > interval {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
