package main

import ("fmt")

type MyInt struct {
	Number int
}

func (m MyInt) SayHello() {
	fmt.Println("hello")
}

func (m *MyInt) SetNumber(value int) {
	m.Number = value
}

func (m MyInt) SauNumber() {
	fmt.Println(m.Number)
}

func main() {
	var my MyInt
	my.SayHello()
	my.Number = 1
	my.SauNumber()
	my.SetNumber(100)
	my.SauNumber()
}