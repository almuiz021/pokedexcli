package pokecache

import "time"

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheItem[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.cacheItem[key]
	return item.val, exists
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
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
