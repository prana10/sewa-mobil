package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindAll() ([]User, error)
	FindByID(ID int) (User, error)
}

type repository struct {
	db *gorm.DB
}

// instance of repository
func NewRepo(db *gorm.DB) *repository {
	return &repository{db}
}

// save user
func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *repository) FindByEmail(email string) (User, error) {
	var user User
	err := repo.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindAll() ([]User, error) {
	var users []User

	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) FindByID(ID int) (User, error) {
	var user User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
