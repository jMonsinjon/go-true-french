package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jMonsinjon/go-true-french/handler"
	"github.com/jMonsinjon/go-true-french/services"
)

func main() {

	services.InitProverbs()
	services.InitCommonNames()
	services.InitQuestions()

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
	})
	r.HandleFunc("/welcome", handler.Welcome).Methods("GET")
	r.HandleFunc("/commonPoint", handler.CommonPoint).Methods("GET")
	r.HandleFunc("/proverb", handler.Proverb).Methods("GET")
	r.HandleFunc("/question", handler.Question).Methods("GET")

	r.PathPrefix("/statics/").Handler(http.StripPrefix("/statics/", http.FileServer(http.Dir("./public/assets"))))

	port := flag.String("p", "8080", "port to serve on")
	log.Printf("Serving on HTTP port: %s\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, r))
}
