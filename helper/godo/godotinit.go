package godo

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() error {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err)
		return err
	}
	return nil
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	return value
}
