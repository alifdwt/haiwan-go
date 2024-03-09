package service

import (
	"errors"

	"github.com/alifdwt/haiwan-go/internal/domain/requests/cart"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/product"
	pro "github.com/alifdwt/haiwan-go/internal/domain/responses/product"
	"github.com/alifdwt/haiwan-go/internal/mapper"
	"github.com/alifdwt/haiwan-go/internal/repository"
	"github.com/alifdwt/haiwan-go/pkg/logger"
	"go.uber.org/zap"
)

type productService struct {
	repository repository.ProductRepository
	logger     logger.Logger
	mapper     mapper.ProductMapping
}

func NewProductService(repository repository.ProductRepository, logger logger.Logger, mapper mapper.ProductMapping) *productService {
	return &productService{
		repository: repository,
		logger:     logger,
		mapper:     mapper,
	}
}

func (s *productService) GetProductAll() ([]*pro.ProductResponse, error) {
	res, err := s.repository.GetProductAll()
	if err != nil {
		s.logger.Error("Error while getting all products", zap.Error(err))
		return nil, err
	}

	mapper := s.mapper.ToProductResponses(res)

	return mapper, nil
}

func (s *productService) CreateProduct(request *product.CreateProductRequest) (*pro.ProductResponse, error) {
	var schema product.CreateProductRequest

	schema.Name = request.Name
	schema.CategoryID = request.CategoryID
	schema.Description = request.Description
	schema.Price = request.Price
	schema.CountInStock = request.CountInStock
	schema.Weight = request.Weight
	schema.Brand = request.Brand
	schema.Rating = request.Rating
	schema.FilePath = request.FilePath

	res, err := s.repository.CreateProduct(&schema)
	if err != nil {
		s.logger.Error("Error while creating product:", zap.Error(err))
		return nil, err
	}

	mapper := s.mapper.ToProductResponse(res)

	return mapper, nil
}

func (s *productService) GetProductById(productId int) (*pro.ProductResponse, error) {
	res, err := s.repository.GetProductById(productId)
	if err != nil {
		s.logger.Error("Error while getting product by id", zap.Error(err))
		return nil, err
	}

	mapper := s.mapper.ToProductResponse(res)

	return mapper, nil
}

func (s *productService) GetProductBySlug(slug string) (*pro.ProductResponse, error) {
	res, err := s.repository.GetProductBySlug(slug)
	if err != nil {
		s.logger.Error("Error while getting product by slug", zap.Error(err))
		return nil, err
	}

	mapper := s.mapper.ToProductResponse(res)

	return mapper, nil
}

func (s *productService) UpdateProduct(productId int, request *product.UpdateProductRequest) (*pro.ProductResponse, error) {
	var schema product.UpdateProductRequest

	schema.Name = request.Name
	schema.CategoryID = request.CategoryID
	schema.Description = request.Description
	schema.Price = request.Price
	schema.CountInStock = request.CountInStock
	schema.Weight = request.Weight
	schema.Rating = request.Rating
	schema.FilePath = request.FilePath

	res, err := s.repository.UpdateProduct(productId, request)
	if err != nil {
		s.logger.Error("Error while updating product", zap.Error(err))
		return nil, err
	}

	mapper := s.mapper.ToProductResponse(res)

	return mapper, nil
}

func (s *productService) DeleteProduct(productId int) (*pro.ProductResponse, error) {
	res, err := s.repository.DeleteProduct(productId)
	if err != nil {
		s.logger.Error("Error while deleting product", zap.Error(err))
		return nil, err
	}

	mapper := s.mapper.ToProductResponse(res)

	return mapper, nil
}

func (s *productService) UpdateQuantity(cart []*cart.CartCreateRequest) (bool, error) {
	if len(cart) == 0 {
		return false, errors.New("no cart data received")
	}

	for _, item := range cart {
		productId := item.ProductID
		quantity := item.Quantity

		product, err := s.GetProductById(productId)
		if err != nil {
			s.logger.Error("Error while getting product by id", zap.Error(err))
			return false, err
		}

		currentStock := product.CountInStock

		newStock := currentStock - quantity

		result, err := s.repository.MyUpdateQuantity(productId, newStock)
		if err != nil {
			return false, err
		}
		if !result {
			return false, nil
		}
	}

	return true, nil
}
