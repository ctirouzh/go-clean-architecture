package database

import (
	"lms/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabase_ConnectToPostgres(t *testing.T) {
	// To ensure that a proper postgres config exists
	cfg, cfgErr := config.Parse("../config/config.json")
	assert.Empty(t, cfgErr)
	// To ensure that your config leads to a connection.
	db, err := ConnectToPostgres(cfg.Postgres)
	assert.Empty(t, err)
	assert.NotEmpty(t, db)
}

func TestPostgres_PostgresConnectionError(t *testing.T) {
	cfg, cfgErr := config.Parse("../config/example.config.json")
	assert.Empty(t, cfgErr)
	// To ensure that a wrong config leads to a connection error.
	db, err := ConnectToPostgres(cfg.Postgres)
	assert.NotEmpty(t, err)
	assert.Empty(t, db)
}
