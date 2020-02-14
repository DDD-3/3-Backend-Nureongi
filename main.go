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
	//s.initializeServiceServer(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

}
