package migrations

import (
	"github.com/shawon325/go-crud/src/models"
	"gorm.io/gorm"
)

var registeredModels = []any{
	&models.User{},
}

func Run(db *gorm.DB) error {
	return db.AutoMigrate(registeredModels...)
}
