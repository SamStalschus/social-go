package app

import "social-go/cmd/api/entrypoint"

type Controllers struct {
	Ping *entrypoint.PingController
	User *entrypoint.UserController
}

// InitializeHandlers func initialize all handlers and dependencies of app
func InitializeHandlers() *Controllers {

	return &Controllers{
		Ping: &entrypoint.PingController{},
		User: entrypoint.NewUserController(),
	}
}
