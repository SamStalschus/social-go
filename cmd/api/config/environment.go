package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "social-go"

// SetupEnvironment initialize variables of environment
func SetupEnvironment() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	if err := godotenv.Load(string(rootPath) + `/.env`); err != nil {
		log.Fatal(err)
	}

}
