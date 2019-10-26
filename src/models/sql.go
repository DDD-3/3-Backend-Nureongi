package models

import (
	"os"
	"fmt"
	"errors"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func initDatabase() ( error ) {

	environment := godotenv.Load()
	if ( environment == nil ) {
		return errors.New( "환경 설정 로드에 실패했습니다. 이후 종료 합니다." )
	}

	fmt.Print( environment )

	username 	:= os.Getenv( "db_user" )
	password 	:= os.Getenv( "db_pass" )
	dbName 		:= os.Getenv( "db_name" )
	dbHost 		:= os.Getenv( "db_host" )

	databaseURL := fmt.Sprintf( "host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password )
	fmt.Println( databaseURL )

	conn, err := gorm.Open( "mysql", databaseURL )
	if err != nil {
		fmt.Println( "DB 연결에 실패했습니다 이후에 프로세스 종료 합니다" )
		return err
	}

	db = conn
	//db.Debug().AutoMigrate( &Account{}, &Contact{} )

	return nil
}

func getDB() ( *gorm.DB ) {
	return db
}