package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mini-ecs/back-end/internal/model"
	"github.com/mini-ecs/back-end/internal/service"
	"github.com/mini-ecs/back-end/pkg/common/response"
	"net/http"
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
	logger.Errorf("%+v", courses)
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
	service.CourseManager.GetSpecificCourse()
	c.JSON(http.StatusOK, response.SuccessMsg("Unimplemented"))
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
		c.JSON(http.StatusOK, response.FailMsg("fail"))
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
	service.CourseManager.ModifyCourse()
	c.JSON(http.StatusOK, response.SuccessMsg("Unimplemented"))
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
	service.CourseManager.DeleteCourse()
	c.JSON(http.StatusOK, response.SuccessMsg("Unimplemented"))
}
