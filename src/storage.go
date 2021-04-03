package main

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

type StorageDB struct {
	config *StorageConfig
	db     *sql.DB
	items  map[int]([]*Question)
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

func (s *StorageDB) LoadIntoMemory() {
	s.items = make(map[int]([]*Question))

	dbVolumes, err := s.db.Query("SELECT id, text FROM volumes")

	if err != nil {
		panic(err)
	}

	v := Volume{}

	for dbVolumes.Next() {
		dbVolumes.Scan(&v.Id, &v.Text)

		volumes = append(volumes, Volume{
			Id:   v.Id,
			Text: v.Text,
		})
	}

	dbQuestions, err := s.db.Query("SELECT id, text, answer, event, cost, volume_id FROM questions")

	if err != nil {
		panic(err)
	}

	q := Question{}

	for dbQuestions.Next() {
		dbQuestions.Scan(&q.Id, &q.Text, &q.Answer, &q.Event, &q.Cost, &q.VolumeId)

		questions = append(questions, Question{
			Id:       q.Id,
			Text:     q.Text,
			Answer:   q.Answer,
			Event:    q.Event,
			Cost:     q.Cost,
			VolumeId: q.VolumeId,
			Answered: false,
		})
	}

	if s.config.QIterStart < 0 {
		panic(errors.New("volumes index must be greater 0"))
	}
	if s.config.QIterStart > s.config.QIterEnd {
		panic(errors.New("volumes start index must be less then end index"))
	}
	if s.config.QIterEnd > len(volumes) {
		panic(errors.New("volumes end index must be less then it's size"))
	}

}
