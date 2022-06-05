package usecase

import "fmt"

type CreateUserUseCase interface {
	Execute(user string)
}

type CreateUser struct{}

func NewCreateUser() *CreateUser {
	return &CreateUser{}
}

func (createUser *CreateUser) Execute(user string) {
	fmt.Println("User created", user)
}
