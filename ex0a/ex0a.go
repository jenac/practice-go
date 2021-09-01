package main

import "fmt"

func FmtStringUsage() {
	values := "golang"
	fmt.Printf("%s\n", values)
	fmt.Printf("%q\n", values)
}

func FmtBoolUsage() {
	ok := true
	fmt.Printf("%t\n", ok)
}

func FmtOtherUsage() {
	a:=1
	b:=2.0
	ok:=true
	number:=&a
	s:=struct {
		Name string `json:"name"`
	} {
		Name: "Go",
	}

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", ok)
	fmt.Printf("%p, %d\n", &a, number)
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s)

}
func main() {
	number := 100.203
	numberInt := 100
	fmt.Printf("%d\n", numberInt)
	fmt.Printf("%o\n", numberInt)
	fmt.Printf("%x\n", numberInt)
	fmt.Printf("%b\n", numberInt)
	fmt.Println("----------")
	fmt.Printf("%f\n", number)
	fmt.Printf("%e\n", number)
	fmt.Printf("%E\n", number)
	fmt.Println("----------")
	FmtStringUsage()
	fmt.Println("----------")
	FmtBoolUsage()
	fmt.Println("----------")
	FmtOtherUsage()
}
