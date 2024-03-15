package service

import (
	"github.com/alifdwt/haiwan-go/internal/domain/requests/review"
	"github.com/alifdwt/haiwan-go/internal/models"
	"github.com/alifdwt/haiwan-go/internal/repository"
)

type reviewService struct {
	reviewRepository repository.ReviewRepository
}

func NewReviewService(reviewRepository repository.ReviewRepository) *reviewService {
	return &reviewService{
		reviewRepository: reviewRepository,
	}
}

func (s *reviewService) GetReviewAll() (*[]models.Review, error) {
	res, err := s.reviewRepository.GetReviewAll()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *reviewService) GetReviewById(reviewId int) (*models.Review, error) {
	res, err := s.reviewRepository.GetReviewById(reviewId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *reviewService) CreateReview(productId, userId int, request *review.CreateReviewRequest) (*models.Review, error) {
	res, err := s.reviewRepository.CreateReview(*request, userId, productId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
