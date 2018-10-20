package service

import (
	"fmt"
	"log"
	"os"

	"github.com/mauricioabreu/smart-monkey/store"
)

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

// Install : Deploy a configuration for the given key
func (s *ConfigurationService) Install(key string) error {
	// Retrieve it from the storage
	configuration, err := s.RetrieveConfiguration(key)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("Configuration %s found: %v\n", key, configuration)
	// Write in on the disk
	writeConfiguration(fmt.Sprintf("/tmp/%s.conf", key), configuration.Template)
	return nil
}

func writeConfiguration(destination string, content string) {
	file, err := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	byteSize, err := file.WriteString(content)
	if err != nil {
		panic(err)
	}

	log.Printf("Wrote %d bytes on %s", byteSize, destination)
}
