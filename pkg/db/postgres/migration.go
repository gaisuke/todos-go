package postgres

import (
	"todos-go/models"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.List{}, &models.Sublist{}); err != nil {
		return err
	}
	return nil
}
