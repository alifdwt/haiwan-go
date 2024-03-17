package service

import (
	"github.com/alifdwt/haiwan-go/internal/domain/requests/auth"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/cart"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/category"
	midtransrequest "github.com/alifdwt/haiwan-go/internal/domain/requests/midtrans_request"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/order"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/product"
	rajaongkirrequest "github.com/alifdwt/haiwan-go/internal/domain/requests/rajaongkir_request"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/review"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/slider"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/user"
	"github.com/alifdwt/haiwan-go/internal/domain/responses"
	catresponse "github.com/alifdwt/haiwan-go/internal/domain/responses/category"
	or "github.com/alifdwt/haiwan-go/internal/domain/responses/order"
	pro "github.com/alifdwt/haiwan-go/internal/domain/responses/product"
	"github.com/alifdwt/haiwan-go/internal/models"
	"github.com/midtrans/midtrans-go/snap"
)

type AuthService interface {
	Register(input *auth.RegisterRequest) (*models.User, error)
	Login(input *auth.LoginRequest) (*responses.Token, error)
	RefreshToken(req auth.RefreshTokenRequest) (*responses.Token, error)
}

type UserService interface {
	GetUserAll() (*[]models.User, error)
	CreateUser(request *auth.RegisterRequest) (*models.User, error)
	GetUserById(id int) (*models.User, error)
	UpdateUserById(id int, request *user.UpdateUserRequest) (*models.User, error)
	DeleteUserById(id int) (*models.User, error)
}

type CategoryService interface {
	GetCategoryAll() ([]*catresponse.CategoryResponse, error)
	GetCategoryById(categoryId int) (*catresponse.CategoryResponse, error)
	GetCategoryBySlug(slug string) (*catresponse.CategoryResponse, error)
	CreateCategory(request *category.CreateCategoryRequest) (*catresponse.CategoryResponse, error)
	UpdateCategoryById(id int, updatedCategory *category.UpdateCategoryRequest) (*catresponse.CategoryResponse, error)
	DeleteCategoryById(categoryId int) (*catresponse.CategoryResponse, error)
}

type ProductService interface {
	GetProductAll() ([]*pro.ProductResponse, error)
	CreateProduct(request *product.CreateProductRequest) (*pro.ProductResponse, error)
	GetProductById(productId int) (*pro.ProductResponse, error)
	GetProductBySlug(slug string) (*pro.ProductResponse, error)
	UpdateProduct(productId int, request *product.UpdateProductRequest) (*pro.ProductResponse, error)
	DeleteProduct(productId int) (*pro.ProductResponse, error)
}

type CartService interface {
	GetCartAll(userId int) (*[]models.Cart, error)
	CreateCart(cartRequest *cart.CartCreateRequest) (*models.Cart, error)
	DeleteCart(cartId int) (*models.Cart, error)
	DeleteCartMany(cartIds cart.DeleteCartRequest) (int64, error)
}

type OrderService interface {
	GetOrderAll() (*[]or.OrderResponses, error)
	CreateOrder(userId int, request *order.CreateOrderRequest) (*models.Order, error)
	GetOrderById(orderId int) (*or.OrderResponse, error)
	GetOrdersByUser(userId int) (*[]models.Order, error)
}

type ReviewService interface {
	GetReviewAll() (*[]models.Review, error)
	GetReviewById(reviewId int) (*models.Review, error)
	CreateReview(productId, userId int, request *review.CreateReviewRequest) (*models.Review, error)
}

type SliderService interface {
	GetSliderAll() (*[]models.Slider, error)
	GetSliderById(sliderId int) (*models.Slider, error)
	CreateSlider(request slider.CreateSliderRequest) (*models.Slider, error)
	UpdateSliderById(sliderId int, request slider.UpdateSliderRequest) (*models.Slider, error)
	DeleteSliderById(sliderId int) (*models.Slider, error)
}

type RajaOngkirService interface {
	GetProvince() (map[string]interface{}, error)
	GetCity(provinceId int) (map[string]interface{}, error)
	GetCost(request rajaongkirrequest.OngkosRequest) (map[string]interface{}, error)
}

type MidtransService interface {
	CreateTransaction(request *midtransrequest.CreateMidtransRequest) (*snap.Response, error)
}
