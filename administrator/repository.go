package administrator

import "gorm.io/gorm"

type Repository interface {
	Create(administrator Administrator) (Administrator, error)
	FindAllAdmin() ([]Administrator, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// create administrator
func (r *repository) Create(administrator Administrator) (Administrator, error) {
	err := r.db.Create(&administrator).Error

	if err != nil {
		return administrator, err
	}

	return administrator, nil
}

// find all admin
func (r *repository) FindAllAdmin() ([]Administrator, error) {
	var admin []Administrator

	err := r.db.Find(&admin).Error
	if err != nil {
		return admin, err
	}

	return admin, nil
}
