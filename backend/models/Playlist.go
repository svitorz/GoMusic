package models

type Playlist struct {
	ID     uint    `json:"id"`                                  // Standard field for the primary key
	Name   string  `json:"name"`                                // Name of the playlist
	Musics []Music `json:"musics" gorm:"foreignKey:PlaylistID"` // A slice of Music structs representing the songs in the playlist
}
