package board

import (
	"github.com/jinzhu/gorm"
)

type Board struct {
	gorm.Model

	boardId     int32
	name        string
	description string
}
