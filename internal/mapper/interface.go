package mapper

import (
	"github.com/alifdwt/haiwan-go/internal/domain/responses/category"
	"github.com/alifdwt/haiwan-go/internal/domain/responses/order"
	"github.com/alifdwt/haiwan-go/internal/domain/responses/product"
	"github.com/alifdwt/haiwan-go/internal/models"
)

type CategoryMapping interface {
	ToCategoryResponse(request *models.Category) *category.CategoryResponse
	ToCategoryResponses(requests *[]models.Category) []*category.CategoryResponse
}

type ProductMapping interface {
	ToProductResponse(request *models.Product) *product.ProductResponse
	ToProductResponses(requests *[]models.Product) []*product.ProductResponse
}

type OrderMapping interface {
	ToOrderResponse(requests *models.Order) *order.OrderResponse
	ToOrderResponses(request *[]models.Order) []order.OrderResponses
}
