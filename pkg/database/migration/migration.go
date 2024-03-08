package migration

import (
	"fmt"
	"sort"
	"strings"

	"github.com/alifdwt/haiwan-go/internal/models"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Cart{},
		&models.Category{},
		&models.Order{},
		&models.OrderItems{},
		&models.Product{},
		&models.Review{},
		&models.ShippingAddress{},
		&models.Slider{},
		&models.User{},
		&models.Migration{},
	)

	var existingTables []string
	db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Scan(&existingTables)

	sort.Strings(existingTables)

	for _, tableName := range existingTables {
		migrationName := fmt.Sprintf("create_%s_table", strings.ToLower(tableName))
		if !isTableMigrated(db, migrationName) {
			err := insertMigrationRecord(db, migrationName)
			if err != nil {
				return err
			}
		}
	}

	return err
}

func isTableMigrated(db *gorm.DB, migrationName string) bool {
	var count int64
	db.Raw("SELECT COUNT(*) FROM migrations WHERE name = ?", migrationName).Scan(&count)
	return count > 0
}

func insertMigrationRecord(db *gorm.DB, migrationName string) error {
	result := db.Exec("INSERT INTO migrations (name, created_at) VALUES (?, NOW())", migrationName)
	if result.Error != nil {
		return result.Error
	}
	return nil
}