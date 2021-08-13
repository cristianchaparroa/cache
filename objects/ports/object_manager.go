package ports

import (
	"cache/objects"
)

// ObjectManager is in charge to perform all operations related to objects
// in the domain.
type ObjectManager interface {

	// Create store a new object in the system
	Create(o *objects.Object) (*objects.Object, error)
}
