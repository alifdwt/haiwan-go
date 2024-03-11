package repository

import (
	"errors"

	"github.com/alifdwt/haiwan-go/internal/domain/requests/category"
	"github.com/alifdwt/haiwan-go/internal/models"
	"github.com/gosimple/slug"
	"gopkg.in/loremipsum.v1"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetCategoryAll() (*[]models.Category, error) {
	var categories []models.Category

	db := r.db.Model(&categories)

	checkCategory := db.Debug().Find(&categories)
	if checkCategory.RowsAffected < 1 {
		return nil, errors.New("there are no categories yet")
	}

	return &categories, nil
}

func (r *categoryRepository) GetCategoryById(categoryId int) (*models.Category, error) {
	var category models.Category

	db := r.db.Model(category)

	checkCategoryById := db.Debug().Where("id = ?", categoryId).First(&category)
	if checkCategoryById.RowsAffected < 0 {
		return &category, errors.New("category id is not found")
	}

	return &category, nil
}

func (r *categoryRepository) GetCategoryBySlug(slug string) (*models.Category, error) {
	var category models.Category

	db := r.db.Model(&category)

	checkCategorySlug := db.Debug().Where("slug_category = ?", slug).First(&category)
	if checkCategorySlug.RowsAffected < 0 {
		return &category, errors.New("error while check category by slug")
	}

	return &category, nil
}

func (r *categoryRepository) CreateCategory(request *category.CreateCategoryRequest) (*models.Category, error) {
	var myCategory models.Category

	slugCategory := slug.Make(request.Name)
	loremDescription := loremipsum.New().Sentence()

	myCategory.Name = request.Name
	myCategory.SlugCategory = slugCategory
	myCategory.Description = loremDescription
	myCategory.ImageCategory = request.FilePath

	db := r.db.Model(&myCategory)

	checkCategoryName := db.Debug().Where("name = ?", request.Name).First(&myCategory)
	if checkCategoryName.RowsAffected > 0 {
		return &myCategory, errors.New("category already exists")
	}

	addCategory := db.Debug().Create(&myCategory).Commit()
	if addCategory.RowsAffected < 1 {
		return &myCategory, errors.New("error while creating category")
	}

	return &myCategory, nil
}

func (r *categoryRepository) UpdateCategoryById(id int, updatedCategory *category.UpdateCategoryRequest) (*models.Category, error) {
	var category models.Category

	slugCategory := slug.Make(updatedCategory.Name)
	loremDescription := loremipsum.New().Sentence()

	db := r.db.Model(&category)

	checkCategoryById := db.Debug().Where("id = ?", id).First(&category)
	if checkCategoryById.RowsAffected > 1 {
		return &category, errors.New("category not found")
	}

	category.Name = updatedCategory.Name
	category.SlugCategory = slugCategory
	category.Description = loremDescription
	category.ImageCategory = updatedCategory.FilePath

	updateCategory := db.Debug().Updates(&category)
	if updateCategory.RowsAffected > 1 {
		return &category, errors.New("error while updating category")
	}

	return &category, nil
}

func (r *categoryRepository) DeleteCategoryById(categoryId int) (*models.Category, error) {
	var category models.Category

	db := r.db.Model(&category)

	checkCategoryById := db.Debug().Where("id = ?", categoryId).First(&category)
	if checkCategoryById.RowsAffected > 1 {
		return &category, errors.New("category not found")
	}

	deleteCategory := db.Debug().Delete(&category)
	if deleteCategory.RowsAffected > 1 {
		return &category, errors.New("failed to delete category")
	}

	return &category, nil
}
