package handler

import (
	"html/template"
  "net/http"

	proverbService "github.com/jMonsinjon/go-true-french/services"
)

type ProverbData struct {
	Title    string
	Question string
	Answer   string
}

func Proverb(w http.ResponseWriter, r *http.Request) {

	proverb := proverbService.GetProverb()

	tmpl := template.Must(template.ParseFiles("public/question.html"))
	data := ProverbData{
		Title: "Proverbe",
		Question: proverb.Begining,
		Answer:   proverb.End,
	}
	tmpl.Execute(w, data)
}