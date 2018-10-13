package main

import (
	"log"

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
	configurationService.Install(key)
}
