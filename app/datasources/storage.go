package datasources

type Storage interface {
	Capacity() int

	Len() int

	IsFull() bool

	Add(key string, object interface{}) bool

	Set(oldKey, newKey string, object interface{}) bool

	Get(key string) (interface{}, bool)

	Delete(key string) (interface{}, bool)

	Front() *Element

	Back() *Element
}
