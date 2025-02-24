// cache is an in-memory map to speed up shortURL fetching.
package cache

import (
	"container/list"
	"sync"
)

// Cache represents a cached item with its short URL and long URL.
type Cache struct {
	shortURL string
	longURL  string
}

// CacheMap is an in-memory cache for short URL to long URL mappings.
//
// It uses a hash map with a doubly linked list for LRU eviction.
type CacheMap struct {
	cache     map[string]*list.Element
	cacheLock sync.RWMutex
	lruList   *list.List
	capacity  int
}

// NewCacheMap creates a new CacheMap with the specified capacity.
func NewCacheMap(capacity int) *CacheMap {
	return &CacheMap{
		cache:    make(map[string]*list.Element),
		capacity: capacity,
		lruList:  list.New(),
	}
}

// Get returns the corresponding long URL for the given short URL.
//
// If the short URL exists, it moves the item to the front of the LRU list.
func (cm *CacheMap) Get(shortURL string) (string, bool) {
	cm.cacheLock.RLock()
	defer cm.cacheLock.RUnlock()

	if elem, exists := cm.cache[shortURL]; exists {
		cm.lruList.MoveToFront(elem)
		cacheItem := elem.Value.(*Cache)
		return cacheItem.longURL, true
	}

	return "", false
}

// Set adds or updates the shortURL:longURL mapping in the cache.
//
// If the cache is full, it evicts the least recently used item.
func (cm *CacheMap) Set(shortURL, longURL string) {
	cm.cacheLock.Lock()
	defer cm.cacheLock.Unlock()

	// if short URL already exists, move it to front and return the long URL
	if elem, exists := cm.cache[shortURL]; exists {
		cm.lruList.MoveToFront(elem)
		cacheItem := elem.Value.(*Cache)
		cacheItem.longURL = longURL
		return
	}

	// If lru list is full, remove the last item (least recently used) from the list
	if cm.lruList.Len() >= cm.capacity {
		lru := cm.lruList.Back()
		if lru != nil {
			item := lru.Value.(*Cache)
			delete(cm.cache, item.shortURL)
			cm.lruList.Remove(lru)
		}
	}

	// cache the shortURL and update the lru list
	cacheItem := &Cache{
		shortURL: shortURL,
		longURL:  longURL,
	}
	elem := cm.lruList.PushFront(cacheItem)
	cm.cache[shortURL] = elem
}
