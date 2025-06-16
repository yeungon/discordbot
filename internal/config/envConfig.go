package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
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
