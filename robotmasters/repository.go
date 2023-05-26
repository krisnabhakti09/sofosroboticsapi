package robotmasters

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Recomendation() ([]Robotmaster, error)
	FindAllRobotic() ([]Robotmaster, error)
	FindAllAutomation() ([]Robotmaster, error)
	Byid(id int) (Robotmaster, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Recomendation() ([]Robotmaster, error) {
	var robotmasters []Robotmaster

	err := r.db.Find(&robotmasters).Error
	if err != nil {
		return robotmasters, err
	}

	return robotmasters, nil
}

func (r *repository) FindAllRobotic() ([]Robotmaster, error) {
	var robotmasters []Robotmaster

	err := r.db.Where("typemaster=?", "robotics").Find(&robotmasters).Error
	if err != nil {
		return robotmasters, err
	}

	return robotmasters, nil
}

func (r *repository) FindAllAutomation() ([]Robotmaster, error) {
	var robotics []Robotmaster

	err := r.db.Where("typemaster=?", "automation").Find(&robotics).Error
	if err != nil {
		return robotics, err
	}

	return robotics, nil
}

func (r *repository) Byid(id int) (Robotmaster, error) {
	var robotmasters Robotmaster

	err := r.db.Where("id=?", id).Find(&robotmasters).Error
	if err != nil {
		return robotmasters, err
	}
	fmt.Println(robotmasters)
	return robotmasters, nil
}
