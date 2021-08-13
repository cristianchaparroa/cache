package handlers

import (
	"cache/core"
	"github.com/gin-gonic/gin"
)

func init() {
	err := core.Injector.Provide(newObjectWriter)
	core.CheckInjection(err, "newObjectWriter")
}

// ObjectsHandler is in charge to handle the HTTP request for objects
type ObjectsHandler struct {
}

// NewObjectsHandler it creates a pointer to ObjectsHandler
func newObjectWriter() *ObjectsHandler {
	return &ObjectsHandler{}
}

// Save storages an object in cache
func (h *ObjectsHandler) Save(c *gin.Context) {

}

// GetByKey retrieves a object stored in cache by key
func (h *ObjectsHandler) GetByKey(c *gin.Context) {

}

// DeleteByKey is in charge to delete an existent object in cache
func (h *ObjectsHandler) DeleteByKey(c *gin.Context) {

}
