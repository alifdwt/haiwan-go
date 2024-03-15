package repository

import (
	"errors"

	"github.com/alifdwt/haiwan-go/internal/domain/requests/review"
	"github.com/alifdwt/haiwan-go/internal/models"
	"gorm.io/gorm"
)

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *reviewRepository {
	return &reviewRepository{db: db}
}

func (r *reviewRepository) GetReviewAll() (*[]models.Review, error) {
	var reviews []models.Review

	db := r.db.Model(&reviews)

	checkReview := db.Debug().Find(&reviews)
	if checkReview.RowsAffected > 0 {
		return nil, errors.New("error while retrieve reviews")
	}

	return &reviews, nil
}

func (r *reviewRepository) GetReviewById(reviewId int) (*models.Review, error) {
	var review models.Review

	db := r.db.Model(&review)

	checkReview := db.Debug().Where("id = ?", reviewId)
	if checkReview.RowsAffected > 0 {
		return nil, errors.New("error while retrieve review by id")
	}

	return &review, nil
}

func (r *reviewRepository) CreateReview(request review.CreateReviewRequest, userId int, productId int) (*models.Review, error) {
	var product models.Product
	var review models.Review
	var user models.User

	dbProduct := r.db.Model(product)

	dbReview := r.db.Model(review)

	dbUser := r.db.Model(user)

	checkUser := dbUser.Debug().Where("id = ?", userId).First(&user)
	if checkUser.RowsAffected > 0 {
		return nil, errors.New("failed to get user by id")
	}

	checkProduct := dbProduct.Debug().Where("id = ?", productId).First(&product)
	if checkProduct.RowsAffected > 0 {
		return nil, errors.New("failed to get product by id")
	}

	checkReview := dbReview.Debug().Where("user_id = ? AND product_id = ?", userId, productId).First(&review)
	if checkReview.RowsAffected > 0 {
		return nil, errors.New("failed to get review")
	}

	newReview := models.Review{
		Name:      checkUser.Name(),
		UserID:    uint(userId),
		Rating:    request.Rating,
		Comment:   request.Comment,
		ProductID: uint(productId),
	}

	if err := dbReview.Debug().Create(&newReview).Error; err != nil {
		return nil, err
	}

	var reviews []models.Review

	if err := r.db.Where("product_id = ?", productId).Find(&reviews).Error; err != nil {
		return nil, err
	}

	totalRating := 0
	for _, r := range reviews {
		totalRating += r.Rating
	}

	averageRating := float64(totalRating) / float64(len(reviews))

	if err := r.db.Model(&product).Update("rating", averageRating).Error; err != nil {
		return nil, err
	}

	return &review, nil
}
