package app

import (
	"social-go/cmd/api/core/usecase"
	"social-go/cmd/api/entrypoint"
	"social-go/cmd/api/persistence"
)

type Controllers struct {
	Ping *entrypoint.PingController
	User *entrypoint.UserController
}

// InitializeHandlers func initialize all handlers and dependencies of app
func InitializeHandlers() *Controllers {

	db, _ := persistence.NewConnectionDB()

	userDao := persistence.NewUserDao(db)

	createUser := usecase.NewCreateUser(*userDao)

	return &Controllers{
		Ping: &entrypoint.PingController{},
		User: entrypoint.NewUserController(*createUser),
	}
}
