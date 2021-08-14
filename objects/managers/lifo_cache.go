package managers

import (
	"cache/app/datasources"
	"cache/objects"
	"cache/objects/ports"
)

type LIFOCache struct {
	storage datasources.Storage
}

func NewLIFOCache(storage datasources.Storage) ports.CacheManager {
	return &LIFOCache{storage: storage}
}

func (c *LIFOCache) Add(key string, o *objects.Object) error {
	panic("implement me")
}

func (c *LIFOCache) Delete(key string) error {
	panic("implement me")
}
