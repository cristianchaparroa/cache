package datasources

import (
	"cache/core"
	"container/list"
	"fmt"
	"github.com/thoas/go-funk"
	"os"
	"strconv"
)

func init() {
	err := core.Injector.Provide(NewMemoryStorage)
	core.CheckInjection(err, "NewMemoryStorage")
}

const (
	defaultCapacity = 10000
)

type Record struct {
	key   interface{}
	value interface{}
}

type memory struct {
	storage map[string]*list.Element
	ll      *list.List
	slots   int
}

func NewMemoryStorage() Storage {
	ns := os.Getenv("SLOTS")

	numberSlots := defaultCapacity
	if !funk.IsEmpty(ns) {
		slots, err := strconv.Atoi(ns)
		if err != nil {
			panic(err)
		}
		numberSlots = slots
	}
	storage := make(map[string]*list.Element, numberSlots)
	return &memory{
		storage: storage,
		slots:   numberSlots,
		ll:      list.New(),
	}
}

func (m *memory) Capacity() int {
	return m.slots
}

func (m *memory) Len() int {
	return len(m.storage)
}

func (m *memory) Add(key string, object interface{}) bool {

	if m.slots == len(m.storage) {
		return false
	}

	_, exist := m.storage[key]

	if !exist {
		element := m.ll.PushBack(&Record{
			key:   key,
			value: object,
		})
		m.storage[key] = element
	} else {
		fmt.Println(m.storage[key])
		m.storage[key].Value.(*Record).value = object
	}

	return !exist
}

func (m *memory) Get(key string) (interface{}, bool) {
	v, exist := m.storage[key]
	if exist {
		return v.Value.(*Record).value, true
	}
	return nil, exist
}

func (m *memory) Delete(key string) (interface{}, bool) {
	v, exist := m.storage[key]

	if exist {
		m.ll.Remove(v)
		delete(m.storage, key)
		return v.Value.(*Record).value, true
	}

	return nil, false
}
