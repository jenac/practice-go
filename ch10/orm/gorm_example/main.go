package main

import (
	"ch10/gorm_example/model"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var engine *gorm.DB

func init() {
	var err error
	engine, err = gorm.Open("mysql", "root:admin123@tcp(127.0.0.1:3306)/ch10_gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Print(err)
		panic("gorm connect to mysql failed")
	}
	engine.LogMode(true)
	syncTables()
	defer engine.Close()
}

func syncTables() {
	engine.DropTableIfExists(&model.Person{}) 
	engine.AutoMigrate(&model.Person{})
}

func main() {
	fmt.Println("Hello")
}