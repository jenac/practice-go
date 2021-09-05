package main

import (
	"errors"
	"fmt"
	"unicode"
)


func UnicodeUsage() {
	var ss = "你好, goland 123"
	for _, i := range ss {
		if unicode.IsLetter(i) {
			fmt.Printf("Yes: %c", i)
		} else {
			fmt.Printf("No: %c", i)
		}
	}
	fmt.Print("\n")
	for _, i := range ss {
		fmt.Printf("%c", unicode.ToUpper(i))
	}
	fmt.Print("\n")

}

func RegisterUserName(name string, table *unicode.RangeTable)  error {
	for _, i := range name {
		if !unicode.Is(table, i) {
			return errors.New("scripts is not correct")
		}
	}
	return nil

}

func main() {
	UnicodeUsage()
	fmt.Println("----------")
	fmt.Println(RegisterUserName("注册名hello", unicode.Scripts["Han"]))
}