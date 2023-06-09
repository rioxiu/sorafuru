package handlers

import (
	"net/http"
	"sorafuru/helpers"
	"sorafuru/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandlers(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	//menangkap atau mengambil dari user
	//map input dari user ke struct InputRegister
	//struct diatas akan di passing sebagai parameter service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnsupportedMediaType, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if true {
		response := helpers.APIResponse("Account has registered", http.StatusBadRequest, "success", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//token

	formatter := user.FormatUser(newUser, "tokenkuy")

	response := helpers.APIResponse("Account has registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	//user menginput email dan password
	// handler menangkap input
	// mapping dari input user ke input struct
	// input struct passing di service
	// service mencari dengan bantuan  di repository user
	// mencocokan password yang di input dengan hash password

}
