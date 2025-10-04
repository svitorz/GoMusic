package models

type Music struct {
	ID       uint
	Title    string
	AlbumID  uint
	Album    Album `gorm:"foreignKey:AlbumID"`
	Duration int
}
