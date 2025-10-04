package models

import "time"

type Album struct {
	ID        uint
	Title     string
	CreatorID uint
	Creator   User `gorm:"foreignKey:CreatorID"`
	CreatedAt time.Time
}
