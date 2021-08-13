package datasources

type Storage interface {

	Add(key string, object interface{}) error

	Delete(key string) (interface{}, bool)
}