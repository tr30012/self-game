package main

type Question struct {
	Id       int
	Cost     int
	Text     string
	Answer   string
	Event    int
	Answered bool
	VolumeId int
}

type Volume struct {
	Id        int
	Text      string
	Questions []Question
}

type Player struct {
	Id     int
	Name   string
	Points int
}

type AdminPageContent struct {
}

type CreatePageContent struct {
}

type GamePageContent struct {
}
