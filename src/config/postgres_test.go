package config

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPgConfig(t *testing.T) {
	conf := PgConfig{
		Addr:     "localhost",
		Port:     5432,
		Username: "lib",
		Name:     "lib",
		Password: "123",
	}
	bytes, err := json.Marshal(conf)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var target PgConfig
	err = json.Unmarshal(bytes, &target)

	assert.Nil(t, err)
	assert.NotNil(t, target.Addr, conf.Addr)
	assert.NotNil(t, target.Port, conf.Port)
	assert.NotNil(t, target.Username, conf.Username)
	assert.NotNil(t, target.Name, conf.Name)
	assert.NotNil(t, target.Password, conf.Password)
}
