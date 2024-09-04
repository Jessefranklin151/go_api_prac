package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func GetTrustedProxies() []string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	trustedProxies := os.Getenv("TRUSTED_PROXIES")

	return strings.Split(trustedProxies, ",")

}
