package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type APIServerConfig struct {
	Address    string `json:"address"`
	Uixpath    string `json:"uixpath"`
	Staticpath string `json:"staticpath"`
}

type StorageConfig struct {
	DBPath string `json:"dbpath"`
}

type Config struct {
	APIServerConfig *APIServerConfig `json:"apiserver"`
	StorageConfig   *StorageConfig   `json:"storage"`
}

func OpenConfig(configPath string) *Config {
	f, err := os.Open(configPath)

	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(f)

	if err != nil {
		panic(err)
	}

	config := Config{}

	err = json.Unmarshal(content, &config)

	if err != nil {
		panic(err)
	}

	return &config
}
