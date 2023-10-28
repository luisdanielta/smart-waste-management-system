package models

import "gorm.io/gorm"

type Container struct {
	gorm.Model
	ControllerId string `gorm:"type:varchar(100);not null" json:"controller_id"`
	Name         string `gorm:"type:varchar(100);not null" json:"name"`
	Location     string `gorm:"type:varchar(254);not null" json:"location"`
	Usecase      string `gorm:"type:varchar(100);not null" json:"usecase"` // School, park, etc.
	Size         string `gorm:"type:varchar(100);not null" json:"size"`    // Small, medium, large
}
