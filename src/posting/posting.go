package posting

import (
	"time"
)

/*
type Posting struct {
	gorm.Model
	postinglId  int32  `gorm:"primary_key;auto_increment" json:"id"`
	title       string `gorm:"size:255;not null;unique" json:"title"`
	description string `json:"description"`

	createdAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	updatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
*/
type Posting struct {
	ID      uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title   string `gorm:"size:255;not null;unique" json:"title"`
	Content string `gorm:"size:255;not null;" json:"content"`
	//Author    User      `json:"author"`
	//AuthorID  uint32    `gorm:"not null" json:"author_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
