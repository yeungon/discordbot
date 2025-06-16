package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Env struct {
	Token            string
	AdminID          string
	PASSCODE         string
	PG               string
	SPH_USERNAME     string
	SPH_PASSWORD     string
	STUDENT_LIST     string
	SPH_URL_ENDPOINT string
	SECRET_FIRST     string
	SECRET_SECOND    string
}

var once sync.Once
var env *Env

func New() {
	once.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		token := os.Getenv("TOKEN")
		admin_id := os.Getenv("ADMIN_ID")
		pass_code := os.Getenv("PASSCODE")
		postgresq := os.Getenv("PG")
		sph_username := os.Getenv("SPH_USERNAME")
		sph_password := os.Getenv("SPH_PASSWORD")
		student_list := os.Getenv("STUDENT_LIST")
		sph_url_endpoint := os.Getenv("SPH_URL_ENDPOINT")
		secret_first := os.Getenv("SECRET_FIRST")
		secret_second := os.Getenv("SECRET_SECOND")

		env = &Env{
			Token:            token,
			AdminID:          admin_id,
			PASSCODE:         pass_code,
			PG:               postgresq,
			SPH_USERNAME:     sph_username,
			SPH_PASSWORD:     sph_password,
			STUDENT_LIST:     student_list,
			SPH_URL_ENDPOINT: sph_url_endpoint,
			SECRET_FIRST:     secret_first,
			SECRET_SECOND:    secret_second,
		}

	})
}

func Get() *Env {
	return env
}

func LoadEnv() (string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("LoadEnv Error loading .env file")
	}
	bot_token := os.Getenv("DISCORD_BOT_TOKEN")
	postgres_url := os.Getenv("POSTGRES_URL")
	return bot_token, postgres_url
}

func Token() string {
	var Token string
	production := os.Getenv("PRODUCTION")
	if production == "TRUE" {
		Token = os.Getenv("DISCORD_BOT_TOKEN")
	} else {
		bot_token, _ := LoadEnv()
		Token = bot_token
	}
	return Token
}

func PostgreSql_URL() string {
	var URL string
	production := os.Getenv("PRODUCTION")
	if production == "TRUE" {
		URL = os.Getenv("POSTGRES_URL")
	} else {
		_, URL_ENV := LoadEnv()
		URL = URL_ENV
	}
	return URL
}
