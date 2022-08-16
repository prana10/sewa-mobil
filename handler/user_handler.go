package handler

import (
	"net/http"
	"sewaMobil/auth"
	"sewaMobil/helper"
	"sewaMobil/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.ServiceUser
	authService auth.Service
}

func NewUserHandler(userService user.ServiceUser, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (handler *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		formatError := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": formatError,
		}
		response := helper.APIResponse("Register Account Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userData, err := handler.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// token
	token, err := handler.authService.GenerateToken(userData.ID)
	if err != nil {
		response := helper.APIResponse("Register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	format := user.FormatUserRegister(userData, token)
	response := helper.APIResponse("Register Account Success", http.StatusOK, "success", format)
	c.JSON(http.StatusOK, response)
}

func (handler *userHandler) LoginUser(c *gin.Context) {
	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorFormat := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": errorFormat,
		}

		response := helper.APIResponse("Login Account Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := handler.userService.LoginUser(input)

	if err != nil {
		errorMsg := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login Account Failed", http.StatusBadRequest, "error", errorMsg)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// token
	token, err := handler.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.APIResponse("Login Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	format := user.FormatUser(loggedinUser, token)
	response := helper.APIResponse("Login Account Success", http.StatusOK, "success", format)
	c.JSON(http.StatusOK, response)
}

func (handler *userHandler) GetAllUsers(c *gin.Context) {
	// service get all users
	users, err := handler.userService.GetAllUser()
	if err != nil {
		errorFormat := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": errorFormat,
		}
		response := helper.APIResponse("Get All User Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// format user
	response := helper.APIListResponse("Get All User Success", http.StatusOK, "success", users, len(users))
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	formatter := user.FormatUser(currentUser, "")
	response := helper.APIResponse("Successfuly fetch user data", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
