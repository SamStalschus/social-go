package config

import (
	"log"

	"github.com/joho/godotenv"
)

// SetupEnvironment initialize variables of environment
func SetupEnvironment() {

	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal(err)
	}

}
