package usecase

import "fmt"

type GetUserUseCase interface {
	Execute(user string)
}

type GetUser struct{}

func NewGetUser() *CreateUser {
	return &CreateUser{}
}

func (GetUser *GetUser) Execute() {
	fmt.Println("User")
}
