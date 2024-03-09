package mapper

import (
	"github.com/alifdwt/haiwan-go/internal/domain/responses/category"
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
