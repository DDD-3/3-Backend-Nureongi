package character

import (
	"github.com/jinzhu/gorm"
)

type Character struct {
	gorm.Model
	characterId int32
	name        string
	gender      byte

	description string
}
