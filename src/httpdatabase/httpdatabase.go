package main

//简单的JSON Restful API演示(服务端)
//author: Xiong Chuan Liang
//date: 2015-2-28

import (
	//"encoding/json"
	//"flag"
	"fmt"
	//"net/http"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/pflag"
)

var (
	argPort         = pflag.Int("port", 9080, "The port to listen to for incoming HTTP requests")
	argDatabaseHost = pflag.String("database-host", "", "The address is the backend database address, eg. mysql. "+
		"address:port. If not specified, the assumption is that the database is running locally. ")
	argUsername = pflag.String("username", "", "The username of the user to login to the mysql.")
	argPassword = pflag.String("password", "", "The password of the mysql user.")
)

func main() {
	pflag.Parse()
	fmt.Println("The argPort is", *argPort)
	fmt.Println("The database-host is", *argDatabaseHost)
	fmt.Println("The database-name is", *argDatabaseName)
}
