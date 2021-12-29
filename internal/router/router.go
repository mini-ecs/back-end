package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/mini-ecs/back-end/api/v1"
	"github.com/mini-ecs/back-end/pkg/common/response"
	"github.com/mini-ecs/back-end/pkg/log"

	"net/http"
)

var logger = log.GetGlobalLogger()

func NewRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	server := gin.Default()
	server.Use(Recovery)
	group := server.Group("api/v1")
	{
		group.GET("/welcome", v1.Welcome)
		group.POST("/user/login", v1.Login)
		group.POST("/user/register", v1.Register)

	}
	return server
}
func Recovery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logger.Errorf("gin catch error: %v", r)
			c.JSON(http.StatusOK, response.FailMsg("系统内部错误"))
		}
	}()
	c.Next()
}
