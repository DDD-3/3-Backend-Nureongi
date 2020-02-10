package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	///////////////////////////////////////////////
	"github.com/hodongman/github.com/HodongMan/nureongi-server/channel"
)

type ServiceServer struct {
	DB     *gorm.DB
	router *mux.Router

	isValid bool
}

func (server *ServiceServer) initializeServiceServer(DatabaseDriver, DatabaseUser, DatabasePassword, DatabasePort, DatabaseHost, DatabaseName string) {

	var errors error

	if DatabaseDriver == "mysql" {

	}

	server.DB.Debug().AutoMigrate(&channel.Channel{})

	server.router = mux.NewRouter()
}

func (server *ServiceServer) runServer(port string) {

	err := http.ListenAndServe(":"+port, server.router)
	if err != nil {
		fmt.Print(err)
	}
}
