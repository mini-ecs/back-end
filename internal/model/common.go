package model

import "gorm.io/gorm"

type status struct {
	gorm.Model
	Status string `json:"status" gorm:"not null"`
}
