package managers

import (
	"cache/app/datasources"
	"cache/objects"
	"cache/objects/ports"
)

type FIFOCache struct {
	*baseCache
}

func NewFIFOCache(storage datasources.Storage) ports.CacheManager {
	return &FIFOCache{
		baseCache: &baseCache{storage: storage},
	}
}

func (c *FIFOCache) Add(key string, o *objects.Object) bool {
	if !c.storage.IsFull() {
		return c.storage.Add(key, o)
	}

	newest := c.storage.Back()
	_, isDeleted := c.storage.Delete(newest.KeyToString())

	if !isDeleted {
		return false
	}

	return c.storage.Add(key, o)
}
