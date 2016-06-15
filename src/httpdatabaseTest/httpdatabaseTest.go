package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	. "httpdatabase/mysql"
)

func TestInsert() {
	user := User{"test1", "test1"}
	db, err := sql.Open("mysql", "root:123456@/zjw")
	if err != nil {
		panic("mysql open error")
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic("mysql connect is wrong")
	}

	err = Insert(db, user)
	if err != nil {
		panic("mysql insert error")
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	TestInsert()
}
