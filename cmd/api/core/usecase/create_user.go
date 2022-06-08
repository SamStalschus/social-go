package usecase

type CreateUserUseCase interface {
	Execute(user string) string
}

type CreateUser struct{}

func NewCreateUser() *CreateUser {
	return &CreateUser{}
}

func (createUser *CreateUser) Execute(user string) string {
	string := "User created"
	return string
}
