package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:123@tcp(mysql:3306)/test")
	if err != nil {
		panic(err)
	}
}

func main() {
	createTable()
	insertTable()
	id := getID()
	for {
		fmt.Println("the id from test_table is: ", id)
		time.Sleep(1 * time.Second)
	}
}

func createTable() {
	stmt, err := db.Prepare("create table test_table(id int(12))")
	defer stmt.Close()
	if err != nil {
		panic(err)
	}
	stmt.Exec()
}

func insertTable() {
	stmt, err := db.Prepare("insert into test_table(id) values(2017)")
	defer stmt.Close()
	if err != nil {
		panic(err)
	}
	if _, err = stmt.Exec(); err != nil {
		panic(err)
	}
}

func getID() (id int64) {
	if err := db.QueryRow("select * from test_table").Scan(&id); err != nil {
		panic(err)
	}
	return
}
