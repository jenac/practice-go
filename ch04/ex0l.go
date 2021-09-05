package main

import (
	"fmt"
	"net/url"
)

func UrlUsage() {
	var urlString = "https://golang.org/pkg/net/url?name=james&age=20"
	urlPath, _ :=url.Parse(urlString)
	fmt.Println(fmt.Sprintf("%#v", urlPath))

	fmt.Println(urlPath.Query())
	v:=urlPath.Query()
	v.Del("name")
	v.Add("school", "google")
	urlPath.RawQuery = v.Encode()
	fmt.Println(urlPath)
}

func UrlValues() {
	values:= "name=jame&age=20&school=google"
	v, _:=url.ParseQuery(values)
	fmt.Println(v)
	v.Add("foo", "bar")
	fmt.Println(v)
}
func main() {
	UrlUsage()
	fmt.Println("----------")
	UrlValues()
}