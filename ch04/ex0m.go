package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func getHandle(rawString string) {
	response, err := http.Get(rawString)
	if err != nil {
		return
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func ClientUsageGet() {
	getHandle("https://google.com/headers")
}

func postHandle(rawString string, body io.Reader) {
	response, err := http.Post(rawString, "application/json", body)
	if err != nil {
		return
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func ClientUsagePost() {
	var buf bytes.Buffer
	buf.WriteString("Hello, GoLang")
	postHandle("https://google.com", &buf)

	val:=strings.NewReader("Hello, GOLang")
	postHandle("https://google.com", val)

	bytesNew := struct {
		PortfolioName string `json:"portfolioName"`
		LossPercent int `json:"lossPercent"`
		NoGainSince string `json:"noGainSince"`
		NoGainPercent int `json:"noGainPercent"`
	} {
		PortfolioName: "Watch_Nop",
		LossPercent: -2,
		NoGainSince: "2021-01-01",
		NoGainPercent: 10,
	}

	fmt.Println("*********")
	byt, _:=json.Marshal(bytesNew)
	postHandle("http://192.168.68.96:3333/alerts", bytes.NewReader(byt))
	fmt.Println("*********")

	//post form
	response, err := http.PostForm("http://google.com/", url.Values{
		"name": []string {"GoLang"},
		"age": []string {"10"},
	})
	if err != nil {
		fmt.Println(err)
		return 
	}
	defer response.Body.Close()
	content, _:= ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
func main() {
	ClientUsageGet()
	fmt.Println("----------")
	ClientUsagePost()

}
