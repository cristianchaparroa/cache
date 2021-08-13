package datasources

import (
	"cache/core"
	"errors"
	"os"
	"github.com/thoas/go-funk"
	"strconv"
)

func init() {
	err := core.Injector.Provide(NewMemoryStorage())
	core.CheckInjection(err, "NewMemoryStorage")
}

type memory struct {
	storage map[string]interface{}
	slots int
}

func NewMemoryStorage() Storage {
	ns := os.Getenv("SLOTS")

	numberSlots := 10000
	if !funk.IsEmpty(ns) {
		slots, err := strconv.Atoi(ns)
		if err != nil {
			panic(err)
		}
		numberSlots = slots
	}
	storage := make(map[string]interface{}, numberSlots)
	return &memory{storage: storage, slots: numberSlots}
}

func (m *memory) Add(key string, object interface{}) error {

	if m.slots == len(m.storage) {
		return errors.New("storage is full")
	}

	m.storage[key] = object
	return nil
}

func (m *memory) Delete(key string) (interface{}, bool) {
	v, exist := m.storage[key]

	if exist {
		delete(m.storage, key)
		return v, true
	}

	return nil, false
}
