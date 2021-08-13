package ports

import (
	"cache/objects"
)

// ObjectWriter is in charge to perform operations to storage objects
type ObjectWriter interface {

	// It saves an object
	Save(key string, o *objects.Object) error

	// It deletes a specific object
	Delete(key string) error
}
