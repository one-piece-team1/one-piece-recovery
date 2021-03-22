package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	DBNAME     = "DBNAME"
	DBADDR     = "DBADDR"
	DBPORT     = "DBPORT"
	DBUSER     = "DBUSER"
	DBPASSWORD = "DBPASSWORD"
)

type PgConfig struct {
	Addr     string `json:"addr"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func PgConnectConfig() *PgConfig {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	port, portErr := strconv.Atoi(os.Getenv(DBPORT))
	if portErr != nil {
		log.Fatalf("Port env not found")
	}
	config := PgConfig{
		Addr:     os.Getenv(DBADDR),
		Port:     port,
		Username: os.Getenv(DBUSER),
		Name:     os.Getenv(DBNAME),
		Password: os.Getenv(DBPASSWORD),
	}
	return &config
}
