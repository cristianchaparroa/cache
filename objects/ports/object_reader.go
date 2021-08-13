package ports

import (
	"cache/objects"
)

// ObjectReader is in charge to retrieves object information
type ObjectReader interface {

	// It retrieves an specific object
	Get(key string) (*objects.Object, error)
}
