package ports

import (
	"cache/objects"
)

// ObjectManager is in charge to perform all operations related to objects
// in the domain.
type ObjectManager interface {

	// Create store a new object in the system
	Save(key string, o *objects.Object) (*objects.Object, error)

	// GetByKey returns an object if exist
	GetByKey(key string) (*objects.Object, error)

	// Delete remove an object
	Delete(key string) error
}
