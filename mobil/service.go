package mobil

type Service interface {
	Create(input AddMobilInput) (Mobil, error)
	FindAllMobil() ([]Mobil, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// create
func (s *service) Create(input AddMobilInput) (Mobil, error) {
	mobil := Mobil{}
	mobil.Name = input.Name
	mobil.Type = input.Type
	mobil.Plat = input.Plat
	mobil.RentalPrice = input.RentalPrice
	mobil.PublicationYear = input.PublicationYear

	newMobil, err := s.repository.Create(mobil)
	if err != nil {
		return newMobil, err
	}

	return newMobil, nil
}

// find all
func (s *service) FindAllMobil() ([]Mobil, error) {
	mobil, err := s.repository.FindAllMobil()
	if err != nil {
		return mobil, err
	}

	return mobil, nil
}
