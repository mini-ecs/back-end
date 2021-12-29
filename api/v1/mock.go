package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mini-ecs/back-end/pkg/common/response"
	"net/http"
)

func Welcome(c *gin.Context) {
	logger.Infof("Welcome page")
	c.JSON(http.StatusOK, response.SuccessMsg("Welcome to mini-ecs, the service is ok"))
}
