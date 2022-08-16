package transaction

import (
	"sewaMobil/user"
	"time"
)

type Transaction struct {
	ID              int
	UserId          int
	MobilId         int
	AdministratorId int
	RentalDate      string
	ReturnDate      string
	Penalty         int
	User            user.User
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
