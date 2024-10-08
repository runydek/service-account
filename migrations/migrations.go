package migrations

import (
	"service-account/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Account{}, &models.TransactionRequest{}, &models.RegisterRequest{}, &models.TransactionHistory{}); err != nil {
		return err
	}
	return nil
}
