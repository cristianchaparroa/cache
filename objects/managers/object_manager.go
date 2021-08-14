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

func (m *objectManager) Create(o *objects.Object) (*objects.Object, error) {
	return nil, nil
}
