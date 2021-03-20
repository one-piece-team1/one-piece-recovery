package config

import (
	"github.com/joho/godotenv"
	"os"
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
