package service

import (
	"github.com/alifdwt/haiwan-go/internal/domain/requests/order"
	or "github.com/alifdwt/haiwan-go/internal/domain/responses/order"
	"github.com/alifdwt/haiwan-go/internal/mapper"
	"github.com/alifdwt/haiwan-go/internal/models"
	"github.com/alifdwt/haiwan-go/internal/repository"
	"github.com/alifdwt/haiwan-go/pkg/logger"
	"go.uber.org/zap"
)

type orderService struct {
	mapper     mapper.OrderMapping
	repository repository.OrderRepository
	logger     logger.Logger
}

func NewOrderService(orderRepository repository.OrderRepository, logger logger.Logger, mapper mapper.OrderMapping) *orderService {
	return &orderService{
		repository: orderRepository,
		logger:     logger,
		mapper:     mapper,
	}
}

func (s *orderService) GetOrderAll() (*[]or.OrderResponses, error) {
	res, err := s.repository.GetOrderAll()
	if err != nil {
		s.logger.Error("Error while getting all orders", zap.Error(err))
		return nil, err
	}

	mapper := s.mapper.ToOrderResponses(res)

	return &mapper, nil
}

func (s *orderService) CreateOrder(userId int, request *order.CreateOrderRequest) (*models.Order, error) {
	res, err := s.repository.CreateOrder(userId, request)
	if err != nil {
		s.logger.Error("Error while creating order", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *orderService) GetOrderById(orderId int) (*or.OrderResponse, error) {
	res, err := s.repository.GetOrderById(orderId)
	if err != nil {
		s.logger.Error("error while retrieve order by id", zap.Error(err))
		return nil, err
	}

	mapper := s.mapper.ToOrderResponse(res)

	return mapper, nil
}

func (s *orderService) GetOrdersByUser(userId int) (*[]models.Order, error) {
	res, err := s.repository.GetOrderByUserId(userId)
	if err != nil {
		s.logger.Error("Error while retrieving order by user Id", zap.Error(err))
		return nil, err
	}

	return res, nil
}
