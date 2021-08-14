package managers

import (
	"cache/app/datasources"
	"cache/core"
	"cache/objects/ports"
)

func init() {
	err := core.Injector.Provide(NewCacheBuilder)
	core.CheckInjection(err, "NewCacheBuilder")
}

const (
	olderFistEvictionPolicy   = "OLDEST_FIRST"
	newestFirstEvictionPolicy = "NEWEST_FIRST"
	rejectEvictionPolicy      = "REJECT"
)

type cacheBuilder struct {
	storage datasources.Storage
}

func NewCacheBuilder(storage datasources.Storage) ports.CacheBuilder {
	return &cacheBuilder{storage: storage}
}

func (b *cacheBuilder) Build(evictionPolicy string) ports.CacheManager {

	if olderFistEvictionPolicy == evictionPolicy {
		return NewLIFOCache(b.storage)
	}

	if newestFirstEvictionPolicy == evictionPolicy {
		return NewLIFOCache(b.storage)
	}

	return NewSimpleCache(b.storage)
}
