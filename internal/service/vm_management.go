package service

import (
	"github.com/mini-ecs/back-end/internal/dao/pool"
	"github.com/mini-ecs/back-end/internal/model"
	"github.com/mini-ecs/back-end/pkg/log"
)

var VMManager = new(vmManager)

type vmManager struct {
}

func (v *vmManager) GetVMList() []model.VM {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("GetVMList")
	var vms []model.VM
	res := db.Find(&vms)
	if res.Error != nil {
		log.GetGlobalLogger().Error(res.Error)
	}
	// todo: 可以加个缓存来减少查询次数
	for i := range vms {
		db.Find(&vms[i].Creator, "ID = ?", vms[i].CreatorID)
		db.Find(&vms[i].Status, "ID = ?", vms[i].StatusID)
		db.Find(&vms[i].SourceCourse, "ID = ?", vms[i].SourceCourseID)
		db.Find(&vms[i].SourceCourse.MachineConfig, "ID = ?", vms[i].SourceCourse.MachineConfigID)
	}
	return vms
}
func (v *vmManager) GetSpecificVM() {

}
func (v *vmManager) CreateVM(opt model.CreateVMOpt) error {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("CreateVM")
	course := model.Course{CourseName: opt.CourseName}
	res := db.First(&course)
	if res.Error != nil {
		return res.Error
	}
	creator := model.User{Uuid: opt.Creator}
	res = db.First(&creator)
	if res.Error != nil {
		return res.Error
	}
	vm := model.VM{
		Name: opt.InstanceName,
		//Creator:        creator,
		CreatorID: creator.ID,
		//SourceCourse:   course,
		SourceCourseID: course.ID,
	}

	fakeCreateVM(&vm)
	res = db.Create(&vm)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func fakeCreateVM(vm *model.VM) {
	vm.IP = "123.123.123.123"
	vm.Port = "999"
	vm.LibvirtXML = "/user/bin"
	vm.StatusID = 1
}
func (v *vmManager) ModifyVM() {

}
func (v *vmManager) DeleteVM(id uint) error {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("DeleteVM, vm id: %v", id)
	vm := model.VM{}
	vm.ID = id
	res := db.Unscoped().Delete(&vm)
	if res.Error != nil {
		log.GetGlobalLogger().Error(res.Error)
		return res.Error
	}
	return nil
}
func (v *vmManager) MakeSnapshotWithVM() {

}
func (v *vmManager) MakeImageWithVM() {

}
func (v *vmManager) ResetVMWithSnapshot() {

}
