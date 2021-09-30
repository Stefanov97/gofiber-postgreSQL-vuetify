package models

// Director represents a Director schema
type Director struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name"`
	Age  string `json:"age"`
}
