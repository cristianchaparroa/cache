package datasources

type Storage interface {
	Capacity() int

	Len() int

	Add(key string, object interface{}) error

	Get(key string) (interface{}, bool)

	Delete(key string) (interface{}, bool)
}
