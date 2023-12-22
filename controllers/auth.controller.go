package controllers

import (
	"net/http"

	"github.com/IrvanWijaya/dealls-tinder/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(DB *gorm.DB) AuthController {
	authServie := services.NewAuthService(DB)

	return AuthController{&authServie}
}

func (ac *AuthController) SignUpUser(ctx *gin.Context) {
	newUser, err := ac.authService.SignUpUser(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	message := "User successfully created." + newUser.Email
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": message})
}

func (ac *AuthController) SignInUser(ctx *gin.Context) {

	token, err := ac.authService.SignInUser(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "token": token})
}

func (ac *AuthController) LogoutUser(ctx *gin.Context) {
	ac.authService.LogoutUser(ctx)
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
