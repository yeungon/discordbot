package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Env struct {
	Token          string
	AdminID        string
	Passcode       string
	PostgresURL    string
	SphUsername    string
	SphPassword    string
	StudentList    string
	SphURLEndpoint string
	SecretFirst    string
	SecretSecond   string
	KhoaluanURL    string
}

var (
	once sync.Once
	env  *Env
)

func New() {
	once.Do(func() {
		var PostgresURLDynamic string
		// the variable ${PRODUCTION} will be set to TRUE inside docker-compose.yml, false by default (set in .env)
		if os.Getenv("PRODUCTION") != "TRUE" {
			err := godotenv.Load(".env")
			if err != nil {
				log.Fatal("Error loading .env file")
			}
			PostgresURLDynamic = os.Getenv("POSTGRES_URL_DEV")
		} else {
			PostgresURLDynamic = os.Getenv("POSTGRES_URL")
		}

		env = &Env{
			Token:          os.Getenv("DISCORD_BOT_TOKEN"),
			AdminID:        os.Getenv("ADMIN_ID"),
			Passcode:       os.Getenv("PASSCODE"),
			PostgresURL:    PostgresURLDynamic,
			SphUsername:    os.Getenv("SPH_USERNAME"),
			SphPassword:    os.Getenv("SPH_PASSWORD"),
			StudentList:    os.Getenv("STUDENT_LIST"),
			SphURLEndpoint: os.Getenv("SPH_URL_ENDPOINT"),
			SecretFirst:    os.Getenv("SECRET_FIRST"),
			SecretSecond:   os.Getenv("SECRET_SECOND"),
			KhoaluanURL:    os.Getenv("KHOALUAN_URL"),
		}
	})
}

func Get() *Env {
	if env == nil {
		log.Fatal("Environment not initialized. Call config.New() first.")
	}
	return env
}
