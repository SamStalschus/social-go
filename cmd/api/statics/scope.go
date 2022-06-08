package statics

import "os"

const (
	Scope      = "SCOPE"
	ScopeProd  = "prod"
	ScopeLocal = "local"
	ScopeTest  = "test"
)

func GetScope() string {
	return os.Getenv(Scope)
}
