package handler

import (
	"html/template"
	"net/http"

	commonPoint "github.com/jMonsinjon/go-true-french/services"
)

type CommonPointData struct {
	Title      string
	NameFirst  string
	NameSecond string
}

func CommonPoint(w http.ResponseWriter, r *http.Request) {
	first, second := commonPoint.GetCommonNames()
    
    tmpl := template.Must(template.ParseFiles("public/common-point.html"))
	data := CommonPointData{
		Title:      "Common point",
		NameFirst:  first,
		NameSecond: second,
	}
	tmpl.Execute(w, data)
}
