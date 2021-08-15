package datasources

import (
	"cache/app/conf"
	"cache/core"
	"container/list"
	"fmt"
)

func init() {
	err := core.Injector.Provide(NewMemoryStorage)
	core.CheckInjection(err, "NewMemoryStorage")
}

const (
	defaultCapacity = 10000
	defaultSlotsEnv = "DEFAULT_SLOTS"
)

type Element struct {
	Key     interface{}
	Value   interface{}
	element *list.Element
}

func newElement(e *list.Element) *Element {
	if e == nil {
		return nil
	}

	element := e.Value.(*Record)

	return &Element{
		element: e,
		Key:     element.key,
		Value:   element.value,
	}
}
func (e *Element) KeyToString() string {
	return e.Key.(string)
}

type Record struct {
	key   interface{}
	value interface{}
}

type memory struct {
	storage map[string]*list.Element
	ll      *list.List
	slots   int
}

func NewMemoryStorage(c *conf.Config) Storage {
	storage := make(map[string]*list.Element, c.Slots)
	return &memory{
		storage: storage,
		slots:   c.Slots,
		ll:      list.New(),
	}
}

func (m *memory) Capacity() int {
	return m.slots
}

func (m *memory) Len() int {
	return len(m.storage)
}

func (m *memory) IsFull() bool {
	return m.Capacity() == m.Len()
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

func (m *memory) Front() *Element {
	front := m.ll.Front()
	return newElement(front)
}

func (m *memory) Back() *Element {
	back := m.ll.Back()
	return newElement(back)
}
