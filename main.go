package main

import (
	"encoding/json"
	"github.com/Equanox/gotron"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// FileRead message type
type FileReadEvent struct {
	*gotron.Event
	Data string `json:"data"`
}

// CheckAnswerEvent message type
type CheckAnswerEvent struct {
	*gotron.Event
	Answer string `json:"answer"`
	ID   string `json:"id"`
}

func main() {
	// Create a new browser window instance
	window, err := gotron.New("webapp")
	if err != nil {
		log.Panic(err)
	}

	// Alter default window size and window title.
	window.WindowOptions.Width = 1200
	window.WindowOptions.Height = 600
	window.WindowOptions.Title = "QuestionBank"

	done, err := window.Start()
	if err != nil {
		panic(err)
	}

	fileHandler := FileHandler{}
	SetWindow(window)
	rand.Seed(time.Now().UnixNano())

	window.On(FileReadEvent{Event: &gotron.Event{Event: "file_read"}}, func(bin []byte) {
		event := &FileReadEvent{}
		json.Unmarshal(bin, event)

		fileHandler.ParseFileData(event.Data)
		SendFeedbackEvent("", len(fileHandler.answeredQuestions),
			fileHandler.numCorrect, fileHandler.numTotal, fileHandler.numUsed)
	})

	window.On(CheckAnswerEvent{Event: &gotron.Event{Event: "check_answer"}}, func(bin []byte) {
		event := &CheckAnswerEvent{}
		json.Unmarshal(bin, event)

		event.Answer = strings.Trim(strings.Trim(strings.Trim(event.Answer, " "), "\n"), "\r")

		intID, _ := strconv.Atoi(event.ID)
		feedback := fileHandler.GetFeedback(intID, event.Answer)

		SendFeedbackEvent(feedback, len(fileHandler.answeredQuestions),
			fileHandler.numCorrect, fileHandler.numTotal, fileHandler.numUsed)
	})

	window.On(&gotron.Event{Event: "show_random"}, func(bin []byte) {
		id, question := fileHandler.RandomQuestion()

		SendRandomQuestionEvent(id, question)
		SendFeedbackEvent("", len(fileHandler.answeredQuestions),
			fileHandler.numCorrect, fileHandler.numTotal, fileHandler.numUsed)
	})


	window.On(&gotron.Event{Event: "reset_questions"}, func(bin []byte) {
		fileHandler.ResetEverything()
		SendFeedbackEvent("", len(fileHandler.answeredQuestions),
			fileHandler.numCorrect, fileHandler.numTotal, fileHandler.numUsed)

		id, question := fileHandler.RandomQuestion()
		SendRandomQuestionEvent(id, question)
	})

	<-done
}
