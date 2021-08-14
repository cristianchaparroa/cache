package managers

import (
	"cache/app/datasources"
	"cache/objects"
	"cache/objects/ports"
	"errors"
)

type LIFOCache struct {
	storage datasources.Storage
}

func NewLIFOCache(storage datasources.Storage) ports.CacheManager {
	return &LIFOCache{storage: storage}
}

func (c *LIFOCache) Add(key string, o *objects.Object) bool {

	if !c.storage.IsFull() {
		return c.storage.Add(key, o)
	}

	back := c.storage.Back()
	_, isDeleted := c.storage.Delete(back.KeyToString())

	if isDeleted {
		return false
	}

	return c.storage.Add(key, o)
}

func (c *LIFOCache) Delete(key string) (*objects.Object, error) {
	obj, bool := c.storage.Delete(key)
	if bool {
		return obj.(*objects.Object), nil
	}
	return nil, errors.New("object not found")
}
