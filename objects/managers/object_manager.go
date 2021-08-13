package managers

import (
	"cache/core"
	"cache/objects"
	"cache/objects/ports"
)

func init() {
	err := core.Injector.Provide(newObjectManager)
	core.CheckInjection(err, "newObjectManager")
}

type objectManager struct{}

func newObjectManager() ports.ObjectManager {
	return &objectManager{}
}

func (m *objectManager) Create(o *objects.Object) (*objects.Object, error) {
	panic("implement me")
}
