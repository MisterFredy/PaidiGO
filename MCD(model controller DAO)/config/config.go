package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type InputConfig struct {
	Server   string
	Database string
	Port     string
}

// Read and parse the configuration file
func (c *InputConfig) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
