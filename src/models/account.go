package models

import (
	"strings"
	"errors"
	"os"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	utils "../utils"
)

type AuthToken struct {
	userId uint
	jwt.StandardClaims
}

type UserAccount struct {
	gorm.Model
	email string `json:"email"`
	password string `json:"password"`
	name string `json:"password"`
	token string `json:"token";sql:"-"`
}

/*
	여기서 부터는 DB Procedure 정의 입니다.
*/

func ( account *UserAccount ) validateNotDBXXX() ( /*map[string] interface{},*/ error ) { // DB 검증 없는 버전
	
	if false == strings.Contains( account.email, "@" ) {
		//return utils.makeHttpMessage( false, "Email address is required" ), false
		return errors.New( "eErrorInvalidEmail" )
	}

	if len( account.password ) < 6 {
		return errors.New( "eErrorInvalidPasswordLength" )
	}

	if len( account.name ) < 2 {
		return errors.New( "eErrorInvalidNameLength" )
	}

	return nil
}

func ( account *UserAccount ) validate() ( /*map[string] interface{},*/ error ) {
	
	err := account.validateNotDBXXX()
	if ( nil != err ) {
		return err
	}

	tempAccount := &UserAccount{}

	err = GetDB().Table( "accounts" ).Where( "email = ? AND name = ?" , account.email, account.name ).First( tempAccount ).Error
	if ( nil == err && err == gorm.ErrRecordNotFound ) {
		return errors.New( "eErrorInvalidEmailAndUserName" )
	}

	if ( "" == tempAccount.email ) {

	}

	return nil
}

func ( account *UserAccount ) createAccount() ( error, /*map[string] interface{}*/ ) {

	return nil
}