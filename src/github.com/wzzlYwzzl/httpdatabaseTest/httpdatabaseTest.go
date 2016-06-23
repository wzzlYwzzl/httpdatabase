package main

import (
	"fmt"

	"github.com/wzzlYwzzl/httpdatabase/client"
)

func main() {
	clientConf := client.Client{Host: "localhost:9080"}

	// b, err := clientConf.JudgeName("zjw1")
	// if err != nil {
	// 	fmt.Println("judegeName Error")
	// 	return
	// }

	// if b == true {
	// 	fmt.Println("user zjw1 exist")
	// }

	// b, err = clientConf.CreateNS("zjw1", "namespace1")
	// if err != nil {
	// 	fmt.Println("create ns wrong with error", err)
	// 	return
	// }
	// if b == true {
	// 	fmt.Println("create ns OK")
	// }

	res, err := clientConf.DeleteUser("test")
	if err != nil {
		fmt.Println("Get namespace Error")
		return
	}
	fmt.Println(res)
}
