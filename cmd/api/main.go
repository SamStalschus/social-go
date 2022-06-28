package main

import (
	"fmt"
	"log"
	"os"
	"social-go/cmd/api/app"
	"social-go/cmd/api/config"

	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	if err := run(port); err != nil {
		log.Fatal("Error to start server on port: " + port)
	}

	fmt.Println("Listening on port 3000...")
}

func run(port string) error {

	config.SetupEnvironment()

	handlers, middlewares := app.InitializeHandlers()

	router := gin.New()

	app.MapRoutes(router, handlers, middlewares)

	return router.Run(":" + port)
}
