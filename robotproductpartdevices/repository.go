package robotproductpartdevices

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAllIDRobotProductPartDevice(robotproductpartdevice int) (Robotproductpartdevice, error)
	FindAllIDRobotMasterAndRobotPartDevice(robotmasters int, robotpartdevice int) ([]Robotproductpartdevice, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllIDRobotMasterAndRobotPartDevice(robotmasters int, robotpartdevice int) ([]Robotproductpartdevice, error) {
	var robotmaster []Robotproductpartdevice

	err := r.db.Where("idrobotmaster=? and idrobotpartdevice=?", robotmasters, robotpartdevice).Find(&robotmaster).Error
	if err != nil {
		return robotmaster, err
	}

	return robotmaster, nil
}

func (r *repository) FindAllIDRobotProductPartDevice(robotproductpartdevice int) (Robotproductpartdevice, error) {
	var robotmaster Robotproductpartdevice

	err := r.db.Where("id=?", robotproductpartdevice).Find(&robotmaster).Error
	if err != nil {
		return robotmaster, err
	}

	return robotmaster, nil
}
