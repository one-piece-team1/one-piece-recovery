package config

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	APPPORT = "APPPORT"
)

func GetPort() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	return os.Getenv(APPPORT)
}
