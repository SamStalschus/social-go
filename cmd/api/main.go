package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"social-go/cmd/api/app"
	"social-go/cmd/api/config"
)

func main() {

	config.SetupEnvironment()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	if err := run(port); err != nil {
		log.Fatal("Error to start server on port: " + port)
	}

	fmt.Println("Listening on port", port)
}

func run(port string) error {

	handlers, middlewares := app.InitializeHandlers()

	router := gin.New()

	app.MapRoutes(router, handlers, middlewares)

	return router.Run(":" + port)
}
