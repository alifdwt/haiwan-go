package mapper

import (
	"github.com/alifdwt/haiwan-go/internal/domain/responses/category"
	"github.com/alifdwt/haiwan-go/internal/models"
)

type categoryMapper struct {
}

func NewCategoryMapper() *categoryMapper {
	return &categoryMapper{}
}

func (m *categoryMapper) ToCategoryResponse(request *models.Category) *category.CategoryResponse {
	return &category.CategoryResponse{
		ID:          request.ID,
		Name:        request.Name,
		Description: request.Description,
		Slug:        request.SlugCategory,
		ImagePath:   request.ImageCategory,
		CreatedAt:   request.CreatedAt.String(),
		UpdatedAt:   request.UpdatedAt.String(),
	}
}

func (m *categoryMapper) ToCategoryWithRelationResponse(request *models.Category) *category.CategoryWithRelationResponse {
	return &category.CategoryWithRelationResponse{
		CategoryResponse: *m.ToCategoryResponse(request),
		Products:         NewMapper().ProductMapper.ToProductResponses(&request.Products),
	}
}

func (m *categoryMapper) ToCategoryResponses(requests *[]models.Category) []*category.CategoryResponse {
	var responses []*category.CategoryResponse
	for _, request := range *requests {
		responses = append(responses, m.ToCategoryResponse(&request))
	}

	return responses
}
