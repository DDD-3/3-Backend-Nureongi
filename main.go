package main

import (
	"os"

	"github.com/HodongMan/nureongi-server/src/server"
)

func main() {

	port := os.Getenv("PORT")
	if "" == port {
		port = "8000"
	}

	s := server.ServiceServer{}
	s.InitializeServiceServer("mysql",
		"root",
		"1234",
		"3306",
		"127.0.0.1",
		"ddd")

	s.RunServer(port)

}
