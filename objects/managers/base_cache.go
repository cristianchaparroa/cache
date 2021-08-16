package managers

import (
	"cache/app/datasources"
	"cache/objects"
)

type baseCache struct {
	storage datasources.Storage
}

func (c *baseCache) Delete(key string) (*objects.Object, error) {
	obj, bool := c.storage.Delete(key)
	if bool {
		return obj.(*objects.Object), nil
	}
	return nil, objectNotFound
}

func (c *baseCache) Get(key string) (*objects.Object, error) {

	obj, exist := c.storage.Get(key)

	if !exist {
		return nil, objectNotFound
	}

	o := obj.(*objects.Object)

	if o.TTL == objects.DefaultTTL {
		return o, nil
	}

	if o.IsExpired() {
		c.storage.Delete(key)
		return nil, objectNotFound
	}

	return o, nil
}

func (c *baseCache) Update(key string, o *objects.Object) bool {

	_, exist := c.storage.Get(key)

	if !exist {
		return false
	}

	return c.storage.Set(key, key, o)
}
