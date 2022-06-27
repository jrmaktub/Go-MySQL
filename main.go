package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go My SQL Test")

	//open may validate  arguments
	//To verify, we call Ping
	//func Open(driverName, dataSourcstring)(*DB error)
	//dataSourceName

	db, err := sql.Open("mysql", "root:kurt0916@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		fmt.Println("error validatinng sql.Open arguments")
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error with DB.PING")
		panic(err.Error())
	}

	insert, err := db.Query("INSERT INTO `testdb`.`students` (`id`, `firstname`, `lastname`) VALUES('2', 'BEN', 'Ford');")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	fmt.Println("succesfully connected to my SQL Database")
}
