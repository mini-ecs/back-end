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

// GetVMList godoc
// @Summary      获取实例列表
// @Description  Unimplemented
// @Tags         virtual machine management
// @Accept       json
// @Produce      json
// @Param        username  query     string        true  "用户名"
// @Param        passwd    query     string        true  "密码"
// @Response     400,200   {object}  response.Msg  ""
// @Router       /vm [get]
func GetVMList(c *gin.Context) {
	logger.Infof("GetVMList")
	userUUID, err := c.Cookie("uuid")
	if err != nil {
		panic(err)
	}
	vms := service.VMManager.GetVMList(userUUID)
	c.JSON(http.StatusOK, response.SuccessMsg(vms))
}

// GetSpecificVM godoc
// @Summary      获取实例信息
// @Description  Unimplemented
// @Tags         virtual machine management
// @Accept       json
// @Produce      json
// @Param        username  query     string        true  "用户名"
// @Param        passwd    query     string        true  "密码"
// @Response     400,200   {object}  response.Msg  ""
// @Router       /vm/:uuid [get]
func GetSpecificVM(c *gin.Context) {
	logger.Infof("GetSpecificVM")
	service.VMManager.GetSpecificVM()
	c.JSON(http.StatusOK, response.SuccessMsg("Unimplemented"))
}

// CreateVM godoc
// @Summary      创建实例
// @Description  Unimplemented
// @Tags         virtual machine management
// @Accept       json
// @Produce      json
// @Param        username  query     string        true  "用户名"
// @Param        passwd    query     string        true  "密码"
// @Response     400,200   {object}  response.Msg  ""
// @Router       /vm [post]
func CreateVM(c *gin.Context) {
	logger.Infof("CreateVM")
	opt := model.CreateVMOpt{}
	err := c.ShouldBindJSON(&opt)
	if err != nil {
		logger.Error(err)
	}
	logger.Info(opt)
	cookie, err := c.Cookie("uuid")
	if err != nil {
		panic(err)
	}
	opt.Creator = cookie
	err = service.VMManager.CreateVM(opt)
	if err != nil {
		c.JSON(http.StatusOK, response.FailCodeMsg(error_msg.ErrorDBOperation, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(""))
}

// ModifyVM godoc
// @Summary      Unimplemented
// @Description  Unimplemented
// @Tags         virtual machine management
// @Accept       json
// @Produce      json
// @Param        username  query     string        true  "用户名"
// @Param        passwd    query     string        true  "密码"
// @Response     400,200   {object}  response.Msg  ""
// @Router       /vm/:uuid [put]
func ModifyVM(c *gin.Context) {
	logger.Infof("ModifyVM")
	service.VMManager.ModifyVM()
	c.JSON(http.StatusOK, response.SuccessMsg("Unimplemented"))
}

// DeleteVM godoc
// @Summary      删除实例
// @Description  Unimplemented
// @Tags         virtual machine management
// @Accept       json
// @Produce      json
// @Param        username  query     string        true  "用户名"
// @Param        passwd    query     string        true  "密码"
// @Response     400,200   {object}  response.Msg  ""
// @Router       /vm/:uuid [delete]
func DeleteVM(c *gin.Context) {
	logger.Infof("DeleteVM")
	idStr := c.Param("uuid")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Errorf("parse string to int error: %v", err)
		return
	}
	userID, err := c.Cookie("uuid")
	if err != nil {
		panic(err)
	}
	err = service.VMManager.DeleteVM(uint(id), userID)
	if err != nil {
		c.JSON(http.StatusOK, response.FailCodeMsg(error_msg.ErrorDBOperation, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg("ok"))
}

// MakeSnapshotWithVM godoc
// @Summary      根据实例创建快照
// @Description  Unimplemented
// @Tags         virtual machine management
// @Accept       json
// @Produce      json
// @Param        username  query     string        true  "用户名"
// @Param        passwd    query     string        true  "密码"
// @Response     400,200   {object}  response.Msg  ""
// @Router       /vm/snapshot [post]
func MakeSnapshotWithVM(c *gin.Context) {
	logger.Infof("MakeSnapshotWithVM")
	//service.VMManager.MakeSnapshotWithVM()
	c.JSON(http.StatusOK, response.SuccessMsg("Unimplemented"))
}

// MakeImageWithVM godoc
// @Summary      根据实例创建镜像
// @Description  Unimplemented
// @Tags         virtual machine management
// @Accept       json
// @Produce      json
// @Param        username  query     string        true  "用户名"
// @Param        passwd    query     string        true  "密码"
// @Response     400,200   {object}  response.Msg  ""
// @Router       /vm/image [post]
func MakeImageWithVM(c *gin.Context) {
	logger.Infof("MakeImageWithVM")
	idStr := c.Param("uuid")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Errorf("parse string to int error: %v", err)
		return
	}
	cookie, err := c.Cookie("uuid")
	if err != nil {
		panic(err)
	}
	err = service.VMManager.MakeImageWithVM(uint(id), "test", cookie)
	if err != nil {
		c.JSON(http.StatusOK, response.FailCodeMsg(error_msg.ErrorInternal, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg("ok"))
}

// ResetVMWithSnapshot godoc
// @Summary      将实例恢复到某个快照
// @Description  Unimplemented
// @Tags         virtual machine management
// @Accept       json
// @Produce      json
// @Param        username  query     string        true  "用户名"
// @Param        passwd    query     string        true  "密码"
// @Response     400,200   {object}  response.Msg  ""
// @Router       /vm/snapshot [patch]
func ResetVMWithSnapshot(c *gin.Context) {
	logger.Infof("ResetVMWithSnapshot")
	//service.VMManager.MakeSnapshotWithVM()
	c.JSON(http.StatusOK, response.SuccessMsg("Unimplemented"))
}

func ShutDownVM(c *gin.Context) {
	logger.Infof("ShutDownVM")
	idStr := c.Param("uuid")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Errorf("parse string to int error: %v", err)
		return
	}
	userID, err := c.Cookie("uuid")
	if err != nil {
		logger.Errorln(err)
		return
	}
	err = service.VMManager.ShutdownVM(uint(id), userID)
	if err != nil {
		c.JSON(http.StatusOK, response.FailCodeMsg(error_msg.ErrorInternal, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg("ok"))
}

func RebootVM(c *gin.Context) {
	logger.Infof("ShutDownVM")
	idStr := c.Param("uuid")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Errorf("parse string to int error: %v", err)
		return
	}
	userID, err := c.Cookie("uuid")
	if err != nil {
		logger.Errorln(err)
		return
	}
	err = service.VMManager.RebootVM(uint(id), userID)
	if err != nil {
		c.JSON(http.StatusOK, response.FailCodeMsg(error_msg.ErrorInternal, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg("ok"))
}

func StartVM(c *gin.Context) {
	logger.Infof("ShutDownVM")
	idStr := c.Param("uuid")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Errorf("parse string to int error: %v", err)
		return
	}
	userID, err := c.Cookie("uuid")
	if err != nil {
		logger.Errorln(err)
		return
	}
	err = service.VMManager.StartVM(uint(id), userID)
	if err != nil {
		c.JSON(http.StatusOK, response.FailCodeMsg(error_msg.ErrorInternal, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg("ok"))
}
