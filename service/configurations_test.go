package service

import (
	"io/ioutil"
	"testing"

	"github.com/mauricioabreu/smart-monkey/store"
)

func TestInstall(t *testing.T) {
	repository := store.InMemoryStore()
	// Insert a new configuration in the storage
	repository.StoreConfiguration(&store.Configuration{Key: "1", Template: "foo"})
	configurationService := InitService(repository)
	configurationService.Install("1")
	// Ensure the file has the content we are looking for
	data, _ := ioutil.ReadFile("/tmp/1.conf")
	if string(data) != "foo" {
		t.Errorf("File content should be 'foo'. Got %s", data)
	}
}
