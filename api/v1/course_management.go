package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mini-ecs/back-end/internal/model"
	"github.com/mini-ecs/back-end/internal/service"
	"github.com/mini-ecs/back-end/pkg/common/error_msg"
	"github.com/mini-ecs/back-end/pkg/common/response"
	"net/http"
	"strconv"
)

// GetCourseList godoc
// @Summary      获取课程列表
// @Description  Unimplemented
// @Tags         course management
// @Accept       json
// @Produce      json
// @Param        username  query     string        true  "用户名"
// @Param        passwd    query     string        true  "密码"
// @Response     400,200   {object}  response.Msg  ""
// @Router       /course [get]
func GetCourseList(c *gin.Context) {
	logger.Infof("GetCourseList")
	courses := service.CourseManager.GetCourseList()
	//logger.Errorf("%+v", courses)
	c.JSON(http.StatusOK, response.SuccessMsg(courses))
}

// GetMachineConfig godoc
// @Summary      获取虚拟机配置列表
// @Description  Unimplemented
// @Tags         course management
// @Accept       json
// @Produce      json
// @Param        username  query     string        true  "用户名"
// @Param        passwd    query     string        true  "密码"
// @Response     400,200   {object}  response.Msg  ""
// @Router       /course/configs [get]
func GetMachineConfig(c *gin.Context) {
	logger.Infof("GetMachineConfig")
	configs := service.CourseManager.GetMachineConfig()
	c.JSON(http.StatusOK, response.SuccessMsg(configs))
}

// GetSpecificCourse godoc
// @Summary      获取课程信息
// @Description  Unimplemented
// @Tags         course management
// @Accept       json
// @Produce      json
// @Param        username  query     string        true  "用户名"
// @Param        passwd    query     string        true  "密码"
// @Response     400,200   {object}  response.Msg  ""
// @Router       /course/:uuid [get]
func GetSpecificCourse(c *gin.Context) {
	logger.Infof("GetSpecificCourse")
	idStr := c.Param("uuid")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Errorf("parse string to int error: %v", err)
		return
	}
	course := service.CourseManager.GetSpecificCourse(id)
	c.JSON(http.StatusOK, response.SuccessMsg(course))
}

// CreateCourse godoc
// @Summary      创建课程
// @Description  Unimplemented
// @Tags         course management
// @Accept       json
// @Produce      json
// @Param        username  query     string        true  "用户名"
// @Param        passwd    query     string        true  "密码"
// @Response     400,200   {object}  response.Msg  ""
// @Router       /course [post]
func CreateCourse(c *gin.Context) {
	logger.Infof("CreateCourse")
	var opt model.CreateCourseOpt
	//json := make(map[string]interface{})

	err := c.ShouldBindJSON(&opt)
	if err != nil {
		logger.Error(err)
	}
	logger.Infof("%+v", opt)
	cookie, err := c.Cookie("uuid")
	if err != nil {
		panic(err)
	}
	opt.Creator = cookie
	err = service.CourseManager.CreateCourse(opt)
	logger.Error("out, ", err)
	if err != nil {
		c.JSON(http.StatusOK, response.FailCodeMsg(error_msg.ErrorDBOperation, "fail"))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg("ok"))
}

// ModifyCourse godoc
// @Summary      修改课程
// @Description  Unimplemented
// @Tags         course management
// @Accept       json
// @Produce      json
// @Param        username  query     string        true  "用户名"
// @Param        passwd    query     string        true  "密码"
// @Response     400,200   {object}  response.Msg  ""
// @Router       /course/:uuid [put]
func ModifyCourse(c *gin.Context) {
	logger.Infof("ModifyCourse")
	var opt model.CreateCourseOpt
	//json := make(map[string]interface{})

	err := c.ShouldBindJSON(&opt)
	if err != nil {
		logger.Error(err)
	}
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
	err = service.CourseManager.ModifyCourse(uint(id), userID, opt)
	if err != nil {
		c.JSON(http.StatusOK, response.FailCodeMsg(error_msg.ErrorDBOperation, err.Error()))
	}
	c.JSON(http.StatusOK, response.SuccessMsg("ok"))
}

// DeleteCourse godoc
// @Summary      删除课程
// @Description  Unimplemented
// @Tags         course management
// @Accept       json
// @Produce      json
// @Param        username  query     string        true  "用户名"
// @Param        passwd    query     string        true  "密码"
// @Response     400,200   {object}  response.Msg  ""
// @Router       /course/:uuid [delete]
func DeleteCourse(c *gin.Context) {
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
	err = service.CourseManager.DeleteCourse(uint(id), userID)
	if err != nil {
		c.JSON(http.StatusOK, response.FailCodeMsg(error_msg.ErrorDBOperation, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg("ok"))
}
