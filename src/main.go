package main

import (
	"fmt"
	"net/http"
)

func main() {
	addr, _ := getAddr()
	fmt.Println("[" + addr + "]Serve has running...")
	if err := initDirs(); err != nil {
		fmt.Println("配置路径出错: ", err)
	}
	initHandle()
	if err := serveStart(); err != nil {
		fmt.Println(err)
	}
}

func serveStart() error {
	addr, err := getAddr()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return http.ListenAndServe(addr, nil)
}
