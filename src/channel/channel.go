package channel

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Channel struct {
	gorm.Model
	channelId   int32  `gorm:"primary_key;auto_increment" json:"id"`
	name        string `gorm:"size:255;not null;unique" json:"name"`
	description string `json:"description"`

	createdAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	updatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
