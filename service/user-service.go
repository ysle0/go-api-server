package service

import (
	"ysl.com/go-api-server/repository"
	"ysl.com/go-api-server/types"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	s := &UserService{userRepository}

	return s
}

func (u *UserService) GetAll() []*types.User {
	allUsers := u.userRepository.GetAll()
	return allUsers
}

func (u *UserService) Create(newUser *types.User) error {
	return u.userRepository.Create(newUser)
}

func (u *UserService) Update(name string, newAge int64) error {
	return u.userRepository.Update(name, newAge)
}

func (u *UserService) Delete(name string) error {
	return u.userRepository.Delete(name)
}
