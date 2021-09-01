package main

import "fmt"

type Val struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func (v Val) String() string {
	return fmt.Sprintf("%s + %d", v.Name, v.Age)
}

func (v Val) GoString() string {
	return fmt.Sprintf("go %d + %s", v.Age, v.Name)
}

func main() {

	a := Val {
		Name: "GoLang",
		Age: 20,
	}
	fmt.Println(a)
	fmt.Printf("%#v\n", a)
}