package app

import (
	"github.com/alifdwt/haiwan-go/internal/handler"
	"github.com/alifdwt/haiwan-go/internal/mapper"
	"github.com/alifdwt/haiwan-go/internal/repository"
	"github.com/alifdwt/haiwan-go/internal/service"
	"github.com/alifdwt/haiwan-go/pkg/auth"
	"github.com/alifdwt/haiwan-go/pkg/cloudinary"
	"github.com/alifdwt/haiwan-go/pkg/database/migration"
	"github.com/alifdwt/haiwan-go/pkg/database/postgres"
	"github.com/alifdwt/haiwan-go/pkg/dotenv"
	"github.com/alifdwt/haiwan-go/pkg/hashing"
	"github.com/alifdwt/haiwan-go/pkg/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// @title HaiwanGo
// @version 1.0
// @description REST API for Haiwan Go Fiber App

// @contact.name Alif Dewantara
// @contact.url https://github.com/alifdwt
// @contact.email aputradewantara@gmail.com

// @host localhost:8000
// @BasePath /api/

// @securityDefinitions.apikey BearerAuth
// @in Header
// @name Authorization
func Run() {
	log, err := logger.NewLogger()
	if err != nil {
		log.Error("Error while initializing logger", zap.Error(err))
	}

	err = dotenv.Viper()
	if err != nil {
		log.Error("Error while loading .env fie", zap.Error(err))
	}

	db, err := postgres.NewClient()
	if err != nil {
		log.Error("Error while connecting to database", zap.Error(err))
	}

	err = migration.RunMigration(db)
	if err != nil {
		log.Error("Error while migrating to database", zap.Error(err))
	}

	myCloudinary, err := cloudinary.NewMyCloudinary()
	if err != nil {
		log.Error("Error while connecting to cloudinary", zap.Error(err))
	}

	hashing := hashing.NewHashingPassword()
	repository := repository.NewRepository(db)

	token, err := auth.NewManager(viper.GetString("JWT_SECRET"))
	if err != nil {
		log.Error("Error while initializing token manager", zap.Error(err))
	}

	mapper := mapper.NewMapper()

	service := service.NewService(service.Deps{
		Repository: repository,
		Hashing:    *hashing,
		Token:      token,
		Logger:     *log,
		Mapper:     *mapper,
	})

	myHandler := handler.NewHandler(service, *myCloudinary, token)

	myHandler.Init().Listen(":8000")
}
