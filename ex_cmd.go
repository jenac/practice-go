// go run ex_cmd.go jen jenac@s.com Amazon
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type userInfoOs struct {
	File    string `json:"file"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Company string `json:"company"`
}

func PrintCmdOS() {
	args := os.Args
	if len(args) != 4 {
		fmt.Println("you need name, email, company")
		return
	}

	var oneUserInfoOs userInfoOs
	oneUserInfoOs.File = os.Args[0]
	oneUserInfoOs.Name = os.Args[1]
	oneUserInfoOs.Email = os.Args[2]
	oneUserInfoOs.Company = os.Args[3]

	jsonByte, _ := json.MarshalIndent(oneUserInfoOs, "", "  ")
	fmt.Println(string(jsonByte))
}

func main() {
	PrintCmdOS()
}
