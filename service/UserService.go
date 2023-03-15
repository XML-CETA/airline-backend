package service

import (
	"main/model"
	"main/repo"
	"errors"
)

type UserService struct {
	Repo *repo.UserRepository
}

func (service *UserService) Create(user *model.User) error {
	_, err := service.Repo.GetOne(user.Username)
	if err == nil {
		return errors.New("User with this username already exists")
	}
	return service.Repo.Create(user)
}

func (service *UserService) GetOne(username string) (model.User, error) {
	return service.Repo.GetOne(username)
}
