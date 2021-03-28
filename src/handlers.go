package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

const (
	CODE_PLAYER = 0
	CODE_VOLUME = 1
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"templates\\game_page.html",
	)

	if err != nil {
		panic(err)
	}

	t.Execute(w, mainPageContent)
}

func getPageContent(w http.ResponseWriter, r *http.Request) {
	logger.Info("Получен запрос на данные от: ", r.RemoteAddr)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mainPageContent)

	logger.Info("Успешно.")
}

func setPageContent(w http.ResponseWriter, r *http.Request) {
	logger.Info("Получен запрос на обновление данных от: ", r.RemoteAddr)

	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&mainPageContent)

	logger.Info("Успешно.")
}

func restartPageContent(w http.ResponseWriter, r *http.Request) {
	logger.Info("Получен запрос перезагрузку: ", r.RemoteAddr)

	w.Header().Set("Content-Type", "application/json")
	mainPageContent = LoadMainPageContent()
	json.NewEncoder(w).Encode(mainPageContent)

	logger.Info("Успешно.")
}
