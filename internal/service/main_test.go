package service

import (
	"os"
	"testing"

	"github.com/alifdwt/haiwan-go/internal/mapper"
	"github.com/alifdwt/haiwan-go/internal/repository"
	"github.com/alifdwt/haiwan-go/pkg/auth"
	"github.com/alifdwt/haiwan-go/pkg/database/migration"
	"github.com/alifdwt/haiwan-go/pkg/database/postgres"
	"github.com/alifdwt/haiwan-go/pkg/dotenv"
	"github.com/alifdwt/haiwan-go/pkg/hashing"
	"github.com/alifdwt/haiwan-go/pkg/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var testQueries *Service

func TestMain(m *testing.M) {
	log, err := logger.NewLogger()
	if err != nil {
		log.Error("Error while initializing logger", zap.Error(err))
	}

	err = dotenv.Viper("../..")
	if err != nil {
		log.Error("Error while loading .env fie", zap.Error(err))
	}

	db, err := postgres.NewClient()
	if err != nil {
		log.Error("Error while connecting to database test", zap.Error(err))
	}

	err = migration.RunMigration(db)
	if err != nil {
		log.Error("Error while migrating to database", zap.Error(err))
	}

	hashing := hashing.NewHashingPassword()
	repository := repository.NewRepository(db)

	token, err := auth.NewManager(viper.GetString("JWT_SECRET"))
	if err != nil {
		log.Error("Error while initializing token manager", zap.Error(err))
	}

	mapper := mapper.NewMapper()

	testQueries = NewService(Deps{
		Repository: repository,
		Hashing:    *hashing,
		Token:      token,
		Logger:     *log,
		Mapper:     *mapper,
	})

	os.Exit(m.Run())
}
