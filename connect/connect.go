package connect

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Main - No .env file found")
	}
}

func LookupEnvOrExit(key string) string {
	value, success := os.LookupEnv(key)
	if !success {
		log.Print("Could not obtain " + key)
		os.Exit(1)
	}
	return value
}
