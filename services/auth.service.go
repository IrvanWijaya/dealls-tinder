package services

import (
	goErrors "errors"
	"strings"
	"time"

	"github.com/IrvanWijaya/dealls-tinder/initializers"
	"github.com/IrvanWijaya/dealls-tinder/models"
	"github.com/IrvanWijaya/dealls-tinder/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var (
	fnUHashPassword = utils.HashPassword
	fnTimeNow       = time.Now
)

type DBItf interface {
	Create(value interface{}) (tx *gorm.DB)
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
}

type AuthService struct {
	DB DBItf
}

func NewAuthService(DB *gorm.DB) AuthService {
	return AuthService{DB}
}

func (ac *AuthService) SignUpUser(ctx *gin.Context) (*models.User, error) {
	var payload *models.SignUpInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		return nil, errors.Wrap(err, "Error ShouldBindJSON")
	}

	if payload.Password != payload.PasswordConfirm {
		return nil, goErrors.New("Passwords do not match")
	}

	hashedPassword, err := fnUHashPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	now := fnTimeNow()
	newUser := models.User{
		Name:      payload.Name,
		Email:     strings.ToLower(payload.Email),
		Password:  hashedPassword,
		IsPremium: false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	ac.DB.Create(&newUser)
	return &newUser, nil
}

func (ac *AuthService) SignInUser(ctx *gin.Context) (string, error) {
	var payload *models.SignInInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		return "", err
	}

	var user models.User
	result := ac.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		return "", result.Error
	}

	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		return "", err
	}

	config, _ := initializers.LoadConfig(".")

	token, err := utils.GenerateToken(config.TokenExpiresIn, user.ID, config.TokenSecret)
	if err != nil {
		return "", err
	}

	ctx.SetCookie("token", token, config.TokenMaxAge*60, "/", "localhost", false, true)

	return token, nil
}

func (ac *AuthService) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
}
