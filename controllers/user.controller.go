package controllers

import (
	"net/http"

	"github.com/IrvanWijaya/dealls-tinder/models"
	"github.com/IrvanWijaya/dealls-tinder/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController() UserController {
	userService := services.NewUserService()

	return UserController{&userService}
}

func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	userResponse := &models.UserResponse{
		ID:        currentUser.ID,
		Name:      currentUser.Name,
		Email:     currentUser.Email,
		IsPremium: currentUser.IsPremium,
		CreatedAt: currentUser.CreatedAt,
		UpdatedAt: currentUser.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}
