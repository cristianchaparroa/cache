package repositories

import (
	"cache/core"
	"cache/objects"
	"cache/objects/ports"
)

func init() {
	err := core.Injector.Provide(newObjectWriter)
	core.CheckInjection(err, "newObjectWriter")
}

type objectWriter struct {
}

func newObjectWriter() ports.ObjectWriter {
	return &objectWriter{}
}

func (w *objectWriter) Save(key string, o *objects.Object) error {
	panic("implement me")
}

func (w *objectWriter) Delete(key string) error {
	panic("implement me")
}
