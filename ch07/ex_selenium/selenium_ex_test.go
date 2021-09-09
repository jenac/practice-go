package exselenium

import (
	"fmt"
	"log"
	"testing"
)

func TestSeleniumGetContent(t *testing.T) {
	contentOne, err := SeleniumGetContent("http://quotes.toscrape.com/js/")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(contentOne)
}