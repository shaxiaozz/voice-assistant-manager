package cache

import (
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	data map[string]cacheItem
}

type cacheItem struct {
	Value      interface{}
	ExpireTime time.Time
}

var defaultCache *Cache

func init() {
	defaultCache = NewCache()
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]cacheItem),
	}
}

func Set(key string, value interface{}, expiration time.Duration) {
	defaultCache.Set(key, value, expiration)
}

func Get(key string) (interface{}, bool) {
	return defaultCache.Get(key)
}

func (c *Cache) Set(key string, value interface{}, expiration time.Duration) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = cacheItem{
		Value:      value,
		ExpireTime: time.Now().Add(expiration),
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()

	item, exists := c.data[key]
	if !exists {
		return nil, false
	}

	if time.Now().After(item.ExpireTime) {
		delete(c.data, key)
		return nil, false
	}

	return item.Value, true
}
