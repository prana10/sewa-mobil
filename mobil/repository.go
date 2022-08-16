package mobil

import "gorm.io/gorm"

type Repository interface {
	Create(mobil Mobil) (Mobil, error)
	FindAllMobil() ([]Mobil, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// create administrator
func (r *repository) Create(mobil Mobil) (Mobil, error) {
	err := r.db.Create(&mobil).Error

	if err != nil {
		return mobil, err
	}

	return mobil, nil
}

// find all admin
func (r *repository) FindAllMobil() ([]Mobil, error) {
	var mobil []Mobil

	err := r.db.Find(&mobil).Error
	if err != nil {
		return mobil, err
	}

	return mobil, nil
}
