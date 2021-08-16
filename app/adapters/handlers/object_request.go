package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
)

type objectRequest struct {
	Key string
}

func NewObjectRequest(c *gin.Context) (*objectRequest, error) {
	keyParam := c.Param("key")

	if funk.IsEmpty(keyParam) {
		return nil, badRequest
	}

	return &objectRequest{Key: keyParam}, nil
}
