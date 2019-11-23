package models

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type AuthToken struct {
	userId uint
	jwt.StandardClaims
}

type UserAccount struct {
	gorm.Model
	email    string `json:"email"`
	password string `json:"password"`
	name     string `json:"password"`
	token    string `json:"token";sql:"-"`
}

/*
	여기서 부터는 DB Procedure 정의 입니다.
*/

func (account *UserAccount) validateNotDBXXX() /*map[string] interface{},*/ error { // DB 검증 없는 버전

	if false == strings.Contains(account.email, "@") {
		//return utils.makeHttpMessage( false, "Email address is required" ), false
		return errors.New("eErrorInvalidEmail")
	}

	if len(account.password) < 6 {
		return errors.New("eErrorInvalidPasswordLength")
	}

	if len(account.name) < 2 {
		return errors.New("eErrorInvalidNameLength")
	}

	return nil
}

func (account *UserAccount) validate() /*map[string] interface{},*/ error {

	err := account.validateNotDBXXX()
	if nil != err {
		return err
	}

	tempAccount := &UserAccount{}

	err = getDB().Table("accounts").Where("email = ? AND name = ?", account.email, account.name).First(tempAccount).Error
	if nil == err && err == gorm.ErrRecordNotFound {
		return errors.New("eErrorInvalidEmailAndUserName")
	}

	if "" == tempAccount.email {
		return errors.New("eErrorInvalidEmailAndUserName")
	}

	return nil
}

func (account *UserAccount) createAccount() error /*map[string] interface{}*/ {

	// data validate
	if errors := account.validateNotDBXXX(); nil != errors {
		return errors
	}

	// data setting
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.password), bcrypt.DefaultCost)
	account.password = string(hashedPassword)

	// Transaction Begin
	tx := getDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	tempAccount := &UserAccount{}

	if err := getDB().Table("accounts").Where("email = ? AND name = ?", account.email, account.name).First(tempAccount).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := getDB().Create(account).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
	// end of transaction

	if account.ID <= 0 {
		return errors.New("DBCreateError")
	}

	autoToken := &AuthToken{userId: account.ID}

	return nil
}
