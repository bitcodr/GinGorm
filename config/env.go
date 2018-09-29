package config

import (
	"log"

	"github.com/joho/godotenv"
)

//Environment func
func Environment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error on loading env file")
	}
}
