package ex706

import (
	"fmt"
	"log"
	"testing"
)

func TestGetRespone(t *testing.T) {
	content, err := GetResponse("http://www.baidu.com")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(content))
}

func TestGetContent(t *testing.T) {
	content, err := GetContent("http://www.bidu.com")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(content))
}