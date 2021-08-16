package managers

import (
	"cache/app/conf"
	"cache/app/datasources"
	"cache/objects"
	"cache/objects/ports"
)

type SimpleCache struct {
	*baseCache
}

func NewSimpleCache(storage datasources.Storage) ports.CacheManager {
	return &SimpleCache{baseCache: &baseCache{storage: storage}}
}

func (c *SimpleCache) GetType() string {
	return conf.RejectEvictionPolicy
}

func (c *SimpleCache) Add(key string, o *objects.Object) bool {
	return c.storage.Add(key, o)
}
