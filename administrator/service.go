package administrator

import "golang.org/x/crypto/bcrypt"

type Service interface {
	AddAdmin(input AddAdminInput) (Administrator, error)
	ListAdmin() ([]Administrator, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// register
func (s *service) AddAdmin(input AddAdminInput) (Administrator, error) {
	administrator := Administrator{}
	administrator.Name = input.Name
	administrator.Email = input.Email

	// hashing
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return administrator, err
	}

	administrator.Password = string(hashPassword)

	newAdmin, err := s.repository.Create(administrator)
	if err != nil {
		return administrator, err
	}

	return newAdmin, nil
}

// list Admin
func (s *service) ListAdmin() ([]Administrator, error) {
	admin, err := s.repository.FindAllAdmin()
	if err != nil {
		return admin, err
	}

	return admin, nil
}
