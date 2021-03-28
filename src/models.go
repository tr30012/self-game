package main

type Question struct {
	Cost     int    `json:"q_cost"`
	Text     string `json:"q_text"`
	Answer   string `json:"q_answer"`
	Event    int    `json:"q_event"`
	Answered bool   `json:"q_answered"`
}
