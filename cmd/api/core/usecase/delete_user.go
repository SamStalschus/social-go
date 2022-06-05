package usecase

import "fmt"

type DeleteUserUseCase interface {
	Execute(user string)
}

type DeleteUser struct{}

func NewDeleteUsers() *DeleteUser {
	return &DeleteUser{}
}

func (deleteUser *DeleteUser) Execute() {
	fmt.Println("Delete users")
}
