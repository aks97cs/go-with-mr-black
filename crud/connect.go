package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	fmt.Println("working !")

	db, err := sql.Open("mysql", "anuj:123@tcp(localhost:3306)/godb")

	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println(" Not able to connect with database ")
		panic(err.Error())
	}
	// executing
	defer db.Close()

	fmt.Println("connected !")

	insert, err := db.Query("insert into test_connect values('anujsingh')")

	if err != nil{
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("inserted ")
}