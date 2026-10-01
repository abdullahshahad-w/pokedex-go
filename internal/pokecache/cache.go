package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	Cache map[string]cacheEntry
	mu *sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		Cache: map[string]cacheEntry{},
		mu: &sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}

	c.mu.Lock()
	c.Cache[key] = entry
	c.mu.Unlock() 

	fmt.Printf("\nAdded to cache: %s\n\n", key)
}

func (c *Cache) Get(key string) ([]byte, bool) {
	fmt.Printf("\nLooking for cache: %s\n", key)
	c.mu.Lock()
	elem, ok := c.Cache[key]
	c.mu.Unlock()

	if !ok {
		fmt.Printf("\nCache Miss: %s\n", key)
		return nil, false
	}
	fmt.Printf("\nCache Hit: %s\n\n", key)
	return elem.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.Cache {
			currTime := time.Now()
			if currTime.After(entry.createdAt.Add(interval)) {
				delete(c.Cache, key)
			}
		}
		c.mu.Unlock()
	}
}