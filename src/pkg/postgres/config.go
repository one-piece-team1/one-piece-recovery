package postgres

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/one-piece-team1/one-piece-recovery/src/config"
)

type DBClient struct {
	client *gorm.DB
}

func (m *DBClient) PgConnect() {
	dbConf := config.PgConnectConfig()

	client, err := gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			dbConf.Addr,
			dbConf.Port,
			dbConf.Username,
			dbConf.Name,
			dbConf.Password,
		),
	)
	if err != nil {
		panic(err)
	}
	m.client = client
}

func (m *DBClient) PgDisConnect() {
	m.client.Close()
}
