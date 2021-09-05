package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type Numbers struct {
	Num []int
}

func (n *Numbers) Set(value string) error {
	sList := strings.Split(value, "|")
	var num []int
	for _, i := range sList {
		in, _ := strconv.Atoi(i)
		num = append(num, in)
	}
	n.Num = num
	return nil
}

func (n *Numbers) String() string {
	return fmt.Sprintf("%#v", n.Num)
}

func FlagSpecial() {
	var n Numbers
	flag.Var(&n, "n", "numbers to parse")
	flag.Parse()
	fmt.Println(n.Num)
}

func main() {
	FlagSpecial()
}
