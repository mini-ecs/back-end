package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mini-ecs/back-end/internal/service"
	"github.com/mini-ecs/back-end/pkg/common/error_msg"
	"github.com/mini-ecs/back-end/pkg/common/response"
	"github.com/mini-ecs/back-end/pkg/config"
	"net/http"
	"strconv"
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
	//service.ImageManagement.GetSpecificImage()
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
	form, err := c.MultipartForm()
	_ = err
	_ = form
	//path := "/home/fangaoyang/work/" + form.File["file"][0].Filename
	path := config.GetConfig().ImageStorage.FilePath + "/" + form.File["file"][0].Filename
	err = c.SaveUploadedFile(form.File["file"][0], path)

	userID, err := c.Cookie("uuid")
	if err != nil {
		logger.Errorln(err)
	}
	if err := service.ImageManagement.UploadImage(form.File["file"][0].Filename, path, userID); err != nil {
		c.JSON(http.StatusOK, response.FailCodeMsg(error_msg.UploadImageFailed, err.Error()))
		return

	}
	c.JSON(http.StatusOK, response.SuccessMsg("ok"))
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
	logger.Infof("DeleteCourse")
	idStr := c.Param("uuid")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Errorf("parse string to int error: %v", err)
		return
	}
	userID, err := c.Cookie("uuid")
	if err != nil {
		logger.Errorln(err)
	}
	err = service.ImageManagement.DeleteImage(uint(id), userID)
	if err != nil {
		c.JSON(http.StatusOK, response.FailCodeMsg(error_msg.ErrorDBOperation, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg("ok"))
}
