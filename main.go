package main

import (
	"fmt"
	"log"
	"sorafuru/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:polo1028@tcp(127.0.0.1:3305)/sorafuru?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}

	userInput.Fullname = "Pavolia Reine"
	userInput.Email = "pavoliareine@hololive.id"
	userInput.Password = "merakpalingmegahwoi"
	userInput.Occupation = "Vtuber"

	userService.RegisterUser(userInput)

	// userRepository.Save(user)
	fmt.Println("insert data berhasil")
}
