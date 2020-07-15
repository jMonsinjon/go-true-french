package services

import (
	"io/ioutil"
	"log"
	"math/rand"

	"gopkg.in/yaml.v2"
)

type Question struct {
	Question string `yaml:"question"`
	Answer   string `yaml:"answer"`
}

var questions []Question

func InitQuestions() {
	yamlContent, err := ioutil.ReadFile("./data/questions.yml")
	if err != nil {
		log.Printf("Unable to read questions data file: #%v ", err)
	}

	err = yaml.Unmarshal(yamlContent, &questions)
	if err != nil {
		log.Fatalf("Unable to unmarshall questions data file: %v", err)
	}
}

func GetQuestion() Question {
	return questions[rand.Intn(len(questions))]
}
