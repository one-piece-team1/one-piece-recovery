package config

import (
	"context"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

const (
	DBADDR     = "DBADDR"
	DBUSER     = "DBUSER"
	DBPASSWORD = "DBPASSWORD"
	DBDATABASE = "DBDATABASE"
)

func PgConnect() *pg.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv(DBADDR),
		User:     os.Getenv(DBUSER),
		Password: os.Getenv(DBPASSWORD),
		Database: os.Getenv(DBDATABASE),
	})
	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		panic(err)
	}
	return db
}
