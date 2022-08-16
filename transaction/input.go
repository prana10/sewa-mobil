package transaction

import (
	"sewaMobil/user"
)

type AddTransactionInput struct {
	User            user.User
	MobilId         int    `json:"mobil_id"`
	AdministratorId int    `json:"administrator_id"`
	RentalDate      string `json:"rental_date"`
	ReturnDate      string `json:"return_date"`
	Penalty         int    `json:"penalty"`
}

type GetTransactionIDInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
