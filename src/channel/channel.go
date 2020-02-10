package channel

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Channel struct {
	channelId   int32  `gorm:"primary_key;auto_increment" json:"id"`
	name        string `gorm:"size:255;not null;unique" json:"name"`
	description string `json:"description"`

	createdAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	updatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (channel *Channel) getChannels(db *gorm.DB) (*[]Channel, error) {
	var errors error

	channels := []Channel{}
	errors = db.Debug().Model(&Channel{}).Limit(100).Find(&channels).Error

	if errors != nil {
		return &[]Channel{}, errors
	}

	return &channels, errors
}
