package managers

import (
	"cache/app/datasources"
	"cache/objects"
	"cache/objects/ports"
)

type baseCache struct {
	storage datasources.Storage
}

func (c *baseCache) Delete(key string) (*objects.Object, error) {
	obj, bool := c.storage.Delete(key)
	if bool {
		return obj.(*objects.Object), nil
	}
	return nil, ports.ObjectNotFound
}

func (c *baseCache) Get(key string) (*objects.Object, error) {

	obj, exist := c.storage.Get(key)

	if !exist {
		return nil, ports.ObjectNotFound
	}

	o := obj.(*objects.Object)

	if o.TTL == objects.DefaultTTL {
		return o, nil
	}

	if o.IsExpired() {
		c.storage.Delete(key)
		return nil, ports.ObjectNotFound
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
