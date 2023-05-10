package handlers
package handlers

import "sorafuru/user"

type userHandler struct {
	userService user.Service
}

func newUserHandlers (userService, user.Service) *userHandler{
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context)