package model

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	Status string `json:"Status" gorm:"not null"`
}
