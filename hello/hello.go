package main

import (
	"fmt"
	"log"
	"example.com/greetings"
	"rsc.io/quote"
)


func main()  {
	fmt.Println("Hello, Hi")
	fmt.Println(quote.Go())

	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// message, err :=greetings.Hello("")
	message, err :=greetings.Hello("James")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	fmt.Println("----------")
	names:=[]string {"James", "Tim", "Adidas"}

	messages, err := greetings.Hellos(names)
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}

func init() {
	fmt.Println("hello::init is called")
}