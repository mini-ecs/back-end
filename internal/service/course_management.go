package service

import (
	"github.com/mini-ecs/back-end/internal/dao/pool"
	"github.com/mini-ecs/back-end/internal/model"
	"github.com/mini-ecs/back-end/pkg/log"
)

var CourseManager = new(courseManager)

type courseManager struct {
}

func (c *courseManager) GetCourseList() []model.Course {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("GetCourseList")
	var courses []model.Course
	res := db.Find(&courses)
	if res.Error != nil {
		log.GetGlobalLogger().Error(res.Error)
	}
	return courses
}

func (c *courseManager) GetMachineConfig() []model.MachineConfig {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("GetMachineConfig")
	var configs []model.MachineConfig
	res := db.Find(&configs)
	if res.Error != nil {
		log.GetGlobalLogger().Error(res.Error)
	}
	return configs
}
func (c *courseManager) GetSpecificCourse() {

}
func (c *courseManager) GetCourseLisCreateCourse() {

}
func (c *courseManager) ModifyCourse() {

}
func (c *courseManager) DeleteCourse() {

}
func (c *courseManager) CreateCourse(opt model.CreateCourseOpt) error {
	db := pool.GetDB()
	log.GetGlobalLogger().Infof("CreateCourse")

	image := model.ImageOrSnapshot{}
	res := db.First(&image, "name = ?", opt.ImageName)
	if res.Error != nil {
		log.GetGlobalLogger().Error(res.Error)
		return res.Error
	}

	machineConfig := model.MachineConfig{}
	res = db.First(&machineConfig, "id = ?", opt.ConfigNumber)
	if res.Error != nil {
		log.GetGlobalLogger().Error(res.Error)
		return res.Error
	}

	creator := model.User{}
	res = db.First(&creator, "uuid = ?", opt.Creator)
	if res.Error != nil {
		log.GetGlobalLogger().Error(res.Error)
		return res.Error
	}

	status := model.Status{Status: "unstart"}
	res = db.First(&status)
	if res.Error != nil {
		log.GetGlobalLogger().Error(res.Error)
		return res.Error
	}

	course := model.Course{
		CourseName:    opt.CourseName,
		MachineConfig: machineConfig,
		Image:         image,
		Teacher:       creator,
		Status:        status,
	}

	res = db.Create(&course)
	if res.Error != nil {
		log.GetGlobalLogger().Error(res.Error)
		return res.Error
	}
	return nil
}
