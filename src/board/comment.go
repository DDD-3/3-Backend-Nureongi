package board

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model

	commentId int32

	description string
}
