package main

import (
	"fmt"
	"log"
	"net/http"
)

type Self struct {

}

func (Self) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "Hello Self Server 2")
}

// func (Self) Say(writer http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(writer, "Hello Self Server 1")
// }

func main() {
	var selfServer http.Server

	var selfHandler Self

	var selfMux *http.ServeMux

	selfMux = &http.ServeMux{}

	selfHandler = Self{}
	selfMux.Handle("/say", selfHandler)

	selfServer = http.Server{
		Handler: selfHandler,
		Addr: "localhost:9099",
	}
	log.Fatal(selfServer.ListenAndServe())
}