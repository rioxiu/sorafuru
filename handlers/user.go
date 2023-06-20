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

	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnsupportedMediaType, response)
		return
	}

	loggedInUser, err := h.userService.LoginUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helpers.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnsupportedMediaType, response)
		return
	}
	//token
	formatter := user.FormatUser(loggedInUser, "tokenkuy")
	response := helpers.APIResponse("Successfuly loggedin", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckingEmail(c *gin.Context) {
	/*
		ada input email dari user saat register
		input email di mapping ke struct input
		struct input di passing ke service
		service akan memanggil repository - email sudah ada atau belum
		repository - db
	*/

	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Check email was failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnsupportedMediaType, response)
		return
	}

	// CheckEmail, err := h.userService.CheckEmail(input)
	CheckEmail, err := h.userService.CheckEmail(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}

		response := helpers.APIResponse("Check email was failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnsupportedMediaType, response)
		return
	}

	data := gin.H{"is_available": CheckEmail}
	var metaMassage string

	if CheckEmail {
		metaMassage = "Email Available"
	} else {
		metaMassage = "Email has been registered"
	}

	response := helpers.APIResponse(metaMassage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) AvatarHandlers(c *gin.Context) {
	//input dari user
	//gambar akan di simpan ke folder "images"
	// services akan memanggil repository
	// menggunakan jwt sebagai validasi
	// repo ambil data user yang id nya di dapat dari jwt
	// repo update data user simpan lokasi file
	// return api response

	// c.SaveUploadedFile()

}
