package database

import (
	"fmt"
	"lms/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToPostgres(cfg config.Postgres) (*gorm.DB, error) {

	log.Printf(
		"PSQL_HOST:%s PSQL_PORT:%d PSQL_USER:%s PSQL_NAME:%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Name,
	)

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.Name,
	)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	return db, nil
}
