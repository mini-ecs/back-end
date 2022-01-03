package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mini-ecs/back-end/internal/service"
	"github.com/mini-ecs/back-end/pkg/common/response"
	"net/http"
)

// GetImageList godoc
// @Summary      获取镜像列表
// @Description  Unimplemented
// @Tags         image management
// @Accept       json
// @Produce      json
// @Param        username  query     string        true  "用户名"
// @Param        passwd    query     string        true  "密码"
// @Response     400,200   {object}  response.Msg  ""
// @Router       /image [get]
func GetImageList(c *gin.Context) {
	logger.Infof("GetImageList")
	images := service.ImageManagement.GetImageList()
	c.JSON(http.StatusOK, response.SuccessMsg(images))
}

// GetSpecificImage godoc
// @Summary      获取镜像具体信息
// @Description  Unimplemented
// @Tags         image management
// @Accept       json
// @Produce      json
// @Param        username  query     string        true  "用户名"
// @Param        passwd    query     string        true  "密码"
// @Response     400,200   {object}  response.Msg  ""
// @Router       /image/:uuid [get]
func GetSpecificImage(c *gin.Context) {
	logger.Infof("GetSpecificImage")
	service.ImageManagement.GetSpecificImage()
	c.JSON(http.StatusOK, response.SuccessMsg("Unimplemented"))
}

// CreateImage godoc
// @Summary      创建（上传）镜像
// @Description  Unimplemented
// @Tags         image management
// @Accept       json
// @Produce      json
// @Param        username  query     string        true  "用户名"
// @Param        passwd    query     string        true  "密码"
// @Response     400,200   {object}  response.Msg  ""
// @Router       /image [post]
func CreateImage(c *gin.Context) {
	logger.Infof("CreateImage")
	service.ImageManagement.CreateImage()
	c.JSON(http.StatusOK, response.SuccessMsg("Unimplemented"))
}

// ModifyImage godoc
// @Summary      修改镜像条目信息
// @Description  Unimplemented
// @Tags         image management
// @Accept       json
// @Produce      json
// @Param        username  query     string        true  "用户名"
// @Param        passwd    query     string        true  "密码"
// @Response     400,200   {object}  response.Msg  ""
// @Router       /image/:uuid [put]
func ModifyImage(c *gin.Context) {
	logger.Infof("ModifyImage")
	service.ImageManagement.ModifyImage()
	c.JSON(http.StatusOK, response.SuccessMsg("Unimplemented"))
}

// DeleteImage godoc
// @Summary      删除镜像
// @Description  Unimplemented
// @Tags         image management
// @Accept       json
// @Produce      json
// @Param        username  query     string        true  "用户名"
// @Param        passwd    query     string        true  "密码"
// @Response     400,200   {object}  response.Msg  ""
// @Router       /image/:uuid [delete]
func DeleteImage(c *gin.Context) {
	logger.Infof("DeleteImage")
	service.ImageManagement.DeleteImage()
	c.JSON(http.StatusOK, response.SuccessMsg("Unimplemented"))
}
