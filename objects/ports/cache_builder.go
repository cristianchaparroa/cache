package ports

type CacheBuilder interface {
	Build(evictionPolicy string) CacheManager
}
