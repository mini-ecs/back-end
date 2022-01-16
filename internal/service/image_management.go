package service

import (
	"github.com/mini-ecs/back-end/internal/dao/pool"
	"github.com/mini-ecs/back-end/internal/model"
	"github.com/mini-ecs/back-end/pkg/log"
)

var ImageManagement = new(imageManagement)

type imageManagement struct {
}

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

func (i *imageManagement) GetSpecificImage() {

}

func (i *imageManagement) CreateImage() {

}

func (i *imageManagement) ModifyImage() {

}

func (i *imageManagement) DeleteImage(id uint) error {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("GetMachineConfig, course id: %v", id)
	image := model.ImageOrSnapshot{}
	image.ID = id
	res := db.Unscoped().Delete(&image)
	if res.Error != nil {
		log.GetGlobalLogger().Error(res.Error)
		return res.Error
	}
	return nil
}
