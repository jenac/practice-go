package ex706

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetResponse(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return content, nil
}

var rateTime = time.Tick(time.Microsecond * 200)

func GetContent(url string) ([]byte, error) {
	<-rateTime
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("User-Agent", "Customized Agent")
	response, err := http.DefaultClient.Do(request)
	if err!= nil {
		log.Println(err)
		return nil, err
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

