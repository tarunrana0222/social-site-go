package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	isEnvLoaded = false
	DB_Url      string
	DB_Name     string
	JWT_Secret  string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error while loading env", err)
	}
	DB_Url = os.Getenv("MONGO_URL")
	DB_Name = os.Getenv("DB_NAME")
	JWT_Secret = os.Getenv("JWT_SECRET")
	isEnvLoaded = true

	fmt.Println("---------- Loaded ENV variables ------")
}
