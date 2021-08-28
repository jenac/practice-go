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

func main() {
	var number [4]int= [...] int {1,2,3,4}
	fmt.Println(reflect.TypeOf(number))
	opList(number)

	var name[]string= [] string {
		"go", "python", "java", "c++", "c#",
	}
	fmt.Println(opSlice(name))
}
