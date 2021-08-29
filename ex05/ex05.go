package main

import (
	"fmt"
	"unsafe"
)

type Info struct {
	Name string
	Age  int
}

type Student struct {
	Name string
	University
}

func (s Student) PrintName() {
	fmt.Println(s.Name)
}

type University struct {
	Name string
	Location string
}

func (u University) PrintName() {
	fmt.Println(u.Name)
}

func main() {
	var infoOne Info = Info{
		Name: "this-name",
		Age:  32,
	}

	fmt.Println(unsafe.Sizeof(infoOne))
	fmt.Printf("Name address: %x, Size: %d\r\n",
		&infoOne.Name, unsafe.Sizeof(infoOne.Name))
	fmt.Printf("Age address:%x, Size: %d\r\n",
		&infoOne.Age, unsafe.Sizeof(infoOne.Age))

	fmt.Println("----------")
	var std Student
	std.Name = "Zhao the Ninjia"
	std.University.Name = "Princeton University"
	std.University.Location = "New Jersey"
	fmt.Println(std)
	std.PrintName()
	std.University.PrintName()
}