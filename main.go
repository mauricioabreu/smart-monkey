package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mauricioabreu/smart-monkey/service"
	"github.com/mauricioabreu/smart-monkey/store"
)

func main() {
	log.Println("smart monkeys is starting...")

	key := "1"
	template := "foo"
	repository := store.InMemoryStore()
	// Insert a new configuration in the storage
	repository.StoreConfiguration(&store.Configuration{Key: key, Template: template})
	configurationService := service.InitService(repository)
	// Retrieve it from the storage
	configuration, err := configurationService.RetrieveConfiguration(key)
	if err != nil {
		panic(err)
	}
	log.Printf("Configuration %s found: %v\n", key, configuration)
	// Write in on the disk
	writeConfiguration(fmt.Sprintf("/tmp/%s.conf", key), "foo")
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

	log.Printf("wrote %d bytes on %s", byteSize, destination)
}
