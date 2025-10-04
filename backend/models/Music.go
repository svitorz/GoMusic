package models

type Music struct {
	ID       uint   `json:"id"`                            // Standard field for the primary key
	Title    string `json:"title"`                         // A regular string field
	Album    Album  `json:"album" gorm:"foreignKey:Album"` // An Album struct representing the album of the music
	Duration int    `json:"duration"`                      // Duration in seconds
}
