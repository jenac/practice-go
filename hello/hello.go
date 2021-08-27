package main

import (
	"fmt"

	"example.com/greetings"
	"rsc.io/quote"
)


func main()  {
	fmt.Println("Hello, Hi")
	fmt.Println(quote.Go())
	message:=greetings.Hello("James")
	fmt.Println(message)
}