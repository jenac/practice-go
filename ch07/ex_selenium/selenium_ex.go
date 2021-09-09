package exselenium

import (
	"fmt"

	"github.com/tebeka/selenium"
)

const (
	PORT = 4444
)

func SeleniumGetContent(url string) (string, error) {
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	webDriver, err := selenium.NewRemote(caps, 
		fmt.Sprintf("http://localhost:%d/wd/hub", PORT))
	if err != nil {
		panic(err)
	}
	if err:= webDriver.Get(url); err != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err))
	}
	return webDriver.PageSource()
}
