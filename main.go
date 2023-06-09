package main

import (
	"fmt"
	"log"
	"sorafuru/handlers"
	"sorafuru/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:riosiu@tcp(127.0.0.1:3306)/sorafuru?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handlers.NewUserHandlers(userService)

	userByEmail, err := userRepository.FindByEmail("rioxiu@icloud.com") //mencari email yang ada di database

	if err != nil {
		fmt.Println("error", err.Error())

	}

	if userByEmail.ID == 0 {
		fmt.Println("user not found")
	} else {

		fmt.Println(userByEmail.Name)
	}

	//router
	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)

	router.Run()

	// userRepository.Save(user)
	fmt.Println(userHandler)
}
