package exmaoyan

import (
	"testing"
	"fmt"
	"time"
)

func TestMaoYan(t *testing.T) {
	MaoYan(
		fmt.Sprintf("https://box.maoyan.com/promovie/api/box/second.json?beginDate=%s", 
		time.Now().Format("20210102")))
}