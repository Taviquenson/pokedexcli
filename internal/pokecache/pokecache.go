package pokecache

import (
	// "fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	entry     []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := new(Cache)
	cache.entries = make(map[string]cacheEntry)
	go cache.reapLoop(interval)
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.entries[key] = cacheEntry{createdAt: time.Now(), entry: val}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cacheEntry, exists := cache.entries[key]
	if exists {
		return cacheEntry.entry, true
	} else {
		return []byte(""), false
	}
}

// remove entries older than `interval“
func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// infinite loop checking cache entries
	for t := range ticker.C { // t saves the time of the tick
		cache.mu.Lock()
		for key, val := range cache.entries {
			if t.Sub(val.createdAt).Seconds() > 5 { // if cache entry is older than 5 secs
				delete(cache.entries, key)
			}
		}
		cache.mu.Unlock()
	}
}
