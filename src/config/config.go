package config

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	APPPORT  = "APPPORT"
	LOGLEVEL = "LOGLEVEL"
)

func GetPort() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	return os.Getenv(APPPORT)
}

func GetLogLevel() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	return os.Getenv(LOGLEVEL)
}

func IsProd() bool {
	return GetLogLevel() == "production" || GetLogLevel() == "prod"
}
