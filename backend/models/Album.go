package models

import "time"

type Album struct {
	ID        uint      `json:"id"`                                // Standard field for the primary key
	Title     string    `json:"title"`                             // A regular string field
	Creator   User      `json:"creator_id" gorm:"foreignKey:User"` // A User struct representing the creator of the album
	CreatedAt time.Time `json:"creator_id"`
}
