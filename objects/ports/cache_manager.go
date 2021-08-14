package ports

import "cache/objects"

// CacheManager is in charge to perform operations with objects to be stored
type CacheManager interface {

	// Add adds a new object, it generate an error accoriding
	// with eviction policies.
	Add(key string, o *objects.Object) bool

	// It removes a object according with the key identifier.
	Delete(key string) (*objects.Object, error)
}
