package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mini-ecs/back-end/internal/model"
	"github.com/mini-ecs/back-end/internal/service"
	"github.com/mini-ecs/back-end/pkg/common/response"
	"github.com/mini-ecs/back-end/pkg/log"
	"net/http"
)

var logger = log.GetGlobalLogger()

// RegisterUser godoc
// @Summary      用户注册
// @Description  用户注册
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        username  query     string        true  "用户名"
// @Param        passwd    query     string        true  "密码"
// @Response     400,200   {object}  response.Msg  ""
// @Router       /user/register [post]
func RegisterUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}
	logger.Infof("User %v try to register", user)

	err = service.UserService.Register(&user)
	if err != nil {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessMsg(user))
}

// Login godoc
// @Summary      用户登录
// @Description  用户登录
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        username  query     string        true  "用户名"
// @Param        passwd    query     string        true  "密码"
// @Response     400,200   {object}  response.Msg  ""
// @Router       /user/login [post]
func Login(c *gin.Context) {
	var user model.User
	// c.BindJSON(&user)
	err := c.ShouldBindJSON(&user)
	if err != nil {
		return
	}
	logger.Infof("User %v try to login", user)

	if service.UserService.Login(&user) {
		c.JSON(http.StatusOK, response.SuccessMsg(user))
		return
	}

	c.JSON(http.StatusOK, response.FailMsg("Login failed"))
}

// ModifyUser godoc
// @Summary      用户修改个人信息
// @Description  用户修改个人信息
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        username  query     string        true  "用户名"
// @Param        passwd    query     string        true  "密码"
// @Response     400,200   {object}  response.Msg  ""
// @Router       /user/modify [post]
func ModifyUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		return
	}
	logger.Infof("User %v try to modify", user)
	// todo
	//if service.UserService.Login(&user) {
	//	c.JSON(http.StatusOK, response.SuccessMsg(user))
	//	return
	//}

	c.JSON(http.StatusOK, response.FailMsg("Login failed"))
}
