package repository

import (
	"github.com/alifdwt/haiwan-go/internal/domain/requests/auth"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/cart"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/category"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/order"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/product"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/review"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/slider"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/user"
	"github.com/alifdwt/haiwan-go/internal/models"
)

type UserRepository interface {
	CreateUser(request *auth.RegisterRequest) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserAll() (*[]models.User, error)
	GetUserById(id int) (*models.User, error)
	UpdateUserById(id int, updatedUser *user.UpdateUserRequest) (*models.User, error)
	DeleteUserById(id int) (*models.User, error)
	CountUser() (int, error)
}

type CategoryRepository interface {
	GetCategoryAll() (*[]models.Category, error)
	CreateCategory(request *category.CreateCategoryRequest) (*models.Category, error)
	GetCategoryById(categoryId int) (*models.Category, error)
	GetCategoryBySlug(slug string) (*models.Category, error)
	UpdateCategoryById(id int, updatedCategory *category.UpdateCategoryRequest) (*models.Category, error)
	DeleteCategoryById(categoryId int) (*models.Category, error)
}

type ProductRepository interface {
	GetProductAll() (*[]models.Product, error)
	GetProductBySlug(slug string) (*models.Product, error)
	CreateProduct(request *product.CreateProductRequest) (*models.Product, error)
	GetProductById(productId int) (*models.Product, error)
	MyUpdateQuantity(productId int, quantity int) (bool, error)
	UpdateProduct(productId int, request *product.UpdateProductRequest) (*models.Product, error)
	DeleteProduct(productId int) (*models.Product, error)
	CountProduct() (int, error)
}

type CartRepository interface {
	GetCartAll(userId int) (*[]models.Cart, error)
	CreateCart(cartRequest *cart.CartCreateRequest) (*models.Cart, error)
	DeleteCart(cartId int) (*models.Cart, error)
	DeleteCartMany(cartIds cart.DeleteCartRequest) (int64, error)
}

type OrderRepository interface {
	GetOrderAll() (*[]models.Order, error)
	CreateOrder(user_id int, request *order.CreateOrderRequest) (*models.Order, error)
	GetOrderById(orderId int) (*models.Order, error)
	GetOrderByUserId(userId int) (*[]models.Order, error)
	CountOrder() (int, error)
	SumTotalPrice() (int, error)
	CalculateYearlyRevenue() ([]int, error)
}

type ReviewRepository interface {
	GetReviewAll() (*[]models.Review, error)
	GetReviewById(reviewId int) (*models.Review, error)
	CreateReview(request review.CreateReviewRequest, userId int, productId int) (*models.Review, error)
}

type SliderRepository interface {
	GetSliderAll() (*[]models.Slider, error)
	GetSliderById(sliderId int) (*models.Slider, error)
	CreateSlider(sliderRequest *slider.CreateSliderRequest) (*models.Slider, error)
	UpdateSliderById(sliderId int, updatedSlider *slider.UpdateSliderRequest) (*models.Slider, error)
	DeleteSliderById(sliderId int) (*models.Slider, error)
}
