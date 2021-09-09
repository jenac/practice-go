package exchromedp

import (
	"fmt"
	"testing"
)

func TestChromedpGetContent(t *testing.T) {
	content := ChromedpGetContent("http://quotes.toscrape.com/js/")
	fmt.Println(content)

}
