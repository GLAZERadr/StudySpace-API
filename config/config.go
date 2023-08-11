package config

import (
	"os"
	"log"

	"github.com/joho/godotenv"
)

//function to take host value from .env file
func ConfigHost() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("HOST")
}

//function to take port value from .env file
func ConfigPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("PORT")
}

//function to take mongodb url value from .env file
func ConfigDB() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("DATABASE_URL")
}

func ConfigDBname() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("DATABASE_NAME")
}

//function to take openai api key value from .env file
func ConfigOpenAI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("OPENAI_API_KEY")
}