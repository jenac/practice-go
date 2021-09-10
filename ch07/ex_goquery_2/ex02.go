package exgoquery2

import (
	"context"
	"log"
	"strings"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/PuerkitoBio/goquery"
)

 type ResultForQuotes struct {
	Text string `json:"text"`
	Author string `json:"author"`
	Tags []string `json:"tags"`
}

func GetQuotesContentByClick(url string) []ResultForQuotes {
	// 获取网页渲染之后的源代码和下一页￼
	content, next := ClickNext(url)
	var results []ResultForQuotes
	// 解析当前页并抓取所需的内容￼
	getResults := func(content string) []ResultForQuotes {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
		if err != nil {
			log.Println(err)
			return nil
		}
		doc.Find(`body > div[class="container"] > div[class="quote"]`).
			Each(func(i int, selection *goquery.Selection) {
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
					Text:   selection.Find("span").Eq(0).Text(),
					Author: selection.Find("span > small").Text(),
					Tags:   tags(),
				}
				results = append(results, one)
				fmt.Println(one)
			})
		return results
	}
	results = append(results, getResults(content)...)
		fmt.Println(next)
		// 如果有下一页，就重复解析操作，即进行递归调用
		if strings.Contains(next, "http://") {
			GetQuotesContentByClick(next)
		} else {
			return results
		}
        return results
    }
	
	func ClickNext(url string) (string, string) {
		var nextPage map[string]string
		var pageSource string
		ctx1, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
		defer cancel()
		ctx, cancel1 := context.WithTimeout(ctx1, 30*time.Second)
		defer cancel1()
		// 若有翻页标签，则获取下一页的URL￼
		err := chromedp.Run(ctx, chromedp.Tasks{
			chromedp.Navigate(url), // 导航至新的一页￼
			chromedp.WaitVisible(".footer", chromedp.ByQuery), // 等待渲染结束￼
			chromedp.OuterHTML("body", &pageSource), // 获取网页源代码￼
			chromedp.Attributes(`li[class="next"] > a`, &nextPage), // 获取下一页，从属性值内提取￼
			chromedp.Click(`li[class="next"] > a`, chromedp.ByQuery), // 单击进入下一页￼
		})
        if err != nil {
			// 不存在下一页的情况，仍然需要返回最后一页的网页源代码￼
			err := chromedp.Run(ctx, chromedp.Tasks{
				chromedp.Navigate(url),
				chromedp.WaitVisible(".footer", chromedp.ByQuery),
				chromedp.OuterHTML("body", &pageSource),
				chromedp.WaitNotPresent(`li[class="next"]`, chromedp.ByQuery), 
				// 判断是否是最后一页，根据最后一页是否存在<li class="next">标签来判断￼
			})
			if err != nil {
				return pageSource, ""
			}
		}
		if _, ok := nextPage["href"]; ok {
			return pageSource, fmt.Sprintf("http://quotes.toscrape.com" + nextPage["href"])
		}
		return pageSource, ""
	}