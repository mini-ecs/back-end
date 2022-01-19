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
		Name:           opt.InstanceName,
		CreatorID:      creator.ID,
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

	////------------------------Image copy operations---------------------
	//vm.ImageFileLocation = uuid.New().String()
	//// 拷贝镜像
	//err := image_manager.LocalMachineImpl.Copy(
	//	fmt.Sprintf("%v/%v", config.GetConfig().ImageStorage.FilePath, vm.ImageFileLocation),
	//	vm.SourceCourse.Image.Location,
	//)
	//if err != nil {
	//	return
	//}
	//
	////-----------------------------------libvirt operations--------------------
	//ip := net.ParseIP(config.GetConfig().NodeInfo.Ip)
	//l, err := virtlib.New(ip, strconv.Itoa(int(config.GetConfig().NodeInfo.Port)))
	//if err != nil {
	//	panic("generate env error: " + err.Error())
	//}
	//err = l.Connect()
	//if err != nil {
	//	panic(err)
	//}
	//defer l.DisConnect()
	//
	//d := virtlib.DefaultCreateDomainOpt
	//d.Uuid = uuid.New().String()
	//d.Name = vm.Name
	//d.Devices.Disk[1].Source.File = vm.ImageFileLocation
	//fmt.Printf("%+v\n", d)
	//err = l.CreateDomain(d)
	//if err != nil {
	//	panic(err)
	//}
	//vm.IP, err = l.GetDomainIPAddress(vm.Name)
	//if err != nil {
	//	panic(err)
	//}
	//------------------------------------------------------
}
func (v *vmManager) ModifyVM() {

}
func (v *vmManager) DeleteVM(id uint) error {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("DeleteVM, vm id: %v", id)
	vm := model.VM{}
	vm.ID = id
	res := db.First(&vm)
	if db.Error != nil {
		return db.Error
	}

	//// libvirt operations ----------
	//ip := net.ParseIP(config.GetConfig().NodeInfo.Ip)
	//l, err := virtlib.New(ip, strconv.Itoa(int(config.GetConfig().NodeInfo.Port)))
	//if err != nil {
	//	panic("generate env error: " + err.Error())
	//}
	//err = l.Connect()
	//if err != nil {
	//	panic(err)
	//}
	//defer l.DisConnect()
	//
	//// 销毁domain
	//err = l.DestroyDomain(vm.Name)
	//if err != nil {
	//	return err
	//}
	////删除镜像
	//err = image_manager.LocalMachineImpl.Delete(vm.ImageFileLocation)
	//if err != nil {
	//	return err
	//}
	//// ------------------

	res = db.Unscoped().Delete(&vm)
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
