package main

import (
	"fmt"
	"log"
	"os"
	"sorafuru/auth"
	models "sorafuru/db"
	"sorafuru/handlers"
	"sorafuru/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	//config database from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbConnectstring := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

	dbConnection, err := gorm.Open(mysql.Open(dbConnectstring), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	models := []interface{}{
		&models.UserDB{},
	}

	err = dbConnection.AutoMigrate(models...)
	if err != nil {
		log.Fatal("Error migrating database:", err)

	}

	//user service, auth services and handlers //
	userRepository := user.NewRepository(dbConnection)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handlers.NewUserHandlers(userService, authService)

	//router (yang diambil dari handlers )
	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/register", userHandler.RegisterUser)
	api.GET("/login", userHandler.LoginUser)
	api.POST("/check-email", userHandler.CheckingEmail)
	api.POST("/avatars", userHandler.AvatarHandlers)

	router.Run()

	// userRepository.Save(user)
	// fmt.Println(userHandler)
}

/*
	membuat middleware
	1. membuat header authorization : dengan bearer token
	2. header authorization sebagai tempat token
	3. validasi token yang didapatkan
	4. jika token valid, maka akan di cek user id nya yang diambil dari database dan service utama
*/
