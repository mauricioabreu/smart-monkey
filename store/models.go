package store

// Configuration : representation of a configuration
type Configuration struct {
	Key      string `json:"key"`
	Template string `json:"template"`
	Digest   string `json:"digest"`
}
