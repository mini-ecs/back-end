package v1

import (
	"github.com/gin-contrib/sessions"
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
// @Param        user     body      model.User    true  "user"
// @Response     400,200  {object}  response.Msg  ""
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
// @Param        user     body      model.User    true  "user"
// @Response     400,200  {object}  response.Msg  ""
// @Router       /user/login [post]
func Login(c *gin.Context) {
	var user model.User
	// c.BindJSON(&user)
	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}
	logger.Infof("User %v try to login", user)

	if service.UserService.Login(&user) {
		c.SetCookie("uuid", user.Uuid, 3600, "/", "localhost", false, true)

		session := sessions.Default(c)
		// 通过session.Get读取session值
		// session是键值对格式数据，因此需要通过key查询数据
		if session.Get(user.Uuid) != "online" {
			// 设置session数据
			session.Set(user.Uuid, "online")
			// 保存session数据
			err := session.Save()
			if err != nil {
				logger.Error(err)
			}
		}

		c.JSON(http.StatusOK, response.SuccessMsg(user))
		return
	}

	c.JSON(http.StatusOK, response.FailMsg("Login failed"))
}

// CurrentUser godoc
// @Summary      用户登录
// @Description  用户登录
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user     body      model.User    true  "user"
// @Response     400,200  {object}  response.Msg  ""
// @Router       /currentUser [get]
func CurrentUser(c *gin.Context) {
	cookie, err := c.Cookie("uuid")
	if err != nil {
		c.JSON(http.StatusOK, response.FailMsg("You should login first"))
		return
	}
	session := sessions.Default(c)
	if session.Get(cookie) != "online" {
		c.JSON(http.StatusOK, response.FailMsg("You should login first"))
		return
	}
	user := service.UserService.CurrentUser(cookie)
	user.Avatar = "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"
	c.JSON(http.StatusOK, response.SuccessMsg(user))
}

// ModifyUser godoc
// @Summary      用户修改个人信息
// @Description  用户修改个人信息
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user     body      model.User    true  "user"
// @Response     400,200  {object}  response.Msg  ""
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
