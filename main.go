package main

import (
	"os"

	"github.com/hodongman/github.com/HodongMan/nureongi-server/server"
)

func main() {

	port := os.Getenv("PORT")
	if "" == port {
		port = "8000"
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

}
