package handler

import (
	"net/http"
	"sewaMobil/helper"
	"sewaMobil/mobil"

	"github.com/gin-gonic/gin"
)

type mobilHandler struct {
	mobilService mobil.Service
}

func NewMobilHandler(mobilService mobil.Service) *mobilHandler {
	return &mobilHandler{mobilService}
}

func (handler *mobilHandler) AddMobil(c *gin.Context) {
	var input mobil.AddMobilInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorFormat := helper.FormatValidationError(err)
		errorMsg := gin.H{
			"errors": errorFormat,
		}
		response := helper.APIResponse("Failed Create Mobil", http.StatusUnprocessableEntity, "error", errorMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	mobil, err := handler.mobilService.Create(input)
	if err != nil {
		errorFormat := helper.FormatValidationError(err)
		errorMsg := gin.H{
			"errors": errorFormat,
		}
		response := helper.APIResponse("Failed Create Mobil", http.StatusBadRequest, "error", errorMsg)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success Create Mobil", http.StatusCreated, "data", mobil)
	c.JSON(http.StatusOK, response)
}

// list mobil
func (handler *mobilHandler) ListMobil(c *gin.Context) {
	mobil, err := handler.mobilService.FindAllMobil()
	if err != nil {
		response := helper.APIResponse("Failed Get All Mobil", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIListResponse("Success Get All Mobil", http.StatusOK, "data", mobil, len(mobil))
	c.JSON(http.StatusOK, response)
}
