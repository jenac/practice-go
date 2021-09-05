package main

import (
	"fmt"
	"sort"
)

func SortIntsUsage() {
	list := []int{1, 5, 1, 2, 74, 21, 45, 13}
	sort.Ints(list)
	fmt.Println(list)
	sort.Sort(sort.Reverse(sort.IntSlice(list)))
	fmt.Println(list)
}

func SortFloatsUsage() {
	list := []float64{1.1, 5.2, 1.3, 2.4, 74.5, 21.6, 45.7, 13.8}
	sort.Float64s(list)
	fmt.Println(list)
	sort.Sort(sort.Reverse(sort.Float64Slice(list)))
	fmt.Println(list)
}

func SortStringsUsage() {
	list := []string{"a", "A", "c", "C", "b", "B"}
	sort.Strings(list)
	fmt.Println(list)
	sort.Sort(sort.Reverse(sort.StringSlice(list)))
	fmt.Println(list)
}

type Language struct {
	Year    int    `json:"year"`
	Name    string `json:"name"`
	Account string `json:"account"`
}

type Languages []Language

func (ls Languages) Len() int {
	return len(ls)
}

func (ls Languages) Less(i, j int) bool {
	return ls[i].Year < ls[j].Year
}

func (ls Languages) Swap(i, j int) {
	ls[i], ls[j] = ls[j], ls[i]
}

func SortLang() {
	list := Languages{
		{30, "Java", "Amazon"},
		{10, "GoLang", "Google"},
		{28, "Python", "Apple"},
	}
	fmt.Println("org", list)
	sort.Sort(list)
	fmt.Println("sorted", list)
}

func main() {
	SortIntsUsage()
	fmt.Println("----------")
	SortFloatsUsage()
	fmt.Println("----------")
	SortStringsUsage()
	fmt.Println("----------")
	SortLang()
}
