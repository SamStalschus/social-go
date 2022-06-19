package app

import (
	"social-go/cmd/api/core/usecase"
	"social-go/cmd/api/entrypoint"
	"social-go/cmd/api/persistence"
)

type Controllers struct {
	User *entrypoint.UserController
	Auth *entrypoint.AuthController
}

// InitializeHandlers func initialize all handlers and dependencies of app
func InitializeHandlers() *Controllers {

	db := persistence.NewConnectionDB()

	userDao := persistence.NewUserDao(db)

	createUser := usecase.NewCreateUser(*userDao)
	getUser := usecase.NewGetUser(*userDao)
	getUsers := usecase.NewGetUsers(*userDao)
	updateUser := usecase.NewUpdateUser(*userDao)
	deleteUser := usecase.NewDeleteUsers(*userDao)

	return &Controllers{
		User: entrypoint.NewUserController(*createUser, *getUser, *getUsers, *updateUser, *deleteUser),
		Auth: entrypoint.NewAuthController(),
	}
}
