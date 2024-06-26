package service

import (
	"github.com/alifdwt/haiwan-go/internal/mapper"
	"github.com/alifdwt/haiwan-go/internal/repository"
	"github.com/alifdwt/haiwan-go/pkg/auth"
	"github.com/alifdwt/haiwan-go/pkg/hashing"
	"github.com/alifdwt/haiwan-go/pkg/logger"
	midtranspkg "github.com/alifdwt/haiwan-go/pkg/midtrans_pkg"
	"github.com/alifdwt/haiwan-go/pkg/rajaongkir"
)

type Store interface {
	Service
}

type Service struct {
	Auth       AuthService
	User       UserService
	Category   CategoryService
	Product    ProductService
	Cart       CartService
	Order      OrderService
	Review     ReviewService
	Slider     SliderService
	RajaOngkir RajaOngkirService
	Midtrans   MidtransService
}

type Deps struct {
	Repository *repository.Repositories
	Hashing    hashing.Hashing
	Token      auth.TokenManager
	Logger     logger.Logger
	Mapper     mapper.Mapper
	RajaOngkir *rajaongkir.RajaOngkirAPI
	Snap       *midtranspkg.SnapClient
}

func NewService(deps Deps) *Service {
	return &Service{
		Auth:       NewAuthService(deps.Repository.User, deps.Hashing, deps.Logger, deps.Token),
		User:       NewUserService(deps.Repository.User, deps.Hashing, deps.Logger),
		Category:   NewCategoryService(deps.Repository.Category, deps.Mapper.CategoryMapper),
		Product:    NewProductService(deps.Repository.Product, deps.Logger, deps.Mapper.ProductMapper),
		Cart:       NewCartService(deps.Repository.Cart, deps.Logger, deps.Mapper.CartMapper),
		Order:      NewOrderService(deps.Repository.Order, deps.Logger, deps.Mapper.OrderMapper),
		Review:     NewReviewService(deps.Repository.Review),
		Slider:     NewSliderService(deps.Repository.Slider, deps.Logger),
		RajaOngkir: NewRajaOngkirService(deps.RajaOngkir),
		Midtrans:   NewMidtransService(deps.Snap),
	}
}
