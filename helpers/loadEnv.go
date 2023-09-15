package helpers

import (
	"log"

	"github.com/joho/godotenv"
)

func GetEnvironmentVariable() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}
