package model

import (
	"gorm.io/gorm"
)

type UserType struct {
	gorm.Model
	Type string `json:"type" gorm:"comment:'type'"`
}

type User struct {
	gorm.Model
	Uuid     string `json:"uuid" gorm:"type:varchar(150);not null;unique_index:idx_uuid;comment:'uuid'" example:"1234345"`
	Username string `json:"username" form:"username" binding:"required" gorm:"unique;not null; comment:'用户名'" example:"account name"`
	Password string `json:"password" form:"password" binding:"required" gorm:"type:varchar(150);not null; comment:'密码'" example:"password"`
	Nickname string `json:"nickname" gorm:"comment:'昵称'" example:"nickname"`
	Avatar   string `json:"avatar" gorm:"type:varchar(150);comment:'头像'" example:"bcdedit"`
	Email    string `json:"email" gorm:"type:varchar(80);column:email;comment:'邮箱'" example:"2123@qq.com"`

	UserTypeID uint     `json:"-"`
	UserType   UserType `json:"userType" gorm:"foreignKey:UserTypeID"`
}
