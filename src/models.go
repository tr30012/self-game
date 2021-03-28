package main

type Player struct {
	Name   string `json:"p_name"`
	Points int    `json:"p_points"`
}

type Question struct {
	Id       int
	Cost     int    `json:"q_cost"`
	Text     string `json:"q_text"`
	Answer   string `json:"q_answer"`
	Event    int    `json:"q_event"`
	Answered bool   `json:"q_answered"`
}

type Volume struct {
	Text      string     `json:"v_text"`
	Questions []Question `json:"v_questions"`
}

type QuestionMap map[string](map[int]Question)

type MainPageContent struct {
	Players   []Player
	Questions QuestionMap
}
