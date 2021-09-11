package exhtmlquery

import (
	"log"
	"testing"
	"fmt"
)
func TestParseWeiBo(t *testing.T) {
	content, err := GetContent(WeiBoRoot)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(content))
	ParseWeiBo(content)
}