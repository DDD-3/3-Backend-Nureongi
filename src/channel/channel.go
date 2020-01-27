package channel

import (
	"github.com/jinzhu/gorm"
)

type Drama struct {
	gorm.Model
	name string

	description string
}
