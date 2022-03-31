package cache

import (
	"time"

	"github.com/ReneKroon/ttlcache/v2"
)

type Cache struct {
	ttlCache ttlcache.SimpleCache
}

type ICache interface {
	Set(key string, value interface{}) error
	SetTTL(key string, value interface{}, t time.Duration) error
	Get(key string) (interface{}, error)
	Close()
}

func NewCache() ICache {
	cache := ttlcache.NewCache()
	cache.SkipTTLExtensionOnHit(true)
	return &Cache{
		ttlCache: cache,
	}
}

var MemCache ICache

func (c *Cache) Set(key string, value interface{}) error {
	err := c.ttlCache.SetWithTTL(key, "value", 10*time.Second)
	return err
}

func (c *Cache) SetTTL(key string, value interface{}, ttl time.Duration) error {
	err := c.ttlCache.SetWithTTL(key, value, ttl)
	return err
}

func (c *Cache) Get(key string) (interface{}, error) {
	value, err := c.ttlCache.Get(key)
	if err == ttlcache.ErrNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

func (c *Cache) Close() {
	c.ttlCache.Close()
}
