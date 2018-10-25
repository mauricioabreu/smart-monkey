package store

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"go.etcd.io/etcd/client"
)

// Etcd : etcd storage client
type Etcd struct {
	client client.Client
}

// EtcdStore : implementation of etcd store
func EtcdStore() Store {
	cfg := client.Config{
		Endpoints:               []string{"http://127.0.0.1:2379"},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &Etcd{client: c}
}

// RetrieveConfiguration : retrieve a configuration
func (etcd *Etcd) RetrieveConfiguration(key string) (*Configuration, error) {
	var configuration Configuration
	kapi := client.NewKeysAPI(etcd.client)
	log.Printf("Getting '%s' key value\n", key)
	resp, err := kapi.Get(context.Background(), key, nil)
	if err != nil {
		log.Printf("Key %s not found\n", key)
		return nil, ErrNotFound
	}
	log.Printf("Key %s found with value %s\n", resp.Node.Key, resp.Node.Value)
	err = json.Unmarshal([]byte(resp.Node.Value), &configuration)
	if err != nil {
		log.Println("Error decoding JSON into struct", err)
		return nil, nil
	}
	return &configuration, nil
}

// StoreConfiguration : store a configuration
func (etcd *Etcd) StoreConfiguration(c *Configuration) {
	kapi := client.NewKeysAPI(etcd.client)
	log.Printf("Setting '%s' key value\n", c.Key)
	value, _ := json.Marshal(c)
	resp, err := kapi.Set(context.Background(), c.Key, string(value), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Key '%s' written with data %q\n", c.Key, resp)
	}
}
