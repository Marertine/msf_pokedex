package pokecache

import (
	//"net/http"
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		now := time.Now().UTC()
		c.mux.Lock()

		for k, v := range c.cache {
			if v.createdAt.Before(now.Add(-interval)) {
				// delete cached information that is older than interval
				delete(c.cache, k)
			}
		}

		c.mux.Unlock()
	}
}

func (c *Cache) Add(key string, value []byte) {
	// 1. lock the mutex
	c.mux.Lock()
	defer c.mux.Unlock()

	// 2. store a cacheEntry in the map under this key
	cachedData := cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
	c.cache[key] = cachedData
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	// look up in c.cache
	cachedData, found := c.cache[key]

	if !found {
		// if not found, return nil and false
		return nil, false
	}

	// if found, return the []byte and true
	return cachedData.val, true
}
