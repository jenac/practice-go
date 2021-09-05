package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println(string([]byte("A b c D")))

	A := []byte("how are you")
	fmt.Println(A)

	B := "this is a string"
	fmt.Println(B)

	fmt.Printf("%T, %T\r\n", A, B)
	JsonMarshal()
	JsonUnmarshal()
}

type JsonExample struct {
	Name   string `json:"name,omitempty"`
	Age    int    `json:"age"`
	School string `json:"university"`
}

func JsonMarshal() {
	jex := JsonExample{
		Name:   "Go",
		Age:    10,
		School: "Google",
	}
	by, _ := json.Marshal(jex)
	fmt.Println(string(by))
}

func JsonUnmarshal() {
	var v JsonExample
	by:=[]byte(`{"name":"Go", "age":12, "university":"Google"}`)
	json.Unmarshal(by, &v)
	fmt.Println(v)

	var vother JsonExample
	byOther:=[]byte(`{"name":"", "age":13, "university":"Google"}`)
	json.Unmarshal(byOther, &vother)
	fmt.Println(vother)
}