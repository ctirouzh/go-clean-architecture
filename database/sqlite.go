package database

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	sqliteInstance *sql.DB
	sqliteOnce     sync.Once
)

func GetSqliteInstance() *sql.DB {
	if sqliteInstance == nil {
		sqliteOnce.Do(func() {
			db, err := sql.Open("sqlite3", "./my.db")
			if err != nil {
				log.Panic("error in establishing connection with sqlite3 database", err.Error())
			}
			defer db.Close()

			sqliteInstance = db
		})
	}

	return sqliteInstance
}
