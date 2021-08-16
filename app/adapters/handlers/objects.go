package handlers

import (
	"cache/core"
	"cache/objects/ports"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	err := core.Injector.Provide(newObjectWriter)
	core.CheckInjection(err, "newObjectWriter")
}

// ObjectsHandler is in charge to handle the HTTP request for objects
type ObjectsHandler struct {
	manager ports.ObjectManager
}

// NewObjectsHandler it creates a pointer to ObjectsHandler
func newObjectWriter(manager ports.ObjectManager) *ObjectsHandler {
	return &ObjectsHandler{manager: manager}
}

// Save storages an object in cache
func (h *ObjectsHandler) Save(c *gin.Context) {
	req, err := NewObjectSaveRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	o, err := h.manager.Save(req.Key, req.ToObject())

	if err != nil {
		// TODO: validate which kind of error should be here
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	c.JSON(http.StatusOK, o)
}

// GetByKey retrieves a object stored in cache by key
func (h *ObjectsHandler) GetByKey(c *gin.Context) {
	req, err := NewObjectRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	o, err := h.manager.GetByKey(req.Key)
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, o)
}

// DeleteByKey is in charge to delete an existent object in cache
func (h *ObjectsHandler) DeleteByKey(c *gin.Context) {
	req, err := NewObjectRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = h.manager.Delete(req.Key)
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}
