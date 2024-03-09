package service

import (
	"github.com/alifdwt/haiwan-go/internal/mapper"
	"github.com/alifdwt/haiwan-go/internal/repository"
	"github.com/alifdwt/haiwan-go/pkg/auth"
	"github.com/alifdwt/haiwan-go/pkg/hashing"
	"github.com/alifdwt/haiwan-go/pkg/logger"
)

type Service struct {
	Auth     AuthService
	User     UserService
	Category CategoryService
	Product  ProductService
}

type Deps struct {
	Repository *repository.Repositories
	Hashing    hashing.Hashing
	Token      auth.TokenManager
	Logger     logger.Logger
	Mapper     mapper.Mapper
}

func NewService(deps Deps) *Service {
	return &Service{
		Auth:     NewAuthService(deps.Repository.User, deps.Hashing, deps.Logger, deps.Token),
		User:     NewUserService(deps.Repository.User, deps.Hashing, deps.Logger),
		Category: NewCategoryService(deps.Repository.Category, deps.Mapper.CategoryMapper),
		Product:  NewProductService(deps.Repository.Product, deps.Logger, deps.Mapper.ProductMapper),
	}
}
