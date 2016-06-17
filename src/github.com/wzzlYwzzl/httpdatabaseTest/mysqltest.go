package main

import (
	//"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/wzzlYwzzl/httpdatabase/sqlop"
)

func testmysql() {
	user := User{"zjw1", "zjw1"}
	host := MysqlCon{Host: "localhost:3306", Db: "zjw", Name: "root", Password: "123456"}
	//db, err := sql.Open("mysql", "root:123456@/zjw")
	db, err := Connect(&host)
	if err != nil {
		panic("mysql open error")
	}

	defer db.Close()

	res, err := Query(db, &user)
	if err != nil {
		panic(err)
	}
	fmt.Println(res, len(res))

	// err = Delete(db, &user)
	// if err != nil {
	// 	panic("mysql query error")
	// }

}
