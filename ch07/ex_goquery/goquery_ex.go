//go get github.com/PuerkitoBio/goquery
package exgoquery

import (
	"context"
	"log"
	"strings"
	"fmt"

	"github.com/chromedp/chromedp"
	"github.com/PuerkitoBio/goquery"
)

func ChromedpGetContent(url string) string {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()
	var response string
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.OuterHTML("body", &response),
	})
	if err != nil {
		log.Println(err)
		return ""
	}
	return response
}


type ResultForQuotes struct {
	Text string `json:"text"`
	Author string `json:"author"`
	Tags []string `json:"tags"`
}

func GetQuotesContent(url string) []ResultForQuotes {
	content:=ChromedpGetContent(url)

	doc, err:=goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		log.Println(err)
		return nil
	}
	
	var results []ResultForQuotes

	doc.Find(`body > div[class="container"] > div[class="quote"]`).
		Each(func(i int, selection *goquery.Selection){
			var one ResultForQuotes
			tags := func() []string {
				var ts []string
				selection.Find("div > a").
					Each(func(i int, selection *goquery.Selection) {
						ts = append(ts, selection.Text())
					})
				return ts
			}
			one = ResultForQuotes{
				Text: selection.Find("span").Eq(0).Text(), //div下第一个span标签的文本
				Author: selection.Find("span > small").Text(), // div下带子节点small标签的文本￼
				Tags:   tags(),
			}
			results = append(results, one)
			fmt.Println(one)
		})
	return results
}
		