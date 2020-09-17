package interfaces

// Entity is...
type Entity interface {
	Id() string
}

// Cache is...
type Cache interface {
	Get(string, Entity) error
	Set(string, Entity, int) error
	Delete(string) error
	DeleteByPattern(string) (int64, error)
	Purge() (int64, error)
}
