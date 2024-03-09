package repository

import (
	"github.com/alifdwt/haiwan-go/pkg/database/postgres"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

func setDatabase() *gorm.DB {
	db, err := postgres.NewClient()
	if err != nil {
		log.Error("error while connecting to database " + err.Error())
	}

	return db
}
