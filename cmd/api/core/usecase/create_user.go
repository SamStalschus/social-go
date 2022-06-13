package usecase

import (
	"log"
	"social-go/cmd/api/core/model"
	"social-go/cmd/api/persistence"
)

type CreateUserUseCase interface {
	Execute(user string) string
}

type CreateUser struct {
	userDao persistence.UserDao
}

func NewCreateUser(userDao persistence.UserDao) *CreateUser {
	return &CreateUser{
		userDao: userDao,
	}
}

func (createUser *CreateUser) Execute(user model.User) (*model.User, error) {

	userCreated, err := createUser.userDao.Save(&user)

	if err != nil {
		log.Fatal("Error to create a new user")
	}

	return userCreated, nil
}
