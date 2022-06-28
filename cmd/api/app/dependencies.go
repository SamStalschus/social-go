package app

import (
	"social-go/cmd/api/core/usecase"
	"social-go/cmd/api/entrypoint"
	"social-go/cmd/api/persistence"
	"social-go/cmd/api/utils/middleware"
)

type Controllers struct {
	User *entrypoint.UserController
	Auth *entrypoint.AuthController
}

type Middlwares struct {
	EnsureAuthenticated *middleware.EnsureAuthenticated
}

// InitializeHandlers func initialize all handlers and dependencies of app
func InitializeHandlers() (*Controllers, *Middlwares) {

	db := persistence.NewConnectionDB()

	userDao := persistence.NewUserDao(db)

	genToken := usecase.NewGenTokenUseCase(*userDao)

	createUser := usecase.NewCreateUser(*userDao)
	getUser := usecase.NewGetUser(*userDao)
	getUsers := usecase.NewGetUsers(*userDao)
	updateUser := usecase.NewUpdateUser(*userDao)
	deleteUser := usecase.NewDeleteUsers(*userDao)

	ensureAuthenticated := middleware.NewEnsureAuthenticated()

	return &Controllers{
			User: entrypoint.NewUserController(createUser, getUser, getUsers, updateUser, deleteUser),
			Auth: entrypoint.NewAuthController(genToken),
		}, &Middlwares{
			EnsureAuthenticated: ensureAuthenticated,
		}
}
