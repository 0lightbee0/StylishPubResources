package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	f, err := os.OpenFile("./log/main.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
	if err = serveStart(); err != nil {
		log.Panic(err)
	}
}

func serveStart() error {
	addr, err := getAddr()
	if err != nil {
		return err
	}
	initHandle()
	return http.ListenAndServe(addr, nil)
}
