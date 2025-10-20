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
	interval  time.Duration
	cacheItem map[string]cacheEntry
	mu        *sync.Mutex
}

func NewCache(interval time.Duration) *Cache {

	item := make(map[string]cacheEntry)

	cac := &Cache{
		cacheItem: item,
		mu:        &sync.Mutex{},
		interval:  interval,
	}
	go cac.reapLoop()
	return cac
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheItem[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, exists := c.cacheItem[key]
	return item.val, exists
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for {
		<-ticker.C
		c.mu.Lock()
		for key, val := range c.cacheItem {
			durationPassed := time.Since(val.createdAt)
			if durationPassed > c.interval {
				delete(c.cacheItem, key)
			}
		}
		c.mu.Unlock()
	}
}
