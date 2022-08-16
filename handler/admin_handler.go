package handler

import (
	"net/http"
	"sewaMobil/administrator"
	"sewaMobil/helper"

	"github.com/gin-gonic/gin"
)

type adminHandler struct {
	adminService administrator.Service
}

func NewAdminHandler(adminService administrator.Service) *adminHandler {
	return &adminHandler{adminService}
}

// register
func (handler *adminHandler) AddAdmin(c *gin.Context) {
	var input administrator.AddAdminInput

	// bind json
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorFormat := helper.FormatValidationError(err)
		errorMsg := gin.H{
			"errors": errorFormat,
		}

		response := helper.APIResponse("Register Admin Failed", http.StatusUnprocessableEntity, "error", errorMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return //stop function
	}

	// handler service
	adminData, err := handler.adminService.AddAdmin(input)
	if err != nil {
		errorFormat := helper.FormatValidationError(err)
		errorMsg := gin.H{
			"errors": errorFormat,
		}

		// response
		response := helper.APIResponse("Register Admin Failed", http.StatusBadRequest, "error", errorMsg)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Register Admin Success", http.StatusOK, "success", adminData)
	c.JSON(http.StatusOK, response)
}

func (handler *adminHandler) GetAllAdmin(c *gin.Context) {
	admin, err := handler.adminService.ListAdmin()
	if err != nil {
		response := helper.APIResponse("Get All Admin Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIListResponse("Get All Admin Success", http.StatusOK, "success", admin, len(admin))
	c.JSON(http.StatusOK, response)
}
