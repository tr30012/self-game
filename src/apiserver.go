package main

type APIServer struct {
	config  *APIServerConfig
	storage *StorageDB
}

func CreateAPIServer(cfg *APIServerConfig, db *StorageDB) *APIServer {
	return &APIServer{
		config:  cfg,
		storage: db,
	}
}

func (s *APIServer) Start() {

}
