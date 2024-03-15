package service

import (
	"github.com/alifdwt/haiwan-go/internal/domain/requests/slider"
	"github.com/alifdwt/haiwan-go/internal/models"
	"github.com/alifdwt/haiwan-go/internal/repository"
	"github.com/alifdwt/haiwan-go/pkg/logger"
	"go.uber.org/zap"
)

type sliderService struct {
	repository repository.SliderRepository
	logger     logger.Logger
}

func NewSliderService(repository repository.SliderRepository, logger logger.Logger) *sliderService {
	return &sliderService{
		repository: repository,
		logger:     logger,
	}
}

func (s *sliderService) GetSliderAll() (*[]models.Slider, error) {
	res, err := s.repository.GetSliderAll()
	if err != nil {
		s.logger.Error("Error while getting all sliders", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *sliderService) GetSliderById(sliderId int) (*models.Slider, error) {
	res, err := s.repository.GetSliderById(sliderId)
	if err != nil {
		s.logger.Error("Error while getting slider by id: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *sliderService) CreateSlider(request slider.CreateSliderRequest) (*models.Slider, error) {
	schema := &slider.CreateSliderRequest{
		Nama:     request.Nama,
		FilePath: request.FilePath,
	}

	res, err := s.repository.CreateSlider(schema)
	if err != nil {
		s.logger.Error("Error while creating slider: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *sliderService) UpdateSliderById(sliderId int, request slider.UpdateSliderRequest) (*models.Slider, error) {
	schema := &slider.UpdateSliderRequest{
		Nama:     request.Nama,
		FilePath: request.FilePath,
	}

	res, err := s.repository.UpdateSliderById(sliderId, schema)
	if err != nil {
		s.logger.Error("Error while updating slider: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *sliderService) DeleteSliderById(sliderId int) (*models.Slider, error) {
	res, err := s.repository.DeleteSliderById(sliderId)
	if err != nil {
		s.logger.Error("Error while deleting slider: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}
