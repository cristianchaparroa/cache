package managers

import (
	"cache/app/conf"
	"cache/app/datasources"
	"cache/objects"
	"cache/objects/ports"
)

type LIFOCache struct {
	*baseCache
}

func NewLIFOCache(storage datasources.Storage) ports.CacheManager {
	return &LIFOCache{
		baseCache: &baseCache{storage: storage},
	}
}

func (c *LIFOCache) GetType() string {
	return conf.OlderFistEvictionPolicy
}

func (c *LIFOCache) Add(key string, o *objects.Object) bool {

	if !c.storage.IsFull() {
		return c.storage.Add(key, o)
	}

	oldest := c.storage.Front()
	oldestKey := oldest.Key.(string)
	return c.storage.Set(oldestKey, key, o)
}
