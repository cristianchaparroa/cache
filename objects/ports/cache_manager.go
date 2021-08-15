package ports

import "cache/objects"

// CacheManager is in charge to perform operations with objects to be stored
type CacheManager interface {

	// Add adds a new object, it generate an error according
	// with eviction policies.
	Add(key string, o *objects.Object) bool

	// Get retrieves an object if exists or is not expired
	Get(key string) (*objects.Object, error)

	// It removes a object according with the key identifier.
	Delete(key string) (*objects.Object, error)
}
