package model

import "gorm.io/gorm"

type VM struct {
	gorm.Model
	IP            string `json:"ip" gorm:"not null"`
	Port          string `json:"port" gorm:"not null"`
	LibvirtXML    string `json:"libvirtXML" gorm:"not null"`
	Name          string `json:"name" gorm:"unique;not null"`
	ImageFileName string `json:"imageFileLocation" gorm:"not null"`

	CreatorID      uint   `json:"-"`
	Creator        User   `json:"creator" gorm:"foreignKey:CreatorID"`
	SourceCourseID uint   `json:"-"`
	SourceCourse   Course `json:"sourceCourse" gorm:"foreignKey:SourceCourseID"`
	StatusID       uint   `json:"-"`
	Status         Status `json:"Status" gorm:"foreignKey:StatusID"`
}

type CreateVMOpt struct {
	CourseName   string `json:"courseName"`
	InstanceName string `json:"instanceName"`
	Creator      string
}
