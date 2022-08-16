package transaction

type Service interface {
	Create(input AddTransactionInput) (Transaction, error)
	FindAllTransaction() ([]Transaction, error)
	GetTransactionById(ID int) (Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// create
func (s *service) Create(input AddTransactionInput) (Transaction, error) {
	// convert date

	transaction := Transaction{}
	transaction.UserId = input.User.ID
	transaction.AdministratorId = input.AdministratorId
	transaction.MobilId = input.MobilId
	transaction.RentalDate = input.RentalDate
	transaction.ReturnDate = input.ReturnDate
	transaction.Penalty = input.Penalty

	newTransaction, err := s.repository.Create(transaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

// find all
func (s *service) FindAllTransaction() ([]Transaction, error) {
	mobil, err := s.repository.FindAllTransaction()
	if err != nil {
		return mobil, err
	}

	return mobil, nil
}

// get by id
func (s *service) GetTransactionById(ID int) (Transaction, error) {
	transaction, err := s.repository.GetById(ID)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
