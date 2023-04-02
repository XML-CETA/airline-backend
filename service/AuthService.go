package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"main/dtos"
	"main/model"
	"main/repo"
	"main/utils"
	"os"
	"time"
)

type AuthService struct {
	Repo *repo.UserRepository
}

func (service *AuthService) Login(loginDto *dtos.LoginDto) (string, error) {
	var user model.User
	user, err := service.ValidateLoginData(loginDto)
	var token string
	if err == nil {
		token, _ = service.GenerateJwt(&user)
		return token, err
	}
	return token, err
}

func (service *AuthService) ValidateLoginData(loginDto *dtos.LoginDto) (model.User, error) {
	user, err := service.Repo.GetOne(loginDto.Username)
	if err != nil {
		return user, errors.New("User not found")
	}
	if user.Password != loginDto.Password {
		return user, errors.New("Wrong password")
	}
	return user, nil
}

func (service *AuthService) GenerateJwt(user *model.User) (string, error) {

	var secretKey = []byte(os.Getenv("SECRET_KEY"))

	claims := utils.Claims{
		CustomClaims: map[string]string{
			"username": user.Username,
			"role":     user.Role,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		},
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenString.SignedString(secretKey)

	return token, err
}
