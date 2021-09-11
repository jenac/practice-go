//go get github.com/antchfx/htmlquery

package exhtmlquery

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
)

var rateTime = time.Tick(time.Millisecond * 200)

func GetContent(url string) ([]byte, error) {
	<-rateTime
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer response.Body.Close()
	fmt.Println("*********")
	return ioutil.ReadAll(response.Body)
}

func ToInt(value string) int {
	if value == "" {
		return 0
	}
	v := strings.TrimSpace(value)
	i, _ := strconv.Atoi(v)
	return i
}

func ParseWeiBo(content []byte) {
	reader := strings.NewReader(string(content))
	doc, err := htmlquery.Parse(reader)
	if err != nil {
		log.Panicln(err)
		return
	}

	tds := htmlquery.Find(doc, `//div[@id="pl_top_realtimehot"] //tbody/tr/td[2]`)
	for index, i := range tds {
		if index == 0 {
			continue
		}
		a := htmlquery.FindOne(i, "/a")
		if len(a.Attr) > 2 {
			continue
		}
		// 提取标题
		aText := htmlquery.InnerText(a)
		// 提取链接
		aHref := htmlquery.InnerText(htmlquery.FindOne(a, "/@href"))
		var result ResultForWeiBo
		result = ResultForWeiBo{
			Title: strings.TrimSpace(aText),
			Url:   fmt.Sprintf("%s%s", HOST, strings.TrimSpace(aHref)),
		}
		// 提取分数
		span := htmlquery.FindOne(i, "/span")
		if span != nil {

			result.Score = ToInt(htmlquery.InnerText(span))
		}
		fmt.Println(result)
	}
}
