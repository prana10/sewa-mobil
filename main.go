package main

import (
	"fmt"
	"log"
	"net/http"
	"sewaMobil/administrator"
	"sewaMobil/auth"
	"sewaMobil/handler"
	"sewaMobil/helper"
	"sewaMobil/mobil"
	"sewaMobil/transaction"
	"sewaMobil/user"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "gorm.io/driver/sqlite"
)

func main() {
	fmt.Println("Hello, World!")

	// set config databases;

	// db develop mysql
	dsn := "prana:1024@tcp(127.0.0.1:3306)/sample?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// db production
	// dsn := "prana:1024@tcp(127.0.0.1:3306)/sewaMobil?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// db sqlite
	// db, err := gorm.Open(sqlite.Open("sewaMobil.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// set repository
	userRepository := user.NewRepo(db)
	adminRepository := administrator.NewRepository(db)
	mobilRepository := mobil.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewService()
	adminService := administrator.NewService(adminRepository)
	mobilService := mobil.NewService(mobilRepository)
	transactionService := transaction.NewService(transactionRepository)

	userHandler := handler.NewUserHandler(userService, authService)
	adminHandler := handler.NewAdminHandler(adminService)
	mobilHandler := handler.NewMobilHandler(mobilService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	// routes config
	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/users/login", userHandler.LoginUser)
	api.GET("/users/", userHandler.GetAllUsers)
	api.GET("/users/me", authMiddleware(authService, userService), userHandler.FetchUser)
	api.POST("/users/email_check", userHandler.CheckEmailAvailability)

	api.POST("/administrator/add", adminHandler.AddAdmin)
	api.GET("/administrator", adminHandler.GetAllAdmin)

	// mobil
	api.GET("/mobil", mobilHandler.ListMobil)
	api.POST("/mobil/add", mobilHandler.AddMobil)
	// api.GET("/mobil/:id", userHandler.LoginUser)
	// api.PUT("/mobil/update/:id", userHandler.LoginUser)
	// api.DELETE("/mobil/delete/:id", userHandler.LoginUser)

	// Transaction
	api.GET("/transaction", transactionHandler.GetAllTransaction)
	api.POST("/transaction/add", authMiddleware(authService, userService), transactionHandler.AddTransaction)
	api.GET("/transaction/:id", authMiddleware(authService, userService), transactionHandler.GetTransactionById)
	// api.PUT("/transaction/:id", userHandler.LoginUser)
	// api.DELETE("/transaction/:id", userHandler.LoginUser)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.ServiceUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
