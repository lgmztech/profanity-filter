package env

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading env file")
		panic(err)
	}
	fmt.Println("Environment variables loaded")
}
