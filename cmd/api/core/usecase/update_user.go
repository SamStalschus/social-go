package usecase

import "fmt"

type UpdateUserUseCase interface {
	Execute(user string)
}

type UpdateUser struct{}

func NewUpdateUser() *UpdateUser {
	return &UpdateUser{}
}

func (updateUser *UpdateUser) Execute() {
	fmt.Println("Update user")
}
