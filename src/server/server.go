package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	///////////////////////////////////////////////
	"github.com/HodongMan/nureongi-server/src/channel"
	"github.com/HodongMan/nureongi-server/src/posting"
)

type ServiceServer struct {
	DB     *gorm.DB
	router *mux.Router

	isValid bool
}

func (server *ServiceServer) InitializeServiceServer(DatabaseDriver, DatabaseUser, DatabasePassword, DatabasePort, DatabaseHost, DatabaseName string) {

	var err error

	if DatabaseDriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DatabaseUser, DatabasePassword, DatabaseHost, DatabasePort, DatabaseName)
		server.DB, err = gorm.Open(DatabaseDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", DatabaseDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", DatabaseDriver)
		}
	}

	server.DB.Debug().AutoMigrate(&channel.Channel{})
	server.DB.Debug().AutoMigrate(&posting.Posting{})

	server.router = mux.NewRouter()
	server.setRoutes()

	server.isValid = true
}

func (server *ServiceServer) RunServer(port string) {

	if false == server.isValid {

	}

	err := http.ListenAndServe(":"+port, server.router)
	if err != nil {
		fmt.Print(err)
	}
}
