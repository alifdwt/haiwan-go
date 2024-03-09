package service

import (
	"github.com/alifdwt/haiwan-go/internal/domain/requests/category"
	catresponse "github.com/alifdwt/haiwan-go/internal/domain/responses/category"
	"github.com/alifdwt/haiwan-go/internal/mapper"
	"github.com/alifdwt/haiwan-go/internal/repository"
)

type categoryService struct {
	mapper     mapper.CategoryMapping
	repository repository.CategoryRepository
}

func NewCategoryService(repository repository.CategoryRepository, mapper mapper.CategoryMapping) *categoryService {
	return &categoryService{
		repository: repository,
		mapper:     mapper,
	}
}

func (s *categoryService) GetCategoryAll() ([]*catresponse.CategoryResponse, error) {
	res, err := s.repository.GetCategoryAll()
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCategoryResponses(res)

	return mapper, nil
}

func (s *categoryService) GetCategoryById(categoryId int) (*catresponse.CategoryResponse, error) {
	res, err := s.repository.GetCategoryById(categoryId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCategoryResponse(res)

	return mapper, nil
}

func (s *categoryService) GetCategoryBySlug(slug string) (*catresponse.CategoryResponse, error) {
	res, err := s.repository.GetCategoryBySlug(slug)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCategoryResponse(res)

	return mapper, nil
}

func (s *categoryService) CreateCategory(request *category.CreateCategoryRequest) (*catresponse.CategoryResponse, error) {
	res, err := s.repository.CreateCategory(request)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCategoryResponse(res)

	return mapper, nil
}

func (s *categoryService) UpdateCategoryById(id int, updatedCategory *category.UpdateCategoryRequest) (*catresponse.CategoryResponse, error) {
	category := category.UpdateCategoryRequest{
		Name:     updatedCategory.Name,
		FilePath: updatedCategory.FilePath,
	}

	res, err := s.repository.UpdateCategoryById(id, &category)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCategoryResponse(res)

	return mapper, nil
}

func (s *categoryService) DeleteCategoryById(categoryId int) (*catresponse.CategoryResponse, error) {
	res, err := s.repository.DeleteCategoryById(categoryId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCategoryResponse(res)

	return mapper, nil
}
