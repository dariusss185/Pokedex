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
	sessions map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

func NewCache(timeout time.Duration) *Cache {
	cache := &Cache{
		sessions: make(map[string]cacheEntry),
		interval: timeout,
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(id string, token []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.sessions[id] = cacheEntry{
		createdAt: time.Now(),
		val:       token,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.sessions[key]
	return entry.val, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.reap()
	}
}

func (c *Cache) reap() {
	// 1. lock the mutex, defer the unlock
	c.mu.Lock()
	defer c.mu.Unlock()

	// 2.
	cutoff := time.Now().Add(-c.interval)

	for key, value := range c.sessions {
		if value.createdAt.Before(cutoff) {
			delete(c.sessions, key)
		}
	}
}
