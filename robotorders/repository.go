package robotorders

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(robotorders Robotorder) (Robotorder, error)
	FindAllUserid(userid int) ([]Robotorder, error)
	FindAllKodeinvoice(kodeinvoice string) ([]Robotorder, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(robotorders Robotorder) (Robotorder, error) {
	err := r.db.Create(&robotorders).Error
	if err != nil {
		return robotorders, err
	}
	return robotorders, nil
}

func (r *repository) FindAllUserid(userid int) ([]Robotorder, error) {
	var robotorders []Robotorder

	err := r.db.Where("userid=?", userid).Find(&robotorders).Error
	if err != nil {
		return robotorders, err
	}

	return robotorders, nil
}

func (r *repository) FindAllKodeinvoice(kodeinvoice string) ([]Robotorder, error) {
	var robotorders []Robotorder

	err := r.db.Where("kodeinvoice=?", kodeinvoice).Find(&robotorders).Error
	if err != nil {
		return robotorders, err
	}

	return robotorders, nil
}
