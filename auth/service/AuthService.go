package service

import (
	"errors"
	"main/auth/dto"
	"main/auth/generator"
	"main/model"
	"main/repo"
)

type AuthService struct {
	Repo         *repo.UserRepository
	JwtGenerator *generator.JwtGenerator
}

func (service *AuthService) Login(loginDto *dto.LoginDto) (string, error) {
	var user model.User
	user, err := service.ValidateLoginData(loginDto)
	var token string
	if err == nil {
		token, _ = service.JwtGenerator.GenerateJwt(&user)
		return token, err
	}
	return token, err
}

func (service *AuthService) ValidateLoginData(loginDto *dto.LoginDto) (model.User, error) {
	user, err := service.Repo.GetOne(loginDto.Username)
	if err != nil {
		return user, errors.New("User not found")
	}
	if user.Password != loginDto.Password {
		return user, errors.New("Wrong password")
	}
	return user, nil
}
