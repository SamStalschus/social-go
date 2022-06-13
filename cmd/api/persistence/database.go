package persistence

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // need to load mysql driver on api
	"github.com/jinzhu/gorm"
)

func NewConnectionDB() (*gorm.DB, error) {
	connectionData := GetConnectionDatabase()

	fmt.Println(connectionData)
	client, dbError := gorm.Open(connectionData.Dialect, GetConnectionString(connectionData))

	if dbError != nil {
		// TODO: Handling error
		log.Fatal("Error to connect in database")
	}

	client.SingularTable(true)

	return client, nil
}
