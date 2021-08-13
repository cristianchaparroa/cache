package datasources

type Storage interface {
	Capacity() int

	Len() int

	Add(key string, object interface{}) error

	Delete(key string) (interface{}, bool)
}
