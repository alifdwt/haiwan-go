package mapper

import (
	"github.com/alifdwt/haiwan-go/internal/domain/responses/product"
	"github.com/alifdwt/haiwan-go/internal/models"
)

type productMapper struct {
}

func NewProductMapper() *productMapper {
	return &productMapper{}
}

func (m *productMapper) ToProductResponse(request *models.Product) *product.ProductResponse {
	return &product.ProductResponse{
		ID:           int(request.ID),
		Name:         request.Name,
		Description:  request.Description,
		Slug:         request.SlugProduct,
		Brand:        request.Brand,
		Weight:       request.Weight,
		Rating:       int(request.Rating),
		ImagePath:    request.ImageProduct,
		Price:        request.Price,
		CountInStock: request.CountInStock,
		CategoryID:   int(request.CategoryID),
		CreatedAt:    request.CreatedAt.String(),
		UpdatedAt:    request.UpdatedAt.String(),
		Category:     *NewCategoryMapper().ToCategoryResponse(&request.Category),
	}
}

func (m *productMapper) ToProductResponses(requests *[]models.Product) []*product.ProductResponse {
	var responses []*product.ProductResponse
	for _, request := range *requests {
		responses = append(responses, m.ToProductResponse(&request))
	}

	return responses
}
