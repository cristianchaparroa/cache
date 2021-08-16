package managers

import (
	"cache/app/conf"
	"cache/core"
	"cache/objects"
	"cache/objects/ports"
)

func init() {
	err := core.Injector.Provide(newObjectManager)
	core.CheckInjection(err, "newObjectManager")
}

type objectManager struct {
	cache ports.CacheManager
}

func newObjectManager(conf *conf.Config, builder ports.CacheBuilder) ports.ObjectManager {
	return &objectManager{cache: builder.Build(conf.Policy)}
}

func (m *objectManager) Save(key string, o *objects.Object) (*objects.Object, error) {
	_, err := m.cache.Get(key)

	if err == ports.ObjectNotFound {
		addedObject := m.cache.Add(key, o)
		if addedObject {
			return o, nil
		}
		return nil, ports.ObjectNotStored
	}

	isUpdated := m.cache.Update(key, o)

	if !isUpdated {
		return nil, ports.ObjectNotStored
	}

	return o, nil
}

func (m *objectManager) GetByKey(key string) (*objects.Object, error) {
	return m.cache.Get(key)
}

func (m *objectManager) Delete(key string) error {
	_, err := m.cache.Delete(key)
	return err
}
