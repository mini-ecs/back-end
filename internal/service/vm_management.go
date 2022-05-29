package service

import (
	"errors"
	"fmt"
	"github.com/digitalocean/go-libvirt"
	"github.com/google/uuid"
	"github.com/mini-ecs/back-end/internal/dao/pool"
	"github.com/mini-ecs/back-end/internal/image_manager"
	"github.com/mini-ecs/back-end/internal/model"
	"github.com/mini-ecs/back-end/internal/virtlib"
	"github.com/mini-ecs/back-end/pkg/config"
	"github.com/mini-ecs/back-end/pkg/log"
	"strconv"
    "time"
)

var VMManager = new(vmManager)

type vmManager struct {
}

func (v *vmManager) GetVMList(uuid string) []model.VM {
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
	// 如果用户是admin，则返回所有列表，否则只返回自己创建的
	user := model.User{}
	res = db.First(&user, "uuid = ?", uuid)
	if res.Error != nil {
		log.GetGlobalLogger().Errorln(res.Error)
	}
	res = db.First(&user.UserType, "id = ?", user.UserTypeID)
	if res.Error != nil {
		log.GetGlobalLogger().Errorln(res.Error)
	}
	if user.UserType.Type == "admin" {
		return vms
	}
	ret := make([]model.VM, 0)
	for _, v := range vms {
		if v.Creator.Uuid == uuid {
			ret = append(ret, v)
		}
	}
	return ret
}
func (v *vmManager) GetSpecificVM(uuid string) (model.VM, error) {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("GetVMList")
	var vms model.VM
	res := db.Find(&vms)
	if res.Error != nil {
		log.GetGlobalLogger().Error(res.Error)
	}
	l := virtlib.GetConnectedLib()
	defer l.DisConnect()

	db.Find(&vms.Creator, "ID = ?", vms.CreatorID)
	db.Find(&vms.SourceCourse, "ID = ?", vms.SourceCourseID)
	db.Find(&vms.SourceCourse.MachineConfig, "ID = ?", vms.SourceCourse.MachineConfigID)
	// virtual machine is still preparing
	if vms.IP == "" {
		var err error
		vms.IP, err = l.GetDomainIPAddress(vms.Name)
		if err != nil {
			log.GetGlobalLogger().Infof("get ip address error: %v", err)
			return model.VM{}, err
		}
		// virtual machine is ok, so it has ip address, set the status to running
		if vms.IP != "" {
			status := model.Status{}
			db.First(&status, "status = ?", "running")
			vms.StatusID = status.ID
			db.Model(&vms).Update("status_id", status.ID)
		}
	}
	// default status is pending
	db.Find(&vms.Status, "ID = ?", vms.StatusID)

	// 如果用户是admin，则返回所有列表，否则只返回自己创建的
	user := model.User{}
	res = db.First(&user, "uuid = ?", uuid)
	if res.Error != nil {
		log.GetGlobalLogger().Errorln(res.Error)
	}
	res = db.First(&user.UserType, "id = ?", user.UserTypeID)
	if res.Error != nil {
		log.GetGlobalLogger().Errorln(res.Error)
	}
	if user.UserType.Type == "admin" {
		return vms, nil
	}
	if vms.Creator.Uuid == uuid {
		return vms, nil
	}

	return model.VM{}, errors.New("no such virtual machine")
}
func (v *vmManager) GetVNCPort(id uint) (string, error) {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("GetVNCPort, vm id: %v", id)
	vm := model.VM{}
	vm.ID = id
	res := db.First(&vm)
	if res.Error != nil {
		return "", db.Error
	}

	l := virtlib.GetConnectedLib()
	vncStr := l.GetDomainVNCPort(vm.Name)
	vnc, err := strconv.Atoi(vncStr)
	port, err := virtlib.ProxyVNCToWebSocket(vnc)
	return strconv.Itoa(port), err
}
func (v *vmManager) GetMemUsage(id uint) (float64, error) {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("GetVNCPort, vm id: %v", id)
	vm := model.VM{}
	vm.ID = id
	res := db.First(&vm)
	if res.Error != nil {
		return 0.0, db.Error
	}
	l := virtlib.GetConnectedLib()
	return l.GetDomMemUsage(vm.Name)
}

