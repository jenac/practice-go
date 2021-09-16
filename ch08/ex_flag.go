// go run ex_flag.go -c Amazon -n JJJK 
package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

type UserInfoFlag struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	Email   string `json:"email"`
	Company string `json:"company"`
}

var (
	name    string
	url     string
	email   string
	company string
)

func PrintCommandFlag() {
	flag.StringVar(&name, "n", "James", "Show user name")
	flag.StringVar(&url, "u", "http://www.google.com", "show url")
	flag.StringVar(&email, "e", "james@s.com", "show email")
	flag.StringVar(&company, "c", "LLC", "show company")

	flag.Parse()

	var onUser UserInfoFlag
	onUser.Name = name
	onUser.Url = url
	onUser.Email = email
	onUser.Company = company
	jsonByte, _ := json.MarshalIndent(onUser, "", "  ")
	fmt.Println(string(jsonByte))
}

func main() {
	PrintCommandFlag()
}
