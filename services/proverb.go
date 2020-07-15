package services

import (
	"io/ioutil"
	"log"
	"math/rand"

	"gopkg.in/yaml.v2"
)

type Proverb struct {
	Begining string `yaml:"begining"`
	End      string `yaml:"end"`
}

var proverbs []Proverb

func InitProverbs() {
	yamlContent, err := ioutil.ReadFile("./data/proverbs.yml")
	if err != nil {
		log.Printf("Unable to read proverb data file: #%v ", err)
	}

	err = yaml.Unmarshal(yamlContent, &proverbs)
	if err != nil {
		log.Fatalf("Unable to unmarshall proverb data file: %v", err)
	}
}

func GetProverb() Proverb {
	return proverbs[rand.Intn(len(proverbs))]
}
