package handler

import (
	"net/http"
	"sewaMobil/helper"
	"sewaMobil/transaction"
	"sewaMobil/user"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transaction.Service
}

func NewTransactionHandler(transactionService transaction.Service) *transactionHandler {
	return &transactionHandler{transactionService}
}

// create
func (handler *transactionHandler) AddTransaction(c *gin.Context) {
	var input transaction.AddTransactionInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorFormat := helper.FormatValidationError(err)
		errorMsg := gin.H{
			"errors": errorFormat,
		}

		response := helper.APIResponse("Register Transaction Failed", http.StatusUnprocessableEntity, "error", errorMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return //stop function
	}

	// current user
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	// handler service
	dataTransaction, err := handler.transactionService.Create(input)
	if err != nil {
		errorFormat := helper.FormatValidationError(err)
		errorMsg := gin.H{
			"errors": errorFormat,
		}

		// response
		response := helper.APIResponse("Register Transaction Failed", http.StatusBadRequest, "error", errorMsg)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Register Transaction Success", http.StatusOK, "success", dataTransaction)
	c.JSON(http.StatusOK, response)
}

// find all
func (handler *transactionHandler) GetAllTransaction(c *gin.Context) {
	dataTransaction, err := handler.transactionService.FindAllTransaction()
	if err != nil {
		response := helper.APIResponse("Get All Transaction Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIListResponse("Get All Transaction Success", http.StatusOK, "success", dataTransaction, len(dataTransaction))
	c.JSON(http.StatusOK, response)
}

func (handler *transactionHandler) GetTransactionById(c *gin.Context) {
	var input transaction.GetTransactionIDInput

	err := c.ShouldBindJSON(input)
	if err != nil {
		errorFormat := helper.FormatValidationError(err)
		errorMsg := gin.H{
			"errors": errorFormat,
		}

		response := helper.APIResponse("Get Transaction By ID Failed", http.StatusUnprocessableEntity, "error", errorMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactionById, err := handler.transactionService.GetTransactionById(input.ID)
	if err != nil {
		errorFormat := helper.FormatValidationError(err)
		errorMsg := gin.H{
			"errors": errorFormat,
		}

		response := helper.APIResponse("Get Transaction By ID Failed", http.StatusBadRequest, "error", errorMsg)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get Transaction By ID Success", http.StatusOK, "success", transactionById)
	c.JSON(http.StatusOK, response)
}
