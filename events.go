package main

import (
	"github.com/Equanox/gotron"
	"log"
	"strconv"
)

const (
	FEEDBACK_EVENT = "feedback"
	SHOW_RANDOM_EVENT = "show_random"
)


// RandomQuestion message type
type RandomQuestionEvent struct {
	*gotron.Event
	Data string `json:"data"`
	ID   string `json:"id"`
}

// Feedback message type
type FeedbackEvent struct {
	*gotron.Event
	Feedback string `json:"feedback"`
	Attempted string `json:"attempted"`
	Correct string `json:"correct"`
	Total string `json:"total"`
	Used string `json:"used"`
}

type windowStuff struct {
	window *gotron.BrowserWindow
}

var windowStuffImpl = windowStuff{}

func SetWindow(window *gotron.BrowserWindow) {
	windowStuffImpl.window = window
}

func SendFeedbackEvent(feedback string, attempted, correct, total, used int) {
	if err := windowStuffImpl.window.Send(&FeedbackEvent{
		Event: &gotron.Event{Event: FEEDBACK_EVENT},
		Feedback:    feedback,
		Attempted: strconv.Itoa(attempted),
		Correct: strconv.Itoa(correct),
		Total: strconv.Itoa(total),
		Used: strconv.Itoa(used),
	}); err != nil {
		log.Println("[Window] ", err)
	}
}

func SendRandomQuestionEvent(id int, question string) {
	if err := windowStuffImpl.window.Send(&RandomQuestionEvent{
		Event: &gotron.Event{Event: SHOW_RANDOM_EVENT},
		ID:    strconv.Itoa(id),
		Data:  question,
	}); err != nil {
		log.Println("[Window] ", err)
	}
}