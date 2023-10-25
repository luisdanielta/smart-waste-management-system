package models

import "gorm.io/gorm"

type Data_t struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type Container_t struct {
	gorm.Model
	Id           int    `gorm:"primaryKey"`
	ControllerId int    `gorm:"not null"`
	Name         string `gorm:"type:varchar(100);not null" json:"name"`
	Location     string `gorm:"type:varchar(254);not null" json:"location"`
	Usecase      string `gorm:"type:varchar(100);not null" json:"usecase"` // School, park, etc.
	Size         string `gorm:"type:varchar(100);not null" json:"size"`    // Small, medium, large
	Data         Data_t `gorm:"type:json;not null" json:"data"`

	// Controller ControllerC `gorm:"foreignKey:ControllerId"`
}
