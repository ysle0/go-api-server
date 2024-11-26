package repository

import (
	"ysl.com/go-api-server/types"
	"ysl.com/go-api-server/types/errors"
)

type UserRepository struct {
	userMap []*types.User
}

func NewUserRepository() *UserRepository {
	r := &UserRepository{
		userMap: []*types.User{
			&types.User{
				Name: "yoonsang",
				Age:  27,
			},
		},
	}

	return r
}

func (u *UserRepository) GetAll() []*types.User {
	return u.userMap
}

func (u *UserRepository) GetOne(newUser *types.User) error {
	return nil
}

func (u *UserRepository) Create(newUser *types.User) error {
	u.userMap = append(u.userMap, newUser)
	return nil
}

func (u *UserRepository) Update(name string, newAge int64) error {
	isExisting := false
	for _, user := range u.userMap {
		if user.Name == name {
			user.Age = newAge
			isExisting = true
			continue
		}
	}

	if !isExisting {
		return errors.Errorf(errors.NotFoundUser, nil)
	}
	return nil
}

func (u *UserRepository) Delete(name string) error {
	isExisting := false
	for i, user := range u.userMap {
		if user.Name == name {
			u.userMap = append(u.userMap[:i], u.userMap[i+1:]...)
			isExisting = true
			break
		}
	}

	if !isExisting {
		return errors.Errorf(errors.NotFoundUser, nil)
	}
	return nil
}
