package repository

import (
	"sync"
)

var (
	initOnce sync.Once
	instance *Repository
)

type Repository struct {
	User *UserRepository
}

func New() *Repository {
	initOnce.Do(func() {
		instance = &Repository{
			User: NewUserRepository(),
		}
	})

	return instance
}
