package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// type error interface {
// 	Error() string
// }

type ErrorMessage struct {
	Err error
	Code int
	Message string
}

func (e *ErrorMessage) Error() string {
	return fmt.Sprintf(
		"e.err=%s, e.code=%d, e.messge=%s",
		e.Err.Error(), e.Code, e.Message)
}

var (
	ErrNotRoute=ErrorMessage{
		Err: errors.New("no route"),
		Code: 404,
		Message: "check route",
	}
	ErrParamNotOk = ErrorMessage{
		Err: errors.New("param not ok"),
		Code: 10000,
		Message: "check param",
	}
)

type University struct {
	Name string `json:"name"` // no space after json:
	Location string `json:"location"`
	Number float64 `json:"student_number,string"`
	President string `json:"-"`
}

type University2 struct {
	Name string `json:"name"` // no space after json:
	Location string `json:"location"`
	Number float64 `json:"student_number,string"`
	President string `json:"-"`
}

func (u University2) MarshalJSON() ([]byte, error) {
	result:=fmt.Sprintf(
		"name=%%%s, location=%%%s, population=%f, president=%s",
		u.Name, u.Location, u.Number, u.President)
	return json.Marshal(result)
}


func main() {
	fmt.Println(ErrNotRoute)
	fmt.Println(ErrParamNotOk)

	fmt.Println("----------")
	// var university University
	university := University{
		Name: "Beijing Institute of Technology",
		Location: "BeiJing",
		Number: 2000000,
		President: "Me",
	}

	universityByte, _:= json.Marshal(university)
	fmt.Println(string(universityByte))

	u2:=University2{
		Name: "Beijing Institute of Technology",
		Location: "BeiJing",
		Number: 2000000,
		President: "Me",
	}
	u2Byte, _:=json.Marshal(u2)
	fmt.Print(string(u2Byte))
}