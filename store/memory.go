package store

// InMemory : in-memory data structure
type InMemory struct {
	m map[string]*Configuration
}

// InMemoryStore : implementation of in-memory storage backend
func InMemoryStore() *InMemory {
	var m = map[string]*Configuration{}
	return &InMemory{
		m: m,
	}
}

// StoreConfiguration : store a configuration
func (r *InMemory) StoreConfiguration(c *Configuration) {
	r.m[c.Key] = c
}

// RetrieveConfiguration : retrieve a configuration
func (r *InMemory) RetrieveConfiguration(key string) (*Configuration, error) {
	if r.m[key] == nil {
		return nil, ErrNotFound
	}
	return r.m[key], nil
}
