package managers

import (
	"cache/app/conf"
	"cache/app/datasources"
	"cache/core"
	"cache/objects/ports"
)

func init() {
	err := core.Injector.Provide(NewCacheBuilder)
	core.CheckInjection(err, "NewCacheBuilder")
}

type cacheBuilder struct {
	storage datasources.Storage
}

func NewCacheBuilder(storage datasources.Storage) ports.CacheBuilder {
	return &cacheBuilder{storage: storage}
}

func (b *cacheBuilder) Build(evictionPolicy string) ports.CacheManager {

	if conf.OlderFistEvictionPolicy == evictionPolicy {
		return NewLIFOCache(b.storage)
	}

	if conf.NewestFirstEvictionPolicy == evictionPolicy {
		return NewFIFOCache(b.storage)
	}

	return NewSimpleCache(b.storage)
}
