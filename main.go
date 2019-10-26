package main

import (
	"github.com/gorilla/mux"
	"os"
	"fmt"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	//router.Use(  ) 인증 미들웨어 등록

	port := os.Getenv("PORT")
	if "" == port {
		port = "8000"
	}

	err := http.ListenAndServe( ":" + port, router )
	if err != nil {
		fmt.Print ( err )
	}
}