package main

import (
	"fmt"
	"reflect"
)

type Info struct {
	Name string `json: "name"`
	Age	int `json:"age"`
	Number int `json:"number"`
}

func hasInfoName(value string) bool {
	var info Info
	info.Name = "some-name"
	info.Age = 20
	info.Number = 100

	var typeInfo reflect.Type
	typeInfo = reflect.TypeOf(info)

	if _, ok := typeInfo.FieldByName(value); ok {
		return ok
	}
	return false
}

func main() {
	fmt.Println(hasInfoName("Name"))
	fmt.Println(hasInfoName("Age"))
	fmt.Println(hasInfoName("Number"))
	fmt.Println(hasInfoName("What?"))
}