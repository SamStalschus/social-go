package persistence

import (
	"fmt"
	"os"
	"social-go/cmd/api/statics"
)

type ConnectionData struct {
	User     string
	Password string
	Schema   string
	Dialect  string
}

const MYSQL = "mysql"

// GetConnectionDatabase return a connection of database based on scope
func GetConnectionDatabase() *ConnectionData {
	scope := statics.GetScope()

	connectionData := ConnectionData{}

	if scope == statics.ScopeLocal {
		return connectionData.setupLocalConnectionData()
	}

	return &connectionData
}

func (cd *ConnectionData) setupLocalConnectionData() *ConnectionData {
	cd.User = os.Getenv("DB_USER")
	cd.Password = os.Getenv("DB_PASSWORD")
	cd.Schema = os.Getenv("DB_SCHEMA")
	cd.Dialect = MYSQL

	return cd
}

func GetConnectionString(cd *ConnectionData) string {
	return fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", cd.User, cd.Password, cd.Schema)
}
