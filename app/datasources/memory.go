package datasources

import (
	"cache/core"
	"os"
	"github.com/thoas/go-funk"
	"strconv"
)

func init() {
	err := core.Injector.Provide(newMemoryStorage())
	core.CheckInjection(err, "newMemoryStorage")
}

type memory struct {
	storage map[string]interface{}
}

func newMemoryStorage() Storage {
	ns := os.Getenv("SLOTS")

	numberSlots := 10000
	if funk.IsEmpty(ns) {
		slots, err := strconv.Atoi(ns)
		if err != nil {
			panic(err)
		}
		numberSlots = slots
	}
	storage := make(map[string]interface{}, numberSlots)
	return &memory{storage: storage}
}

func (m *memory) Add(key string, object interface{}) error {
	panic("implement me")
}

func (m *memory) Delete(key string) (interface{}, error) {
	panic("implement me")
}
