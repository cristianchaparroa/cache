package managers

import (
	"cache/app/datasources"
	"cache/objects"
	"errors"
)

type baseCache struct {
	storage datasources.Storage
}

func (c *baseCache) Delete(key string) (*objects.Object, error) {
	obj, bool := c.storage.Delete(key)
	if bool {
		return obj.(*objects.Object), nil
	}
	return nil, errors.New("object not found")
}

func (c *baseCache) Get(key string) (*objects.Object, error) {

	obj, exist := c.storage.Get(key)

	if !exist {
		return nil, objectNotFound
	}

	o := obj.(*objects.Object)
	if o.IsExpired() {
		return nil, objectNotFound
	}

	return o, nil
}
