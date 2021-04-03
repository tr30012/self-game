package main

const (
	configPath = "./config.json"
)

var (
	config  *Config
	api     *APIServer
	storage *StorageDB

	players   []Player
	questions []Question
	volumes   []Volume
)

func main() {
	players = make([]Player, 0)
	questions = make([]Question, 0)
	volumes = make([]Volume, 0)

	config = OpenConfig(configPath)
	storage = CreateStorageDB(config.StorageConfig)
	api = CreateAPIServer(config.APIServerConfig, storage)

	api.Start()
}
