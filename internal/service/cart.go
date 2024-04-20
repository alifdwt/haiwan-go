package service

import (
	"github.com/alifdwt/haiwan-go/internal/domain/requests/cart"
	cartRes "github.com/alifdwt/haiwan-go/internal/domain/responses/cart"
	"github.com/alifdwt/haiwan-go/internal/mapper"
	"github.com/alifdwt/haiwan-go/internal/repository"
	"github.com/alifdwt/haiwan-go/pkg/logger"
	"go.uber.org/zap"
)

type cartService struct {
	repository repository.CartRepository
	logger     logger.Logger
	mapper     mapper.CartMapping
}

func NewCartService(repository repository.CartRepository, logger logger.Logger, mapper mapper.CartMapping) *cartService {
	return &cartService{repository: repository, logger: logger, mapper: mapper}
}

func (s *cartService) GetCartAll(userId int) ([]*cartRes.CartResponse, error) {
	res, err := s.repository.GetCartAll(userId)
	if err != nil {
		s.logger.Error("Error while retrieving cart by user id", zap.Error(err))
		return nil, err
	}

	mapper := s.mapper.ToCartResponses(res)

	return mapper, nil
}

func (s *cartService) CreateCart(cartRequest *cart.CartCreateRequest) (*cartRes.CartResponse, error) {
	cart := &cart.CartCreateRequest{
		Name:         cartRequest.Name,
		Price:        cartRequest.Price,
		ImageProduct: cartRequest.ImageProduct,
		Quantity:     cartRequest.Quantity,
		Weight:       cartRequest.Weight,
		ProductID:    cartRequest.ProductID,
		UserID:       cartRequest.UserID,
	}

	res, err := s.repository.CreateCart(cart)
	if err != nil {
		s.logger.Error("Error while creating cart", zap.Error(err))
		return nil, err
	}

	mapper := s.mapper.ToCartResponse(res)

	return mapper, err
}

func (s *cartService) DeleteCart(cartId int) (*cartRes.CartResponse, error) {
	res, err := s.repository.DeleteCart(cartId)
	if err != nil {
		s.logger.Error("Error while deleting cart", zap.Error(err))
		return nil, err
	}

	mapper := s.mapper.ToCartResponse(res)

	return mapper, nil
}

func (s *cartService) DeleteCartMany(cartIds cart.DeleteCartRequest) (int64, error) {
	res, err := s.repository.DeleteCartMany(cartIds)
	if err != nil {
		s.logger.Error("Error while deleting many carts", zap.Error(err))
		return 0, err
	}

	return res, nil
}
