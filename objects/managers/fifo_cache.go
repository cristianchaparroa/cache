package managers

import (
	"cache/app/datasources"
	"cache/objects"
	"cache/objects/ports"
	"errors"
)

type FIFOCache struct {
	storage datasources.Storage
}

func NewFIFOCache(storage datasources.Storage) ports.CacheManager {
	return &FIFOCache{storage: storage}
}

func (c *FIFOCache) Add(key string, o *objects.Object) bool {
	if !c.storage.IsFull() {
		return c.storage.Add(key, o)
	}

	front := c.storage.Front()
	_, isDeleted := c.storage.Delete(front.KeyToString())

	if isDeleted {
		return false
	}

	return c.storage.Add(key, o)
}

func (c *FIFOCache) Delete(key string) (*objects.Object, error) {
	obj, bool := c.storage.Delete(key)
	if bool {
		return obj.(*objects.Object), nil
	}
	return nil, errors.New("object not found")
}
