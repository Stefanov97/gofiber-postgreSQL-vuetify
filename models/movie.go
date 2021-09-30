package models

// Movie represents a Movie schema
type Movie struct {
	ID   uint   `gorm:"primaryKey"`
	Title string `json:"title" gorm:"unique"`
	Genre  string `json:"genre"`
	Country string `json:"country"`
	Resume string `json:"resume"`
	DirectorID uint `json:"director"`
}