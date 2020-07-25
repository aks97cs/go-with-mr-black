package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Shop struct {
	gorm.Model
	id string
	name string
	city string
	locality string
	created string
	modified string 
	status bool
}

func main() {

	db, err := gorm.Open("mysql", "anuj:123@tcp(localhost:3306)/godb")

	if err != nil {
		panic("Failed to connect database")
	}
	
	defer db.Close()

	db.AutoMigrate(&Shop{})

	db.Create(&Shop{id: "1", name: "Anuj", city: "Varanasi", locality: "DurgaKund", created: "now", 
	modified: "now", status: true})

	var shop Shop

	d:=db.First(&shop, 1)

	fmt.Printf("data: %v", d)


}