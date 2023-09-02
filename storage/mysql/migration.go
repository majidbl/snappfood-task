package mysql

import (
	"gorm.io/gorm"

	"task/models"
)

var migrationsModels = []interface{}{
	models.Order{},
	models.Vendor{},
	models.Agent{},
	models.DelayReport{},
	models.Trip{},
}

func MigrateUp(db *gorm.DB) error {
	return db.AutoMigrate(migrationsModels...)
}

func MigrateDown(db *gorm.DB) error {
	return db.Migrator().DropTable(migrationsModels...)
}
