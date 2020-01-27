package posting

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Posting struct {
	gorm.Model
	postinglId  int32  `gorm:"primary_key;auto_increment" json:"id"`
	title       string `gorm:"size:255;not null;unique" json:"title"`
	description string `json:"description"`

	createdAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	updatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
