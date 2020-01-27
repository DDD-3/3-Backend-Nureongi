package user

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	userId    uint32    `gorm:"primary_key;auto_increment" json:"id"`
	name      string    `gorm:"size:255;not null;unique" json:"nickname"`
	email     string    `gorm:"size:100;not null;unique" json:"email"`
	password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func makeHash(password string) ([]byte, error) {

	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
