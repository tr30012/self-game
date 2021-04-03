package main

const (
	configPath = "./config.json"
)

var (
	config  *Config
	api     *APIServer
	storage *StorageDB
)

func main() {
	config = OpenConfig(configPath)
	storage = CreateStorageDB(config.StorageConfig)
	api = CreateAPIServer(config.APIServerConfig, storage)

	api.Start()
}
