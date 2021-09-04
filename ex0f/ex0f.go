package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func DefaultUsageForLog() {
	log.Print("Hello, world, GoLang")
	log.Println("Hello, world, GoLang")
	log.Printf("Hello, world, %s", "GoLang")
	log.Println("log prefix", log.Prefix())
	log.Println("log Flags", log.Flags())
}

func SpecialUsageLog() {
	logger := log.New(os.Stdout, "GoLang - ", log.Lshortfile)
	logger.Println("Hello world, GoLang")
}

func SpecialUsageWithBytes() {
	var buf bytes.Buffer
	logger := log.New(&buf, "Hi! ", log.Lshortfile)
	logger.Println("Hello world, GoLang")
	fmt.Println(buf.String())
}

func SpecialUsageWithFile() {
	file, _ := os.Create("log.log")
	logger := log.New(file, "Go - ", log.Lshortfile)
	logger.Println("Hello, world, GoLang")
}

func main() {
	DefaultUsageForLog()
	fmt.Println("----------")
	SpecialUsageLog()
	fmt.Println("----------")
	SpecialUsageWithBytes()
	fmt.Println("----------")
	SpecialUsageWithFile()
}
