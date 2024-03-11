package postgres

import (
	"fmt"

	"github.com/alifdwt/haiwan-go/pkg/dotenv"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewClient() (*gorm.DB, error) {
	return createClient(viper.GetString("DB_USER"), viper.GetString("DB_PASSWORD"), viper.GetString("DB_HOST"), viper.GetString("DB_NAME"))
}

func NewClientTest() (*gorm.DB, error) {
	config, err := dotenv.LoadConfig("../..")
	if err != nil {
		log.Error("cannot load .env file", err)
	}

	dbName := fmt.Sprintf("%s_test", config.DB_NAME)

	return createClient(config.DB_USER, config.DB_PASSWORD, config.DB_HOST, dbName)
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

func createClient(dbUser string, dbPassword string, dbHost string, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, "postgres")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database postgres")
	}

	checkAndCreateDatabase(db, dbName)

	dsnNew := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbName)

	db, err = gorm.Open(postgres.Open(dsnNew), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database with new DSN")
	}

	return db, nil
}
