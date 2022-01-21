package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mini-ecs/back-end/internal/dao/pool"
	"github.com/mini-ecs/back-end/internal/image_manager"
	"github.com/mini-ecs/back-end/internal/model"
	"github.com/mini-ecs/back-end/internal/virtlib"
	"github.com/mini-ecs/back-end/pkg/config"
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
	l := virtlib.GetConnectedLib()
	defer l.DisConnect()
	// todo: 可以加个缓存来减少查询次数
	for i := range vms {
		db.Find(&vms[i].Creator, "ID = ?", vms[i].CreatorID)
		db.Find(&vms[i].SourceCourse, "ID = ?", vms[i].SourceCourseID)
		db.Find(&vms[i].SourceCourse.MachineConfig, "ID = ?", vms[i].SourceCourse.MachineConfigID)
		// virtual machine is still preparing
		if vms[i].IP == "" {
			var err error
			vms[i].IP, err = l.GetDomainIPAddress(vms[i].Name)
			if err != nil {
				log.GetGlobalLogger().Infof("get ip address error: %v", err)
				continue
			}
			// virtual machine is ok, so it has ip address, set the status to running
			if vms[i].IP != "" {
				status := model.Status{}
				db.First(&status, "status = ?", "running")
				vms[i].StatusID = status.ID
				db.Model(&vms[i]).Update("status_id", status.ID)
			}
		}
		// default status is pending
		db.Find(&vms[i].Status, "ID = ?", vms[i].StatusID)
		//log.GetGlobalLogger().Infof("vm state: %+v", vms[i])
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
	log.GetGlobalLogger().Infof("course info: %+v", course)
	if res.Error != nil {
		return res.Error
	}
	res = db.First(&course.Image)
	if res.Error != nil {
		return res.Error
	}
	creator := model.User{Uuid: opt.Creator}
	res = db.First(&creator)
	if res.Error != nil {
		return res.Error
	}
	status := model.Status{}
	res = db.First(&status, "status = ?", "pending")
	log.GetGlobalLogger().Errorf("get status: %+v", status)
	if res.Error != nil {
		return res.Error
	}
	vm := model.VM{
		Name:           opt.InstanceName,
		CreatorID:      creator.ID,
		SourceCourseID: course.ID,
		SourceCourse:   course,
		StatusID:       status.ID,
	}

	fakeCreateVM(&vm)
	res = db.Create(&vm)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func fakeCreateVM(vm *model.VM) {
	//vm.Port = "999"
	//vm.LibvirtXML = "/user/bin"

	//------------------------Image copy operations---------------------
	vm.ImageFileName = uuid.New().String()
	// 拷贝镜像
	err := image_manager.LocalMachineImpl.Copy(
		fmt.Sprintf("%v/%v", config.GetConfig().ImageStorage.FilePath, vm.ImageFileName),
		vm.SourceCourse.Image.Location,
	)
	if err != nil {
		panic(err)
	}

	//-----------------------------------libvirt operations--------------------
	l := virtlib.GetConnectedLib()
	defer l.DisConnect()
	d := virtlib.DefaultCreateDomainOpt
	d.Uuid = uuid.New().String()
	d.Name = vm.Name
	d.Devices.Disk[1].Source.File = fmt.Sprintf("%v/%v", config.GetConfig().ImageStorage.FilePath, vm.ImageFileName)
	fmt.Printf("%+v\n", d)
	err = l.CreateDomain(d)
	if err != nil {
		panic(err)
	}
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
	if res.Error != nil {
		return db.Error
	}

	//// libvirt operations ----------
	l := virtlib.GetConnectedLib()
	defer l.DisConnect()

	// 销毁domain
	err := l.DestroyDomain(vm.Name)
	if err != nil {
		return err
	}
	//删除镜像
	err = image_manager.LocalMachineImpl.Delete(fmt.Sprintf("%v/%v", config.GetConfig().ImageStorage.FilePath, vm.ImageFileName))
	if err != nil {
		return err
	}

	res = db.Unscoped().Delete(&vm)
	if res.Error != nil {
		log.GetGlobalLogger().Error(res.Error)
		return res.Error
	}
	return nil
}
func (v *vmManager) MakeSnapshotWithVM(id uint) error {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("DeleteVM, vm id: %v", id)
	vm := model.VM{}
	vm.ID = id
	res := db.First(&vm)
	if res.Error != nil {
		return db.Error
	}
	l := virtlib.GetConnectedLib()
	defer l.DisConnect()
	snapshots, _ := l.ListSnapshots(vm.Name)
	cnt := len(snapshots)
	opt := virtlib.DomainSnapshot{
		Name: fmt.Sprintf("%v-snap%v", vm.Name, cnt+1),
	}
	return l.CreateSnapshot(vm.Name, opt)
}

// MakeImageWithVM todo: 将该函数改进为，删除所有的快照
func (v *vmManager) MakeImageWithVM(id uint) error {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("DeleteVM, vm id: %v", id)
	vm := model.VM{}
	vm.ID = id
	res := db.First(&vm)
	if res.Error != nil {
		return db.Error
	}
	// make sure the vm has stopped
	imagePath := fmt.Sprintf("%v/%v", config.GetConfig().ImageStorage.FilePath, vm.ImageFileName)
	name := fmt.Sprintf("%v's image", vm.ImageFileName)
	newPath := fmt.Sprintf("%v/%v", config.GetConfig().ImageStorage.FilePath, name)
	return image_manager.LocalMachineImpl.Copy(newPath, imagePath)
}
func (v *vmManager) ResetVMWithSnapshot() {

}
