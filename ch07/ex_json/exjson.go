package exjson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func ParseJSON() {
	file, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Println(err)
		return
	}

	var result ResultForJSON
	err = json.Unmarshal(file, &result)

	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(result)
}

func MarshallJSON() {
	var object ResultForJSON
	object.Data.Directors = []string{"James", "Dokcer"}
	object.Data.Casts = []string{"Hello, world", "Tomorrow never die",
		"How does it work?"}
	object.Data.Title = "Good Title"
	object.Data.Rate = "7/10"
	object.Data.Star = "4"
	object.Data.Cover = 3000
	object.Data.URL = "http://www.google.com"

	content, err := json.Marshal(object)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(content))
}
