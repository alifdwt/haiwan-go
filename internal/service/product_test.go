package service

import (
	"fmt"
	"testing"

	proreq "github.com/alifdwt/haiwan-go/internal/domain/requests/product"
	pro "github.com/alifdwt/haiwan-go/internal/domain/responses/product"
	"github.com/alifdwt/haiwan-go/pkg/random"
	"github.com/stretchr/testify/require"
	"gopkg.in/loremipsum.v1"
)

func createRandomProduct(t *testing.T, categoryId uint) pro.ProductResponse {
	categoryIdStr := fmt.Sprint(categoryId)
	productName := random.RandomOwner()
	rating := random.RandomInt(3, 5)
	description := loremipsum.New().Paragraph()

	arg := &proreq.CreateProductRequest{
		Name:         productName,
		CategoryID:   categoryIdStr,
		Description:  description,
		Price:        random.RandomMoney(),
		CountInStock: random.RandomInt(1, 10),
		Weight:       random.RandomInt(500, 1000),
		Brand:        random.RandomString(6),
		Rating:       &rating,
		FilePath:     random.RandomProductPhotoUrl(),
	}

	product, err := testQueries.Product.CreateProduct(arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.CategoryID, fmt.Sprint(product.CategoryID))
	require.Equal(t, arg.Description, product.Description)
	require.Equal(t, arg.Price, product.Price)
	require.Equal(t, arg.CountInStock, product.CountInStock)
	require.Equal(t, arg.Weight, product.Weight)
	// require.Equal(t, arg.Rating, product.Rating)
	require.Equal(t, arg.FilePath, product.ImagePath)

	require.NotZero(t, product.CreatedAt)
	require.NotZero(t, product.UpdatedAt)

	return *product
}

func TestCreateProduct(t *testing.T) {
	category := createRandomCategory(t)
	createRandomProduct(t, category.ID)
}

func TestGetProductAll(t *testing.T) {
	category := createRandomCategory(t)
	for i := 0; i < 5; i++ {
		createRandomProduct(t, category.ID)
	}

	products, err := testQueries.Product.GetProductAll()
	require.NoError(t, err)
	require.NotEmpty(t, products)

	for _, product := range products {
		require.NotEmpty(t, product)
	}
}

func TestGetProductById(t *testing.T) {
	category := createRandomCategory(t)
	product1 := createRandomProduct(t, category.ID)

	product2, err := testQueries.Product.GetProductById(product1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1.Name, product2.Name)
	require.Equal(t, product1.CategoryID, product2.CategoryID)
	require.Equal(t, product1.Description, product2.Description)
	require.Equal(t, product1.Price, product2.Price)
	require.Equal(t, product1.CountInStock, product2.CountInStock)
	require.Equal(t, product1.Weight, product2.Weight)
	// require.Equal(t, product1.Rating, product2.Rating)
	require.Equal(t, product1.ImagePath, product2.ImagePath)

	require.NotZero(t, product2.CreatedAt)
	require.NotZero(t, product2.UpdatedAt)
}

func TestGetProductBySlug(t *testing.T) {
	category := createRandomCategory(t)
	product1 := createRandomProduct(t, category.ID)

	product2, err := testQueries.Product.GetProductBySlug(product1.Slug)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1.Name, product2.Name)
	require.Equal(t, product1.CategoryID, product2.CategoryID)
	require.Equal(t, product1.Description, product2.Description)
	require.Equal(t, product1.Price, product2.Price)
	require.Equal(t, product1.CountInStock, product2.CountInStock)
	require.Equal(t, product1.Weight, product2.Weight)
	// require.Equal(t, product1.Rating, product2.Rating)
	require.Equal(t, product1.ImagePath, product2.ImagePath)

	require.NotZero(t, product2.CreatedAt)
	require.NotZero(t, product2.UpdatedAt)
}

func TestUpdateProduct(t *testing.T) {
	category := createRandomCategory(t)
	product1 := createRandomProduct(t, category.ID)

	categoryIdStr := fmt.Sprint(product1.CategoryID)
	productName := random.RandomOwner()
	rating := int(random.RandomInt(3, 5))
	description := loremipsum.New().Paragraph()

	arg := &proreq.UpdateProductRequest{
		Name:         productName,
		CategoryID:   categoryIdStr,
		Description:  description,
		Price:        int(random.RandomMoney()),
		CountInStock: int(random.RandomInt(1, 10)),
		Weight:       int(random.RandomInt(500, 1000)),
		Brand:        random.RandomString(6),
		Rating:       rating,
		FilePath:     "https://picsum.photos/960/540",
	}

	product2, err := testQueries.Product.UpdateProduct(product1.ID, arg)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, arg.Name, product2.Name)
	require.Equal(t, arg.CategoryID, fmt.Sprint(product2.CategoryID))
	require.Equal(t, arg.Description, product2.Description)
	require.Equal(t, arg.Price, product2.Price)
	require.Equal(t, arg.CountInStock, product2.CountInStock)
	require.Equal(t, arg.Weight, product2.Weight)
	// require.Equal(t, arg.Rating, product2.Rating)
	require.Equal(t, arg.FilePath, product2.ImagePath)

	require.NotZero(t, product2.CreatedAt)
	require.NotZero(t, product2.UpdatedAt)
	require.NotEqual(t, product1.CreatedAt, product2.UpdatedAt)
}

func TestDeleteProduct(t *testing.T) {
	category := createRandomCategory(t)
	product1 := createRandomProduct(t, category.ID)

	_, err := testQueries.Product.DeleteProduct(product1.ID)
	require.NoError(t, err)
}
