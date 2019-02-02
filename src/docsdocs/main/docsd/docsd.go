package main

import (
	"docsdocs/log"
	"docsdocs/server"
)

func main() {
	log.Settings("json", "stdout", "debug")
	serv, err := server.NewServer("tcp", "localhost:3000")
	if err != nil {
		panic(err)
	}
	quit := serv.Run()
	defer func() { quit <- struct{}{} }()
}
