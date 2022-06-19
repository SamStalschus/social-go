package config

import (
	"context"
	"social-go/cmd/api/app"
	"time"

	"github.com/shaj13/go-guardian/auth"
	"github.com/shaj13/go-guardian/auth/strategies/basic"
	"github.com/shaj13/go-guardian/auth/strategies/bearer"
	"github.com/shaj13/go-guardian/store"
)

func setupGoGuardian() {
	authenticator := auth.New()
	cache := store.NewFIFO(context.Background(), time.Minute*5)
	basicStrategy := basic.New(app.ValidateUser, cache)
	tokenStrategy := bearer.New(app.VerifyToken, cache)
	authenticator.EnableStrategy(basic.StrategyKey, basicStrategy)
	authenticator.EnableStrategy(bearer.CachedStrategyKey, tokenStrategy)
}
