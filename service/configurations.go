package service

import "github.com/mauricioabreu/smart-monkey/store"

// ConfigurationService : handle configuration deploy
type ConfigurationService struct {
	store store.Store
}

// InitService : service to execute actions on configurations
func InitService(s store.Store) *ConfigurationService {
	return &ConfigurationService{
		store: s,
	}
}

// RetrieveConfiguration : retrieve a configuration from the storage
func (s *ConfigurationService) RetrieveConfiguration(key string) (*store.Configuration, error) {
	return s.store.RetrieveConfiguration(key)
}
