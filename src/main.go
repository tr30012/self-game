package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/gorilla/mux"
)

const (
	webAddress = "127.0.0.1:8080"
)

var (
	logger          *Logger
	router          *mux.Router
	mainPageContent *MainPageContent
)

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func LoadQuestions(si int, ei int) QuestionMap {

	f, err := os.Open("static\\json\\questions.json")

	if err != nil {
		panic(err)
	}

	fContent, err := ioutil.ReadAll(f)

	if err != nil {
		panic(err)
	}

	var v []Volume
	json.Unmarshal(fContent, &v)

	defer f.Close()

	q := make(QuestionMap)

	id := 0
	for i := si; i < ei; i++ {
		q[v[i].Text] = make(map[int]Question)

		for j := 0; j < len(v[i].Questions); j++ {
			v[i].Questions[j].Id = id
			v[i].Questions[j].Answered = false
			q[v[i].Text][v[i].Questions[j].Cost] = v[i].Questions[j]
			id++
		}
	}

	logger.Info("Успешно загружены вопросы.")

	return q
}

func LoadPlayers() []Player {
	f, err := os.Open("static\\json\\players.json")

	if err != nil {
		panic(err)
	}

	fContent, err := ioutil.ReadAll(f)

	if err != nil {
		panic(err)
	}

	var v []Player
	json.Unmarshal(fContent, &v)

	for i := 0; i < len(v); i++ {
		v[i].Id = i
	}

	defer f.Close()

	logger.Info("Успешно загружены игроки.")

	return v
}

func LoadMainPageContent() *MainPageContent {
	return &MainPageContent{
		Questions: LoadQuestions(0, 5),
		Players:   LoadPlayers(),
	}
}

func main() {
	logger = NewLogger(os.Stdout, os.Stdout, os.Stderr, llInfo)
	logger.Info("Web-address of server is " + webAddress)

	router = mux.NewRouter()

	router.HandleFunc("/", mainPage).Methods("GET")
	router.HandleFunc("/get", getPageContent).Methods("GET")
	router.HandleFunc("/set", setPageContent).Methods("POST")
	router.HandleFunc("/restart", restartPageContent).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("./static/"))))

	mainPageContent = LoadMainPageContent()

	go open("http://" + webAddress)

	if err := http.ListenAndServe(webAddress, router); err != nil {
		panic(err)
	}
}
