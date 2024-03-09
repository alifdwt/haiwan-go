package repository

import (
	"testing"

	"github.com/alifdwt/haiwan-go/internal/domain/requests/category"
	"github.com/alifdwt/haiwan-go/internal/models"
	"github.com/alifdwt/haiwan-go/pkg/random"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) *models.Category {
	categoryName := random.RandomString(6)
	arg := &category.CreateCategoryRequest{
		Name:     categoryName,
		FilePath: "https://picsum.photos/200",
	}

	categoryRepository := NewCategoryRepository(setDatabase())
	category, err := CategoryRepository.CreateCategory(categoryRepository, arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.Name, category.Name)
	require.Equal(t, arg.FilePath, category.ImageCategory)

	require.NotZero(t, category.ID)
	require.NotZero(t, category.CreatedAt)
	require.Zero(t, category.DeletedAt)
	require.NotZero(t, category.Description)

	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}
