package orm

import (
	"log"

	"gorm.io/gorm"
)

type ORM struct {
	DB     *gorm.DB
	models []interface{}
}

// NewORM returns a pointer to ORM, which uses its models and migrate them to the given database.
// It drops the unused columns of the database tables internally.
func NewORM(db *gorm.DB) *ORM {
	orm := &ORM{DB: db}
	orm.setModels()
	orm.migrate()
	return orm
}

func (orm *ORM) setModels() {
	// TODO: Think about these two options:
	// 1. Add gorm tag to domain entities and use them as a orm model here,
	// 2. Define new models here and use them.
	orm.models = []interface{}{
		// &user.User{},
	}
}

// Migrate calls Automigrate method of the orm db for all orm models.
// It also drops unused column internally.
func (orm *ORM) migrate() {
	for _, model := range orm.models {
		if err := orm.DB.AutoMigrate(model); err != nil {
			log.Fatal("Fatal: migration", err)
		}
		orm.dropUnusedColumns(model)
	}
	log.Println("database migrated...")
}

// dropUnusedColumns drops unused columns of a given destination.
func (orm *ORM) dropUnusedColumns(dst interface{}) {

	stmt := &gorm.Statement{DB: orm.DB}
	stmt.Parse(dst)
	fields := stmt.Schema.Fields
	columns, _ := orm.DB.Debug().Migrator().ColumnTypes(dst)
	for i := range columns {
		found := false
		for j := range fields {
			if columns[i].Name() == fields[j].DBName {
				found = true
				break
			}
		}
		if !found {
			orm.DB.Migrator().DropColumn(dst, columns[i].Name())
		}
	}
}
