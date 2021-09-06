//go get -u github.com/gobuffalo/packr
package main

import (
	"log"
	"os"
	"text/template"

	"github.com/gobuffalo/packr"
)
type Detail struct {
	Class string
	Title string
	Content string
}

var detail Detail

func init() {
	detail = Detail{
		Class: "test-class",
		Title: "test-title",
		Content: "test-content",
	}
}

func withPackr() {
	box:=packr.NewBox(".")
	index, err := box.FindString("index.html")
	if err != nil {
		log.Println(err)
		return
	}

	t, _ := template.New("").Parse(index)
	err = t.Execute(os.Stdout, &detail)
	if err!= nil {
		log.Println(err)
		return
	}
}

func main() {
	withPackr()
}