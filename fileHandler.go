package main

import (
	"bufio"
	"math/rand"
	"strings"
)

// FileHandler structure holds question data
type FileHandler struct {
	data              string
	allQuestions      map[int]string
	allAnswers        map[int]string
	usedQuestions     map[int]bool
	answeredQuestions map[int]bool
	numCorrect        int
	numTotal          int
	numUsed           int
}

// ResetEverything clears question bank and resets data
func (fileHandler *FileHandler) ResetEverything() {
	fileHandler.allQuestions = map[int]string{}
	fileHandler.allAnswers = map[int]string{}

	if fileHandler.data != "" {
		sections := strings.Split(fileHandler.data, "\n\n")

		for _, section := range sections {
			section = strings.TrimSpace(section)

			if section == "" {
				continue
			}

			question := ""
			answer := ""

			scanner := bufio.NewScanner(strings.NewReader(section))
			for scanner.Scan() {
				text := strings.TrimSpace(scanner.Text())

				if strings.Contains(strings.ToLower(text), "answer:") {
					answer = strings.ToUpper(strings.Replace(strings.ToLower(text), "answer:", "", 1))
					break
				} else {
					question += text + "\n"
				}
			}

			fileHandler.allQuestions[len(fileHandler.allQuestions)+1] = question
			fileHandler.allAnswers[len(fileHandler.allAnswers)+1] =
				strings.Trim(strings.Trim(strings.Trim(answer, " "), "\n"), "\r")
		}
	}

	fileHandler.usedQuestions = map[int]bool{}
	fileHandler.answeredQuestions = map[int]bool{}
	fileHandler.numCorrect = 0
	fileHandler.numTotal = len(fileHandler.allQuestions)
	fileHandler.numUsed = 0
}

// ParseFileData parses question bank and creates question and answers
func (fileHandler *FileHandler) ParseFileData(data string) {
	fileHandler.data = data
	fileHandler.ResetEverything()
}

// RandomQuestion provides a random questions
func (fileHandler *FileHandler) RandomQuestion() (int, string) {
	if len(fileHandler.allQuestions) == len(fileHandler.usedQuestions) {
		return -1, "All Questions used! Please reset or choose new file"
	}
	for {
		i := rand.Intn(len(fileHandler.allQuestions)-1) + 1
		if _, exists := fileHandler.usedQuestions[i]; !exists {
			fileHandler.usedQuestions[i] = true
			fileHandler.numUsed++
			return i, fileHandler.allQuestions[i]
		}
	}
}

// GetFeedback provodes feedback based on the provided user answer
func (fileHandler *FileHandler) GetFeedback(id int, answer string) string {
	if id > 0 {
		correctAnswer := fileHandler.allAnswers[id]

		feedback := "Correct Answer is " + correctAnswer + "\nYour answer is "
		correct := correctAnswer == answer

		if correct {
			feedback += "Correct!"

			if _, exists := fileHandler.answeredQuestions[id]; !exists {
				fileHandler.numCorrect++
			}
		} else {
			feedback += "Incorrect"
		}

		fileHandler.answeredQuestions[id] = true
		return feedback
	} else {
		return "Invalid question"
	}
}
