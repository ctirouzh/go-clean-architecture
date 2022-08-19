package database

import (
	"log"

	"gorm.io/gorm"
)

var DB *gorm.DB
var models = []interface{}{
	// TODO: Think about these two options:
	// 1. Add gorm tag to domain entities and use them as a orm model here,
	// 2. Define new models here and use them.

	// &user.User{},

}

// Migrate calls Automigrate method of the given gorm db for all database models.
// It also drops unused column internally.
func Migrate(db *gorm.DB) {
	DB = db
	for _, model := range models {
		if err := DB.AutoMigrate(model); err != nil {
			log.Fatal("Fatal: migration", err)
		}
		dropUnusedColumns(model)
	}
	log.Println("database migrated...")
}

// dropUnusedColumns drops unused columns of a given destination.
func dropUnusedColumns(dst interface{}) {

	stmt := &gorm.Statement{DB: DB}
	stmt.Parse(dst)
	fields := stmt.Schema.Fields
	columns, _ := DB.Debug().Migrator().ColumnTypes(dst)
	for i := range columns {
		found := false
		for j := range fields {
			if columns[i].Name() == fields[j].DBName {
				found = true
				break
			}
		}
		if !found {
			DB.Migrator().DropColumn(dst, columns[i].Name())
		}
	}
}
