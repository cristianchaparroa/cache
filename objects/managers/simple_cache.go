package managers

import (
	"cache/app/datasources"
	"cache/objects"
	"cache/objects/ports"
	"errors"
)

type SimpleCache struct {
	storage datasources.Storage
}

func NewSimpleCache(storage datasources.Storage) ports.CacheManager {
	return &SimpleCache{storage: storage}
}

func (c *SimpleCache) Add(key string, o *objects.Object) bool {
	return c.storage.Add(key, o)
}

func (c *SimpleCache) Delete(key string) (*objects.Object, error) {
	obj, bool := c.storage.Delete(key)
	if bool {
		return obj.(*objects.Object), nil
	}
	return nil, errors.New("object not found")
}
