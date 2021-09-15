package expersist

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func SaveTxt(users []User) {
	FILE := "users.txt"
	f, err := os.Open(FILE)
	if err != nil {
		f, err = os.Create(FILE)
	}

	var w *bufio.Writer
	w = bufio.NewWriter(f)
	for _, i := range users {
		c, err := json.Marshal(i)
		if err != nil {
			log.Println(err)
			return
		}
		w.Write(c)
		w.WriteString("\n")
	}
	w.Flush()
}

func SaveJSON(users []User) {
	content, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile("users.json", content, 0644)
	if err != nil {
		log.Println(err)
		return
	}
}

func SaveCSV(users []User) {
	FILE := "users.csv"
	f, err := os.Open(FILE)
	if err != nil {
		f, err = os.Create(FILE)
	}
	header := []string{"Username", "Email", "Role"}
	var values [][]string
	for _, i := range users {
		var line []string
		line = append(line, i.Username, i.Email, i.Role)
		values = append(values, line)
	}
	w := csv.NewWriter(f)
	w.Write(header)
	for _, i := range values {
		w.Write(i)
	}
	w.Flush()
	err = w.Error()
	if err != nil {
		log.Panicln(err)
		return
	}
}
