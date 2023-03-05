package pkg

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvironment() {
	if err := godotenv.Load(".env"); err != nil {
		log.Panicf("panic occurred while loading environment : %v", err)
	}
}
