package service

import (
	"github.com/alifdwt/haiwan-go/internal/repository"
	"github.com/alifdwt/haiwan-go/pkg/auth"
	hashing "github.com/alifdwt/haiwan-go/pkg/hasing"
	"github.com/alifdwt/haiwan-go/pkg/logger"
)

type Service struct {
	Auth AuthService
	User UserService
}

type Deps struct {
	Repository *repository.Repositories
	Hashing    hashing.Hashing
	Token      auth.TokenManager
	Logger     logger.Logger
}

func NewService(deps Deps) *Service {
	return &Service{
		Auth: NewAuthService(deps.Repository.User, deps.Hashing, deps.Logger, deps.Token),
		User: NewUserService(deps.Repository.User, deps.Hashing, deps.Logger),
	}
}
