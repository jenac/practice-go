package main

import "fmt"

type Info struct {
	Name string
	_ int
	Age int
}

func main() {
	var infoOne Info = Info{
		Name: "one",
		Age: 10,
	}

	var infoTwo = Info{"two", 2000, 20}

	var infoThree = new(Info)
	infoThree = &Info {
		Name: "three",
		Age: 30,
	}

	fmt.Println("1", infoOne)
	fmt.Println("2", infoTwo)
	fmt.Println("3", *infoThree)
	fmt.Println("3", infoThree)
}