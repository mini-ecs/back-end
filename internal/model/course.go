package model

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	CourseName     string `json:"courseName" gorm:"unique;index;not null"`
	BaseLibvirtXML string `json:"baseLibvirtXML"`

	TeacherID       uint            `json:"-"`
	Teacher         User            `json:"teacher" gorm:"foreignKey:TeacherID"`
	StatusID        uint            `json:"-" `
	Status          Status          `json:"status" gorm:"foreignKey:StatusID"`
	MachineConfigID uint            `json:"-" `
	MachineConfig   MachineConfig   `json:"machineConfig" gorm:"foreignKey:MachineConfigID"`
	ImageID         uint            `json:"-" `
	Image           ImageOrSnapshot `json:"image" gorm:"foreignKey:ImageID"`
}

type MachineConfig struct {
	gorm.Model
	CPU int32 `json:"cpu" gorm:"not null"`
	RAM int32 `json:"ram" gorm:"not null"`
}

type CreateCourseOpt struct {
	CourseName   string `json:"courseName"`
	ImageName    string `json:"imageName"`
	Note         string `json:"note"`
	ConfigNumber uint   `json:"configs"`
	Creator      string
}
