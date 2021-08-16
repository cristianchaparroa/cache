package handlers

import (
	"cache/app/conf"
	"cache/objects"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"io/ioutil"
	"strconv"
)

type objectSaveRequest struct {
	Key  string
	Data string
	TTL  int64
}

func NewObjectSaveRequest(c *gin.Context, config *conf.Config) (*objectSaveRequest, error) {
	keyParam := c.Param("key")
	if funk.IsEmpty(keyParam) {
		return nil, badRequest
	}

	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, badRequest
	}

	ttlParam := c.Query("ttl")
	ttl := int64(0)

	if !funk.IsEmpty(ttlParam) {
		ttlParsed, err := strconv.ParseInt(ttlParam, 10, 64)
		if err != nil {
			ttl = ttlParsed
		}
	} else {
		ttl = config.TTL
	}

	return &objectSaveRequest{
		Key:  keyParam,
		Data: string(b),
		TTL:  ttl,
	}, nil
}

func (r *objectSaveRequest) ToObject() *objects.Object {
	if r.TTL > 0 {
		return objects.NewObjectWithTTL(r.Data, r.TTL)
	}
	return objects.NewObject(r.Data)
}
