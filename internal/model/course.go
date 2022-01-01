package model

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	CourseName     string `json:"courseName" gorm:"unique;index;not null"`
	BaseLibvirtXML string `json:"baseLibvirtXML"`

	TeacherID       uint            `json:"-"`
	Teacher         User            `json:"userType" gorm:"foreignKey:TeacherID"`
	StatusID        uint            `json:"-" `
	Status          status          `json:"status" gorm:"foreignKey:StatusID"`
	MachineConfigID uint            `json:"-" `
	MachineConfig   machineConfig   `json:"machineConfig" gorm:"foreignKey:MachineConfigID"`
	ImageID         uint            `json:"-" `
	Image           ImageOrSnapshot `json:"image" gorm:"foreignKey:ImageID"`
}

type machineConfig struct {
	gorm.Model
	CPU int32 `json:"cpu" gorm:"not null"`
	RAM int32 `json:"ram" gorm:"not null"`
}
