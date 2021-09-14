package exmaoyan

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

var getContent = func(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Origin", "http://piaofang.maoyan.com")
	req.Header.Set("Referer", "http://piaofang.maoyan.com/dashboard")
	req.Header.Set("User-Agent", "My Agent")
	content, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	return ioutil.ReadAll(content.Body)
}

func MaoYan(url string) {
	c, _ := getContent(url)
	doc := gjson.ParseBytes(c).Get("data.list").Array()
	for _, i := range doc {
		var r ResultForMaoYan
		r = ResultForMaoYan{
			AvgSeatView:  i.Get("avgSeatView").String(),
			AvgShowView:  i.Get("avgShowView").String(),
			BoxInfo:      i.Get("boxInfo").String(),
			BoxRate:      i.Get("boxRate").String(),
			MovieName:    i.Get("movieName").String(),
			ReleaseInfo:  i.Get("releaseInfo").String(),
			SplitBoxRate: i.Get("splitBoxRate").String(),
			SumBoxInfo:   i.Get("sumBoxRate").String(),
			ShowInfo:     i.Get("showInfo").String(),
			ShowRate:     i.Get("showRate").String(),
		}
		fmt.Println(r)
	}
}
