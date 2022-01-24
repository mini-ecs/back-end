package service

import (
	"errors"
	"github.com/mini-ecs/back-end/internal/dao/pool"
	"github.com/mini-ecs/back-end/internal/image_manager"
	"github.com/mini-ecs/back-end/internal/model"
	"github.com/mini-ecs/back-end/pkg/log"
)

var ImageManagement = new(imageManagement)

type imageManagement struct {
}

// GetImageList 获取当前用户拥有的image
func (i *imageManagement) GetImageList() []model.ImageOrSnapshot {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("GetCourseList")
	var images []model.ImageOrSnapshot
	res := db.Find(&images)
	if res.Error != nil {
		log.GetGlobalLogger().Error(res.Error)
	}
	for i := range images {
		db.Find(&images[i].Creator, "ID = ?", images[i].CreatorID)
	}
	return images
}

// GetSpecificImage 根据名字获取镜像
func (i *imageManagement) GetSpecificImage(name string) model.ImageOrSnapshot {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("GetCourseList")
	var image model.ImageOrSnapshot
	res := db.Find(&image, "name = ?", name)
	if res.Error != nil {
		log.GetGlobalLogger().Error(res.Error)
	}
	return image
}

// CreateImageByCopy 创建镜像
func (i *imageManagement) CreateImageByCopy(old, new string) error {
	o := i.GetSpecificImage(old)
	n := i.GetSpecificImage(new)
	err := image_manager.LocalMachineImpl.Copy(o.Location, n.Location)
	if err != nil {
		return err
	}
	return nil
}

func (i *imageManagement) ModifyImage() {

}

func (i *imageManagement) DeleteImage(id uint, userID string) error {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("GetMachineConfig, course id: %v", id)
	image := model.ImageOrSnapshot{}
	image.ID = id
	res := db.Find(&image)
	if res.Error != nil {
		log.GetGlobalLogger().Errorln(res.Error)
	}
	res = db.Find(&image.Creator, "id = ?", image.CreatorID)
	if res.Error != nil {
		return db.Error
	}
	res = db.Find(&image.Creator.UserType, "id = ?", image.Creator.UserTypeID)
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
	if operator.UserType.Type != "admin" && image.Creator.Uuid != userID {
		return errors.New("unauthorized operation")
	}

	err := image_manager.LocalMachineImpl.Delete(image.Location)
	if err != nil {
		return err
	}

	res = db.Unscoped().Delete(&image)
	if res.Error != nil {
		log.GetGlobalLogger().Error(res.Error)
		return res.Error
	}

	return nil
}
