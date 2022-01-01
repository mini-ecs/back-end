package model

import "gorm.io/gorm"

type VM struct {
	gorm.Model
	IP         string `json:"ip" gorm:"not null"`
	Port       string `json:"port" gorm:"not null"`
	LibvirtXML string `json:"libvirtXML" gorm:"not null"`

	CreatorID      uint   `json:"-"`
	Creator        User   `json:"creator" gorm:"foreignKey:CreatorID"`
	SourceCourseID uint   `json:"-"`
	SourceCourse   Course `json:"sourceCourse" gorm:"foreignKey:CreatorID"`
	StatusID       uint   `json:"-"`
	Status         status `json:"status" gorm:"foreignKey:CreatorID"`
}
