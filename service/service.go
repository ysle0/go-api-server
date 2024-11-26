package service

import (
	"sync"

	"ysl.com/go-api-server/repository"
)

var (
	initOnce sync.Once
	instance *Service
)

type Service struct {
	repository *repository.Repository

	User *UserService
}

func New(repository *repository.Repository) *Service {
	userService := NewUserService(repository.User)

	initOnce.Do(func() {
		instance = &Service{
			repository: repository,
			User:       userService,
		}
	})

	return instance
}
