package models

import (
	"time"

	"gorm.io/gorm"
)

type ControllerC struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Ip       string `gorm:"type:varchar(100);not null" json:"ip"`
	Status   bool   `gorm:"type:bool;not null" json:"status"`
	Battery  int    `gorm:"type:int;not null" json:"battery"`
	Signal   int    `gorm:"type:int;not null" json:"signal"`
	Location string `gorm:"type:varchar(254);not null" json:"location"`

	/* config */
	CreatedAt time.Time `gorm:"autoCreateTime:false"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:false"`
	DeletedAt gorm.DeletedAt
}

func (c *ControllerC) SetStatus() bool {
	c.Status = true
	return true
}
