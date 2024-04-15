package repository

import (
	"errors"
	"strconv"

	"github.com/alifdwt/haiwan-go/internal/domain/requests/product"
	"github.com/alifdwt/haiwan-go/internal/models"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetProductAll() (*[]models.Product, error) {
	var products []models.Product

	db := r.db.Model(&products)

	checkProduct := db.Debug().Preload("Category").Find(&products)
	if checkProduct.RowsAffected < 1 {
		return nil, errors.New("there is no products yet")
	}

	return &products, nil
}

func (r *productRepository) GetProductBySlug(slug string) (*models.Product, error) {
	var product models.Product

	db := r.db.Model(&product)

	checkProductBySlug := db.Preload("Review").Debug().Where("slug_product = ?", slug).First(&product)
	if checkProductBySlug.RowsAffected < 0 {
		return nil, errors.New("failed to get product by slug")
	}

	return &product, nil
}

func (r *productRepository) CreateProduct(request *product.CreateProductRequest) (*models.Product, error) {
	var product models.Product

	slugProduct := slug.Make(request.Name)
	if request.CategoryID == "" {
		return nil, errors.New("category id is empty")
	}

	categoryId, err := strconv.Atoi(request.CategoryID)
	if err != nil {
		return nil, errors.New("failed to convert category id to int: " + err.Error())
	}

	ratingFloat := float64(*request.Rating)
	product.Name = request.Name
	product.Description = request.Description
	product.SlugProduct = slugProduct
	product.ImageProduct = request.FilePath
	product.Price = request.Price
	product.Weight = request.Weight
	product.Brand = request.Brand
	product.CategoryID = uint(categoryId)
	product.CountInStock = request.CountInStock
	product.Rating = ratingFloat

	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, errors.New("failed to create product: " + tx.Error.Error())
	}

	if err := tx.Create(&product).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("failed to create product: " + err.Error())
	}

	tx.Commit()
	return &product, nil
}

func (r *productRepository) GetProductById(productId int) (*models.Product, error) {
	var product models.Product

	db := r.db.Model(&product)

	checkProductById := db.Debug().Where("id", productId).First(&product)
	if checkProductById.RowsAffected < 0 {
		return nil, errors.New("failed to get id")
	}

	return &product, nil
}

func (r *productRepository) MyUpdateQuantity(productId int, quantity int) (bool, error) {
	var dbProduct models.Product
	if err := r.db.Where("id = ?", productId).First(&dbProduct).Error; err != nil {
		return false, err
	}

	if dbProduct.ID != 0 {
		dbProduct.CountInStock = quantity

		if err := r.db.Save(&dbProduct).Error; err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil
}

func (r *productRepository) UpdateProduct(productId int, request *product.UpdateProductRequest) (*models.Product, error) {
	var product models.Product

	slugProduct := slug.Make(request.Name)

	db := r.db.Model(&product)

	if err := db.Debug().Where("id = ?", productId).First(&product).Error; err != nil {
		return nil, errors.New("failed to find product: " + err.Error())
	}

	categoryId, err := strconv.Atoi(request.CategoryID)
	if err != nil {
		return nil, errors.New("failed to convert category id to int: " + err.Error())
	}

	ratingStr := strconv.Itoa(request.Rating)
	ratingFloat, err := strconv.ParseFloat(ratingStr, 64)
	if err != nil {
		return nil, errors.New("failed to convert rating to float64: " + err.Error())
	}

	product.Name = request.Name
	product.SlugProduct = slugProduct
	product.ImageProduct = request.FilePath
	product.Description = request.Description
	product.Price = request.Price
	product.Brand = request.Brand
	product.Weight = request.Weight
	product.CategoryID = uint(categoryId)
	product.Brand = request.Brand
	product.CountInStock = request.CountInStock
	product.Rating = ratingFloat

	if err := db.Debug().Save(&product).Error; err != nil {
		return nil, errors.New("failed to update product: " + err.Error())
	}

	return &product, nil
}

func (r *productRepository) DeleteProduct(productId int) (*models.Product, error) {
	var product models.Product

	dbProduct, err := r.GetProductById(productId)
	if err != nil {
		return nil, err
	}

	if dbProduct == nil {
		return nil, errors.New("product not found")
	}

	db := r.db.Model(&product)
	if err := db.Debug().Delete(dbProduct).Error; err != nil {
		return nil, err
	}

	return dbProduct, nil
}

func (r *productRepository) CountProduct() (int, error) {
	var product models.Product

	db := r.db.Model(&product)

	var totalProduct int64

	db.Debug().Model(&product).Count(&totalProduct)

	return int(totalProduct), nil
}
