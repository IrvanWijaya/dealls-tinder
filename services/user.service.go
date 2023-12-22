package services

import (
	"github.com/IrvanWijaya/dealls-tinder/models"
	"github.com/gin-gonic/gin"
)

type UserService struct {
}

func NewUserService() UserService {
	return UserService{}
}

func (ac *UserService) GetMe(ctx *gin.Context) (*models.UserResponse, error) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	userResponse := &models.UserResponse{
		ID:        currentUser.ID,
		Name:      currentUser.Name,
		Email:     currentUser.Email,
		IsPremium: currentUser.IsPremium,
		CreatedAt: currentUser.CreatedAt,
		UpdatedAt: currentUser.UpdatedAt,
	}

	return userResponse, nil
}
