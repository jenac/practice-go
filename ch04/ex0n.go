package main

import (
	"log"
	"net/http"
)

type SelfHandler struct {
}

func (SelfHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("Hello, python"))
}
func main() {
	http.HandleFunc("/hello_golang",
		func(rw http.ResponseWriter, r *http.Request) {
			rw.Write([]byte("Hello, GoLang"))
		})
	var self SelfHandler
	http.Handle("/hello_python", self)
	log.Fatal(http.ListenAndServe(":8011", nil))
}
