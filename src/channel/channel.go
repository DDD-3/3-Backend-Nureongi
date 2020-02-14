package channel

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Channel struct {
	gorm.Model
	Id          int64  `gorm:"primary_key;auto_increment" json:"id"`
	Name        string `gorm:"size:255" json:"name"`
	Description string `gorm:"text" json:"description"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
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
