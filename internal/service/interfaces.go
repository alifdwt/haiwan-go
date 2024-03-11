package service

import (
	"github.com/alifdwt/haiwan-go/internal/domain/requests/auth"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/cart"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/category"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/order"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/product"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/user"
	"github.com/alifdwt/haiwan-go/internal/domain/responses"
	catresponse "github.com/alifdwt/haiwan-go/internal/domain/responses/category"
	or "github.com/alifdwt/haiwan-go/internal/domain/responses/order"
	pro "github.com/alifdwt/haiwan-go/internal/domain/responses/product"
	"github.com/alifdwt/haiwan-go/internal/models"
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
