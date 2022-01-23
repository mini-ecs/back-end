package service

import (
	"github.com/google/uuid"
	"github.com/mini-ecs/back-end/internal/dao/pool"
	"github.com/mini-ecs/back-end/internal/model"
	"github.com/mini-ecs/back-end/me-errors"
	"github.com/mini-ecs/back-end/pkg/log"
)

var UserService = new(userService)

type userService struct {
}

func (u *userService) Login(user *model.User) bool {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("User %v try to login", user)

	queryUser := &model.User{}
	db.First(&queryUser, "username = ?", user.Username)
	log.GetGlobalLogger().Infof("query user %v ...", queryUser)

	user.Uuid = queryUser.Uuid
	return queryUser.Password == user.Password
}

func (u *userService) Register(user *model.User) error {
	db := pool.GetDB()
	var userCount int64
	db.Model(user).Where("username", user.Username).Count(&userCount)
	if userCount > 0 {
		return me_errors.New("user already exists")
	}
	user.Uuid = uuid.New().String()
	userType := model.UserType{Type: "student"}
	db.First(&userType, "type = ?", "student")
	user.UserTypeID = userType.ID
	db.Create(&user)
	return nil
}
func (u *userService) CurrentUser(uuid string) model.User {
	db := pool.GetDB()

	queryUser := &model.User{}
	db.First(&queryUser, "uuid = ?", uuid)
	db.Find(&queryUser.UserType, "ID = ?", queryUser.UserTypeID)
	return *queryUser
}
