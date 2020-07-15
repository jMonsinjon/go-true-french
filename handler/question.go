package handler

import (
	"html/template"
  "net/http"

	questionService "github.com/jMonsinjon/go-true-french/services"
)

type QuestionData struct {
	Title    string
	Question string
	Answer   string
}

func Question(w http.ResponseWriter, r *http.Request) {
	question := questionService.GetQuestion()

	tmpl := template.Must(template.ParseFiles("public/question.html"))
	data := ProverbData{
		Title: "Question",
		Question: question.Question,
		Answer:   question.Answer,
	}
	tmpl.Execute(w, data)
}
