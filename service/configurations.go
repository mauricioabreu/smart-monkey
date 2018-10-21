package service

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
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

	filePath := fmt.Sprintf("/tmp/%s.conf", key)

	// Compare files before copying them
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		if compareDigest(filePath, configuration.Digest) == true {
			log.Printf("Files are equal. File %s matches the digest %s", filePath, configuration.Digest)
			return nil
		}
	}
	log.Printf("File %s is outdated. Copying...", filePath)
	// Write in on the disk
	writeConfiguration(filePath, configuration.Template)
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

func md5HashFromFile(filePath string) (string, error) {
	var md5Hash string

	file, err := os.Open(filePath)
	if err != nil {
		return md5Hash, err
	}

	defer file.Close()

	hash := md5.New()

	if _, err := io.Copy(hash, file); err != nil {
		return md5Hash, err
	}

	hashBytes := hash.Sum(nil)[:16]

	md5Hash = hex.EncodeToString(hashBytes)

	return md5Hash, nil
}

func compareDigest(filePath string, digest string) bool {
	md5Hash, err := md5HashFromFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return md5Hash == digest
}
