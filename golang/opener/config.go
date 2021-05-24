package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/pelletier/go-toml"
)

const (
	TYPE_FILE string = "file"
	TYPE_DIR  string = "dir"
	TYPE_WWW  string = "www"
	TYPE_BEAR string = "bear"
)

type Item struct {
	Key         string
	Type        string
	Description string
	Uri         string
}

type Config struct {
	Item []Item
}

func LoadConfig(configPath string) Config {
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	config := Config{}
	toml.Unmarshal(data, &config)
	return config
}

func (config *Config) Find(key string) *Item {
	for _, item := range config.Item {
		if item.Key == key {
			return &item
		}
	}

	return nil
}
