package store

// Store : interface for the storage backend
type Store interface {
	RetrieveConfiguration(key string) (*Configuration, error)
	StoreConfiguration(c *Configuration)
}
