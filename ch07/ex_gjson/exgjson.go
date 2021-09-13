//go get -u github.com/tidwall/gjson
package exgjson

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

var rateTime = time.Tick(time.Microsecond * 200)

func GetContent(url string) ([]byte, error) {
	<-rateTime
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}


func V2ex(url string) {
	content, err := GetContent(url) 
	if err!=nil {
		panic(err)
	}

	doc := gjson.ParseBytes(content).Array()
	for _, i := range doc {
		var r ResultForV2ex
		r = ResultForV2ex {
			Title: i.Get("title").String(),
			URL: i.Get("url").String(),
			Description: strings.TrimSpace(i.Get("content").String()),
			Content: strings.TrimSpace( i.Get("content_rendered").String()),
		}
		fmt.Println(r)
	}
}