package postgres

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewClient() (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", viper.GetString("DB_USER"), viper.GetString("DB_PASSWORD"), viper.GetString("DB_HOST"), "postgres")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database")
	}

	dbName := viper.GetString("DB_NAME")
	checkAndCreateDatabase(db, dbName)

	dsnNew := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", viper.GetString("DB_USER"), viper.GetString("DB_PASSWORD"), viper.GetString("DB_HOST"), dbName)

	db, err = gorm.Open(postgres.Open(dsnNew), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database with new DSN")
	}

	return db, nil
}

func checkAndCreateDatabase(DB *gorm.DB, dbName string) {
	var count int64
	result := DB.Raw("SELECT count(*) FROM pg_database WHERE datname = ?", dbName).Scan(&count)
	if result.Error != nil {
		panic(result.Error)
	}

	if count == 0 {
		createDatabase(DB, dbName)
	} else {
		fmt.Printf("Database %s already exists\n", dbName)
	}
}

func createDatabase(DB *gorm.DB, dbName string) {
	result := DB.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	// DB.Exec(fmt.Sprintf("INSERT INTO migrations (id, name, created_at) VALUES (1, 'create_database_%s', NOW())", dbName))
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Printf("Database %s created\n", dbName)
}
