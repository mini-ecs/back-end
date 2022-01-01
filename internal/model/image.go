package model

import "gorm.io/gorm"

type ImageOrSnapshot struct {
	gorm.Model
	Type         string `json:"type" gorm:"not null"`
	Location     string `json:"location" gorm:"not null"`
	GenerateType int32  `json:"generateType" gorm:"not null"`

	CreatorID uint `json:"-"`
	Creator   User `json:"creator" gorm:"foreignKey:CreatorID;not null"`
}
