package persistence

import (
	"log"

	"github.com/jinzhu/gorm"
)

func NewConnectionDB() (client *gorm.DB) {
	connectionData := GetConnectionDatabase()

	client, dbError := gorm.Open(connectionData.Dialect, GetConnectionString(connectionData))

	if dbError != nil {
		log.Fatal("Error to connect in database")
	}

	return client
}
