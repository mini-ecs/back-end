package model

import "gorm.io/gorm"

type ImageOrSnapshot struct {
	gorm.Model
	Type         string `json:"type" gorm:"not null"`
	Location     string `json:"location" gorm:"not null;unique"`
	GenerateType int32  `json:"generateType" gorm:"not null"`
	Name         string `json:"name" gorm:"not null"`

	CreatorID uint `json:"-"`
	Creator   User `json:"creator" gorm:"foreignKey:CreatorID;not null"`
}

type Snapshot struct {
	gorm.Model
	VMName           string
	SnapshotName     string
	SnapshotLocation string
}