func (v *vmManager) GetDiskUsage(id uint) (float64, error) {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("GetVNCPort, vm id: %v", id)
	vm := model.VM{}
	vm.ID = id
	res := db.First(&vm)
	if res.Error != nil {
		return 0.0, db.Error
	}
	l := virtlib.GetConnectedLib()
	return l.GetDomDiskUsage(vm.Name)
}

func (v *vmManager) GetCPUUsage(id uint) (float64, error) {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("GetVNCPort, vm id: %v", id)
	vm := model.VM{}
	vm.ID = id
	res := db.First(&vm)
	if res.Error != nil {
		return 0.0, db.Error
	}
	l := virtlib.GetConnectedLib()
	return l.GetDomCPUUsage(vm.Name, 1)
}
func (v *vmManager) CreateVM(opt model.CreateVMOpt) error {
    defer func(start time.Time) {
        log.GetGlobalLogger().Infof("创建虚拟机耗时：%v", time.Since(start).String())
    }(time.Now())
    db := pool.GetDB()
    log.GetGlobalLogger().Infof("CreateVM")
    course := model.Course{}
    res := db.First(&course, "course_name = ?", opt.CourseName)
    log.GetGlobalLogger().Infof("course info: %+v", course)
    if res.Error != nil {
        return res.Error
    }
    res = db.First(&course.Image, "id = ?", course.ImageID)
    if res.Error != nil {
        return res.Error
    }
    res = db.First(&course.MachineConfig, "id = ?", course.MachineConfigID)
    if res.Error != nil {
        return res.Error
    }
    creator := model.User{}
    res = db.First(&creator, "uuid = ?", opt.Creator)
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

    //------------------------Image copy operations---------------------
    // 注意，该文件是最基本的镜像文件，此后创建的快照都会在改文件后添加".Snapshot"样式的后缀
    // 用数据库记录该实例拥有的快照的地址，以便于后续操作管理(删除、恢复、合并快照)
    vm.BaseImageFileName = uuid.New().String()
    snapshotPath := getImageFilePath(vm.BaseImageFileName)
    // 拷贝镜像
    err := image_manager.LocalMachineImpl.Copy(
        snapshotPath,
        vm.SourceCourse.Image.Location,
	)
	if err != nil {
		panic(err)
	}
	log.GetGlobalLogger().Infof("Create VM, copy the source Image Finished, Image path: %v, source file: %v", snapshotPath, vm.SourceCourse.Image.Location)
	snapshot := model.Snapshot{
		VMName:           vm.Name,
		SnapshotName:     vm.BaseImageFileName,
		SnapshotLocation: snapshotPath,
    }
    res = db.Create(&snapshot)
    if res.Error != nil {
        panic(res.Error)
    }

    //-----------------------------------libvirt operations--------------------
    l := virtlib.GetConnectedLib()
    defer l.DisConnect()
    d := virtlib.DefaultCreateDomainOpt
    d.Memory.Text = fmt.Sprintf("%v", course.MachineConfig.RAM)
    d.Vcpu = fmt.Sprintf("%v", course.MachineConfig.CPU)
    d.Uuid = uuid.New().String()
    d.Name = vm.Name
    d.Devices.Disk[1].Source.File = snapshotPath
    //d.Devices.Disk = d.Devices.Disk[1:]

    fmt.Printf("%+v\n", d)
    err = l.CreateDomain(d)
    if err != nil {
        panic(err)
    }
    //------------------------------------------------------

	res = db.Create(&vm)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
func getImageFilePath(imageName string) string {
	return fmt.Sprintf("%v/%v", config.GetConfig().ImageStorage.FilePath, imageName)
}
func (v *vmManager) ModifyVM() {

}
func (v *vmManager) DeleteVM(id uint, userID string) error {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("DeleteVM, vm id: %v", id)
	vm := model.VM{}
	vm.ID = id
	res := db.First(&vm)
	if res.Error != nil {
		return db.Error
	}
	res = db.Find(&vm.Creator, "id = ?", vm.CreatorID)
	if res.Error != nil {
		return db.Error
	}
	res = db.Find(&vm.Creator.UserType, "id = ?", vm.Creator.UserTypeID)
	if res.Error != nil {
		return db.Error
	}
	operator := model.User{}
	res = db.Find(&operator, "uuid = ?", userID)
	if res.Error != nil {
		return db.Error
	}
	res = db.Find(&operator.UserType, "id = ?", operator.UserTypeID)
	if res.Error != nil {
		return db.Error
	}
	if operator.UserType.Type != "admin" && vm.Creator.Uuid != userID {
		return errors.New("unauthorized operation, " + vm.Creator.UserType.Type)
	}
	//// libvirt operations ----------
	l := virtlib.GetConnectedLib()
	defer l.DisConnect()

	// 销毁domain
	err := l.DestroyDomain(vm.Name)
	if err != nil {
		return err
	}
	err = l.UnDefineDomain(vm.Name)
	if err != nil {
		return err
	}
	//删除镜像
	var snapshots []model.Snapshot
	res = db.Find(&snapshots, "vm_name = ?", vm.Name)
	if res.Error != nil {
		return res.Error
	}
	for _, v := range snapshots {
		err = image_manager.LocalMachineImpl.Delete(v.SnapshotLocation)
		if err != nil {
			log.GetGlobalLogger().Errorln(err)
		}
	}
	if _, err = l.GetCurrentSnapshot(vm.Name); err == nil {
		// 获取所有快照
		allSnapshots, err := l.ListSnapshots(vm.Name)
		if err != nil {
			return err
		}
		for _, v := range allSnapshots {
			err = l.DeleteSnapshot(vm.Name, v.Name, libvirt.DomainSnapshotDeleteMetadataOnly)
			if err != nil {
				log.GetGlobalLogger().Errorln(err)
				continue
			}
		}
	}

	res = db.Unscoped().Delete(&vm)
	if res.Error != nil {
		log.GetGlobalLogger().Error(res.Error)
		return res.Error
	}
	return nil
}

func (v *vmManager) MakeSnapshotWithVM(id uint, userID string) error {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("DeleteVM, vm id: %v", id)
	vm := model.VM{}
	vm.ID = id
	res := db.First(&vm)
	if res.Error != nil {
		return db.Error
	}
	res = db.Find(&vm.Creator, "id = ?", vm.CreatorID)
	if res.Error != nil {
		return db.Error
	}
	res = db.Find(&vm.Creator.UserType, "id = ?", vm.Creator.UserTypeID)
	if res.Error != nil {
		return db.Error
	}
	operator := model.User{}
	res = db.Find(&operator, "uuid = ?", userID)
	if res.Error != nil {
		return db.Error
	}
	res = db.Find(&operator.UserType, "id = ?", operator.UserTypeID)
	if res.Error != nil {
		return db.Error
	}
	// 1. 学生不允许操作
	// 2. 如果是老师，则需要满足操作的是他自己的实例
	// 3. 其余的是admin，没限制
	if operator.UserType.Type != "admin" &&
		(operator.UserType.Type != "teacher" || vm.Creator.Uuid != userID) {
		return errors.New("unauthorized operation")
	}
	l := virtlib.GetConnectedLib()
	defer l.DisConnect()
	snapshots, _ := l.ListSnapshots(vm.Name)
	cnt := len(snapshots)
	opt := virtlib.DomainSnapshot{
		Name: fmt.Sprintf("snap-%v", cnt+1),
	}
	snapshot := model.Snapshot{
		VMName:           vm.Name,
		SnapshotName:     opt.Name,
		SnapshotLocation: getImageFilePath(vm.BaseImageFileName) + "." + opt.Name,
	}
	res = db.Create(&snapshot)
	if res.Error != nil {
		return db.Error
	}
	return l.CreateSnapshot(vm.Name, opt)
}

// MakeImageWithVM 合并所有快照，只保留最后合并的快照，其余快照将被删除，然后以传入的名字创建一个新镜像
func (v *vmManager) MakeImageWithVM(id uint, imageName, userUUid string) error {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("MakeImageWithVM, vm id: %v", id)
	vm := model.VM{}
	vm.ID = id
	res := db.First(&vm)
	if res.Error != nil {
		return db.Error
	}
	res = db.Find(&vm.Creator, "id = ?", vm.CreatorID)
	if res.Error != nil {
		return db.Error
	}
	res = db.Find(&vm.Creator.UserType, "id = ?", vm.Creator.UserTypeID)
	if res.Error != nil {
		return db.Error
	}
	// 1. 学生不允许操作
	// 2. 如果是老师，则需要满足操作的是他自己的实例
	// 3. 其余的是admin，没限制
	if vm.Creator.UserType.Type != "admin" &&
		(vm.Creator.UserType.Type != "teacher" || vm.Creator.Uuid != userUUid) {
		return errors.New("unauthorized operation")
	}

	l := virtlib.GetConnectedLib()
	defer l.DisConnect()
	// 确认是否有快照，没有快照，则说明已经只有一份镜像了
	// 获取当前快照
	cur, err := l.GetCurrentSnapshot(vm.Name)
	var curSnap model.Snapshot
	if err == nil {
		//将所有的快照都合并到最新的镜像文件中，此时老文件都没用了
		err = l.PullAllSnapshots(vm.Name)
		if err != nil {
			return err
		}
		// 获取所有快照
		snapshots, err := l.ListSnapshots(vm.Name)
		if err != nil {
			return err
		}
		// 删除除了当前快照之外的所有快照的元数据和真实文件
		for _, v := range snapshots {
			// 从db获取snapshot数据
			snapshot := model.Snapshot{}
			res = db.Find(&snapshot, "snapshot_name = ?", v.Name)
			if res.Error != nil {
				log.GetGlobalLogger().Errorf("Deleteing snapshot, get %v error: %v", v.Name, res.Error)
				continue
			}
			// 如果是当前的快照，则获取一下其地址
			if cur.Name == v.Name {
				curSnap = snapshot
				continue
			}
			// 删除文件
			err := image_manager.LocalMachineImpl.Delete(snapshot.SnapshotLocation)
			if err != nil {
				log.GetGlobalLogger().Errorf("Deleteing snapshot, get %v error: %v", v.Name, err)
				continue
			}
			// 清除数据库记录
			res = db.Delete(&snapshot)
			if res.Error != nil {
				log.GetGlobalLogger().Errorf("Deleteing snapshot, get %v error: %v", v.Name, res.Error)
				continue
			}
			// 删除快照元数据
			err = l.DeleteSnapshot(vm.Name, v.Name, libvirt.DomainSnapshotDeleteMetadataOnly)
			if err != nil {
				log.GetGlobalLogger().Errorln(err)
				continue
			}
		}
	} else {
		// 不存在快照，则把这个值赋为当前使用的镜像路径
		curSnap.SnapshotLocation = getImageFilePath(vm.BaseImageFileName)
	}
	log.GetGlobalLogger().Infof("开始拷贝镜像")
	// make sure the vm has stopped
	newPath := getImageFilePath(vm.Name)
	err = image_manager.LocalMachineImpl.Copy(newPath, curSnap.SnapshotLocation)
	if err != nil {
		return err
	}
	// 创建数据库的镜像记录
	creator := model.User{Uuid: userUUid}
	res = db.First(&creator)
	image := model.ImageOrSnapshot{
		Type:         "image",
		Location:     newPath,
		GenerateType: 1,
		Name:         fmt.Sprintf("%v[IMAGE]", vm.Name),
		Creator:      creator,
	}
	res = db.Create(&image)
	return res.Error
}
func (v *vmManager) ResetVMWithSnapshot() {

}
func (v *vmManager) ShutdownVM(id uint, userID string) error {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("MakeImageWithVM, vm id: %v", id)
	vm := model.VM{}
	vm.ID = id
	res := db.First(&vm)
	if res.Error != nil {
		return db.Error
	}
	res = db.Find(&vm.Creator, "id = ?", vm.CreatorID)
	if res.Error != nil {
		return db.Error
	}
	res = db.Find(&vm.Creator.UserType, "id = ?", vm.Creator.UserTypeID)
	if res.Error != nil {
		return db.Error
	}
	operator := model.User{}
	res = db.Find(&operator, "uuid = ?", userID)
	if res.Error != nil {
		return db.Error
	}
	res = db.Find(&operator.UserType, "id = ?", operator.UserTypeID)
	if res.Error != nil {
		return db.Error
	}
	if operator.UserType.Type != "admin" && vm.Creator.Uuid != userID {
		return errors.New("unauthorized operation")
	}

	l := virtlib.GetConnectedLib()
	err := l.ShutdownDomain(vm.Name)
	if err != nil {
		return err
	}
	state := model.Status{}
	db.Find(&state, "status = ?", "unstart")
	db.Model(&vm).Update("status_id", state.ID)
	return nil
}

func (v *vmManager) RebootVM(id uint, userID string) error {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("MakeImageWithVM, vm id: %v", id)
	vm := model.VM{}
	vm.ID = id
	res := db.First(&vm)
	if res.Error != nil {
		return db.Error
	}
	res = db.Find(&vm.Creator, "id = ?", vm.CreatorID)
	if res.Error != nil {
		return db.Error
	}
	res = db.Find(&vm.Creator.UserType, "id = ?", vm.Creator.UserTypeID)
	if res.Error != nil {
		return db.Error
	}
	operator := model.User{}
	res = db.Find(&operator, "uuid = ?", userID)
	if res.Error != nil {
		return db.Error
	}
	res = db.Find(&operator.UserType, "id = ?", operator.UserTypeID)
	if res.Error != nil {
		return db.Error
	}
	if operator.UserType.Type != "admin" && vm.Creator.Uuid != userID {
		return errors.New("unauthorized operation")
	}
	l := virtlib.GetConnectedLib()
	return l.RebootDomain(vm.Name)
}

func (v *vmManager) StartVM(id uint, userID string) error {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("MakeImageWithVM, vm id: %v", id)
	vm := model.VM{}
	vm.ID = id
	res := db.First(&vm)
	if res.Error != nil {
		return db.Error
	}
	res = db.Find(&vm.Creator, "id = ?", vm.CreatorID)
	if res.Error != nil {
		return db.Error
	}
	res = db.Find(&vm.Creator.UserType, "id = ?", vm.Creator.UserTypeID)
	if res.Error != nil {
		return db.Error
	}
	operator := model.User{}
	res = db.Find(&operator, "uuid = ?", userID)
	if res.Error != nil {
		return db.Error
	}
	res = db.Find(&operator.UserType, "id = ?", operator.UserTypeID)
	if res.Error != nil {
		return db.Error
	}
	if operator.UserType.Type != "admin" && vm.Creator.Uuid != userID {
		return errors.New("unauthorized operation")
	}
	l := virtlib.GetConnectedLib()
	err := l.StartDomain(vm.Name)
	if err != nil {
		return err
	}
	state := model.Status{}
	db.Find(&state, "status = ?", "running")
	db.Model(&vm).Update("status_id", state.ID)
	return nil
}
