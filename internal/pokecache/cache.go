package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	interval  time.Duration
	cacheItem map[string]cacheEntry
	mu        *sync.RWMutex
}

func NewCache(interval time.Duration) *Cache {

	item := make(map[string]cacheEntry)

	cac := &Cache{
		cacheItem: item,
		mu:        &sync.RWMutex{},
		interval:  interval,
	}
	go cac.reapLoop()
	return cac
}
