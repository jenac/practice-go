package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("hello, Go, Java", "Go"))
	fmt.Println(strings.Compare("Java", "Go"))
	fmt.Println(strings.ToUpper("goLang"))
	fmt.Println(strings.ToLower("goLang"))
	fmt.Println(strings.ToTitle("goLang"))
	fmt.Println(strings.Count("go Go go", "go"))
	fmt.Println(strings.HasPrefix("GoLang", "GoL"))
	fmt.Println(strings.HasSuffix("GoLang", "Lang"))
	splitted := strings.Split("How are you?", " ")
	fmt.Println(reflect.TypeOf(splitted))
	fmt.Println(splitted[1])
	fmt.Println(strings.Join(splitted, ","))
	fmt.Println(strings.Index("Thisisgood", "i"))
	fmt.Println(strings.TrimSpace("   how are you?   a   "))
	fmt.Println(strings.Trim("##   how are you?   a   ##", "#"))
	replacer := strings.NewReplacer("1", "one", "2", "two", "3", "three")
	fmt.Println(replacer.Replace("1 dog, 2 dogs, 3 cats"))
}