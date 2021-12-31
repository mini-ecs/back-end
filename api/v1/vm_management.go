package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mini-ecs/back-end/internal/service"
	"github.com/mini-ecs/back-end/pkg/common/response"
	"net/http"
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
	service.VMManager.GetVMList()
	c.JSON(http.StatusOK, response.SuccessMsg("Unimplemented"))
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
	service.VMManager.CreateVM()
	c.JSON(http.StatusOK, response.SuccessMsg("Unimplemented"))
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
	service.VMManager.DeleteVM()
	c.JSON(http.StatusOK, response.SuccessMsg("Unimplemented"))
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
	service.VMManager.MakeSnapshotWithVM()
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
	service.VMManager.MakeImageWithVM()
	c.JSON(http.StatusOK, response.SuccessMsg("Unimplemented"))
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
	service.VMManager.MakeSnapshotWithVM()
	c.JSON(http.StatusOK, response.SuccessMsg("Unimplemented"))
}
