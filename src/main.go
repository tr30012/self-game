package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	webAddress = "127.0.0.1:8080"
)

var (
	logger *Logger
	router *mux.Router
)

func main() {
	logger = NewLogger(os.Stdout, os.Stdout, os.Stderr, llInfo)
	logger.Info("Web-address of server is " + webAddress)

	router = mux.NewRouter()
	router.HandleFunc("/", mainPage)

	if err := http.ListenAndServe(webAddress, router); err != nil {
		panic(err)
	}
}
