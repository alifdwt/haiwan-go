package mapper

import (
	"github.com/alifdwt/haiwan-go/internal/domain/responses/cart"
	"github.com/alifdwt/haiwan-go/internal/models"
)

type cartMapper struct {
}

func NewCartMapper() *cartMapper {
	return &cartMapper{}
}

func (m *cartMapper) ToCartResponse(request *models.Cart) *cart.CartResponse {
	return &cart.CartResponse{
		ID:        request.ID,
		Name:      request.Name,
		Price:     request.Price,
		Image:     request.Image,
		Quantity:  request.Quantity,
		Weight:    request.Weight,
		UserID:    request.UserID,
		ProductID: request.ProductID,
		CreatedAt: request.CreatedAt,
		UpdatedAt: request.UpdatedAt,
	}
}

func (m *cartMapper) ToCartResponses(requests *[]models.Cart) []*cart.CartResponse {
	var responses []*cart.CartResponse
	for _, request := range *requests {
		responses = append(responses, m.ToCartResponse(&request))
	}

	return responses
}
