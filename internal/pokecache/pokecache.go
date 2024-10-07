package pokecache

import (
	"log"
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	log.Printf("Creating new cache with interval: %v", interval)
	c := &Cache{
		entries: make(map[string]cacheEntry),
		mu:      &sync.Mutex{},
	}
	go reapLoop(c, interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	log.Printf("Adding key: %s", key)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	log.Printf("Added key: %s", key)
}

func (c *Cache) Get(key string) ([]byte, bool) {
	log.Printf("Getting key: %s", key)
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[key]
	if !ok {
		log.Printf("Key not found: %s", key)
		return nil, false
	}
	log.Printf("Found key: %s", key)
	return entry.val, true
}

func reapLoop(c *Cache, interval time.Duration) {
	// log.Printf("Starting reap loop with interval: %v", interval)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		log.Printf("Reaping cache")
		c.mu.Lock()
		beforeCount := len(c.entries)
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > interval {
				delete(c.entries, key)
				log.Printf("Deleted expired key: %s", key)
			}
		}
		afterCount := len(c.entries)
		c.mu.Unlock()
		log.Printf("Reaped cache. Before: %d, After: %d", beforeCount, afterCount)
	}
}
