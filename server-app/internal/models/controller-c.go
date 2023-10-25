package models

import "gorm.io/gorm"

type Location_t struct {
	City    string  `json:"city"`
	Country string  `json:"country"`
	Address string  `json:"address"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
}

type Status_t struct {
	Battery int  `json:"battery"`
	Signal  int  `json:"signal"`
	Online  bool `json:"online"`
}

type ControllerC struct {
	gorm.Model
	Id       int         `gorm:"primaryKey"`
	Name     string      `gorm:"type:varchar(100);not null" json:"name"`
	Version  string      `gorm:"type:varchar(100);not null" json:"version"`
	Ip       string      `gorm:"type:varchar(100);not null" json:"ip"`
	Status   Status_t    `gorm:"type:json;not null" json:"status"`
	Location Location_t  `gorm:"type:json;not null" json:"location"`
	Devices  []Container `gorm:"foreignKey:ControllerId"`
}
