package usecase

import "fmt"

type GetUsersUseCase interface {
	Execute(user string)
}

type GetUsers struct{}

func NewGetUsers() *GetUsers {
	return &GetUsers{}
}

func (getUsers *GetUsers) Execute() {
	fmt.Println("Users")
}
