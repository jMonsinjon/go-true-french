package services

import (
	"io/ioutil"
	"log"
	"math/rand"

	"gopkg.in/yaml.v2"
)

var commonNames []string

func InitCommonNames() {
	yamlContent, err := ioutil.ReadFile("./data/common-points.yml")
	if err != nil {
		log.Printf("Unable to read common-points data file: #%v ", err)
	}

	err = yaml.Unmarshal(yamlContent, &commonNames)
	if err != nil {
		log.Fatalf("Unable to unmarshall common-points data file: %v", err)
	}
}

func GetCommonNames() (string, string) {
	return commonNames[rand.Intn(len(commonNames))], commonNames[rand.Intn(len(commonNames))]
}
