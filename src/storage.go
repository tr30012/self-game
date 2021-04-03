package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type StorageDB struct {
	config *StorageConfig
	db     *sql.DB
}

func CreateStorageDB(cfg *StorageConfig) *StorageDB {
	connection, err := sql.Open("sqlite3", cfg.DBPath)

	if err != nil {
		panic(err)
	}

	if err = connection.Ping(); err != nil {
		panic(err)
	}

	return &StorageDB{
		config: cfg,
		db:     connection,
	}
}
