package main

import (
	"fmt"
	"reflect"
)

var opList = func(number [4]int) {
	fmt.Println(number[1])
	fmt.Println(len(number))
	fmt.Println(number[1:], reflect.TypeOf(number[1:]))
	for index, one := range number {
		fmt.Println(index, one)
	}
	for i := 0; i < len(number); i++ {
		fmt.Println(i, number[i])
	}
}

var opSlice = func(name []string) []string {
	fmt.Println(name[1], reflect.TypeOf(name[1]))
	for index, one := range name {
		fmt.Println(index, one)
	}
	name = append(name, "appended")
	return name
}

var opMap = func (name map[string] int)  {
	for key, val := range name {
		fmt.Println(key, val)
	}

	name["Life"] = 100

	if val, ok := name["Go"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("no exists Go")
	}

	delete(name, "java")
	fmt.Println("delete java1")
	delete(name, "java1")
	fmt.Println(name)
}

func main() {
	var number [4]int= [...] int {1,2,3,4}
	fmt.Println(reflect.TypeOf(number))
	opList(number)

	var name[]string= [] string {
		"go", "python", "java", "c++", "c#",
	}
	fmt.Println(opSlice(name))

	nameMap := make(map[string]int)
	nameMap["java"] = 100
	nameMap["php"] = 200
	nameMap["python"] = 300
	nameMap["js"] = 100
	opMap(nameMap)
}
