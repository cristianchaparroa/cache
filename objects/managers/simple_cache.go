package managers

import (
	"cache/app/datasources"
	"cache/objects"
	"cache/objects/ports"
)

type SimpleCache struct {
	storage datasources.Storage
}

func NewSimpleCache(storage datasources.Storage) ports.CacheManager {
	return &SimpleCache{storage: storage}
}

func (c *SimpleCache) Add(key string, o *objects.Object) error {
	panic("implement me")
}

func (c *SimpleCache) Delete(key string) error {
	panic("implement me")
}
