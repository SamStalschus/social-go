package persistence

import "gorm.io/gorm"

func NewConnectionDB() (client *gorm.DB) {
	connectionData := GetConnectionDatabase()

	client, dbError := gorm.Open(connectionData.Dialect, GetConnectionString(connectionData))
}
