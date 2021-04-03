package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

func (s *APIServer) MainPage(w http.ResponseWriter, r *http.Request) {

}

func (s *APIServer) CreateGamePage(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles(s.config.Uixpath + "create_game.html")

	if err != nil {
		panic(err)
	}

	page.Execute(w, nil)
}

func (s *APIServer) Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func (s *APIServer) GetQuestion(w http.ResponseWriter, r *http.Request) {
	idx, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(idx)
	fmt.Fprint(w, "OK")
}

func (s *APIServer) Start() {
	router := mux.NewRouter()

	s.storage.LoadIntoMemory()

	router.HandleFunc("/", s.MainPage)
	router.HandleFunc("/Ping", s.Ping)

	router.HandleFunc("/GetQuestion/{id}", s.GetQuestion).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir(s.config.Staticpath))))

	http.ListenAndServe(s.config.Address, router)
}
