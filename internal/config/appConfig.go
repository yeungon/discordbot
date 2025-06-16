package config

import (
	"log"

	db "github.com/yeungon/discordbot/internal/pg"
)

type AppConfig struct {
	UseCache      bool
	ErrorLog      *log.Logger
	InfoLog       *log.Logger
	InProduction  bool
	AUTH_USER     string
	AUTH_PASSWORD string
	AppName       string
	Query         *db.Queries
	Debug         bool
}

func NewApp(cacheState bool, ProductionState bool) *AppConfig {
	return &AppConfig{
		UseCache:     cacheState,
		InProduction: ProductionState,
	}

}
