package transaction

import "gorm.io/gorm"

type Repository interface {
	Create(transaction Transaction) (Transaction, error)
	FindAllTransaction() ([]Transaction, error)
	GetById(ID int) (Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// create administrator
func (r *repository) Create(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

// find all admin
func (r *repository) FindAllTransaction() ([]Transaction, error) {
	var transaction []Transaction

	err := r.db.Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) GetById(ID int) (Transaction, error) {
	var transaction Transaction

	err := r.db.Where("id = ?", ID).First(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
