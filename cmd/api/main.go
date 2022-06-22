package main

import (
	"fmt"
	"log"
	"os"
	"social-go/cmd/api/app"
	"social-go/cmd/api/config"
	"social-go/cmd/api/utils/middleware"

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

	handlers := app.InitializeHandlers()

	router := gin.New()

	router.Use(middleware.GinErrorInterceptor())

	app.MapRoutes(router, handlers)

	return router.Run(":" + port)
}
