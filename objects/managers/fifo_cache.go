package managers

import (
	"cache/app/datasources"
	"cache/objects"
	"cache/objects/ports"
)

type FIFOCache struct {
	storage datasources.Storage
}

func NewFIFOCache(storage datasources.Storage) ports.CacheManager {
	return &FIFOCache{storage: storage}
}

func (c *FIFOCache) Add(key string, o *objects.Object) error {
	panic("implement me")
}

func (c *FIFOCache) Delete(key string) error {
	panic("implement me")
}
