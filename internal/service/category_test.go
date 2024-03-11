package service

import (
	"testing"

	"github.com/alifdwt/haiwan-go/internal/domain/requests/category"
	catresponse "github.com/alifdwt/haiwan-go/internal/domain/responses/category"
	"github.com/alifdwt/haiwan-go/pkg/random"
	"github.com/gosimple/slug"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) catresponse.CategoryResponse {
	arg := &category.CreateCategoryRequest{
		Name:     random.RandomOwner(),
		FilePath: "https://picsum.photos/200",
	}

	category, err := testQueries.Category.CreateCategory(arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.Name, category.Name)
	require.Equal(t, arg.FilePath, category.ImagePath)
	require.Equal(t, slug.Make(arg.Name), category.Slug)

	require.NotZero(t, category.ID)
	require.NotZero(t, category.CreatedAt)
	require.NotZero(t, category.Description)

	return *category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategoryAll(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomCategory(t)
	}

	categories, err := testQueries.Category.GetCategoryAll()
	require.NoError(t, err)
	require.NotEmpty(t, categories)

	for _, category := range categories {
		require.NotEmpty(t, category)
	}
}

func TestGetCategoryById(t *testing.T) {
	category1 := createRandomCategory(t)
	category2, err := testQueries.Category.GetCategoryById(int(category1.ID))
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Name, category2.Name)
	require.Equal(t, category1.Description, category2.Description)
	require.Equal(t, category1.ImagePath, category2.ImagePath)
	require.Equal(t, category1.Slug, category2.Slug)

	require.Equal(t, category1.CreatedAt[:20], category2.CreatedAt[:20])
	require.Equal(t, category1.UpdatedAt[:20], category2.UpdatedAt[:20])

	// notExistCategory, err := testQueries.Category.GetCategoryById(1234567)
	// require.Error(t, err)
	// require.Empty(t, notExistCategory)
}

func TestGetCategoryBySlug(t *testing.T) {
	category1 := createRandomCategory(t)
	category2, err := testQueries.Category.GetCategoryBySlug(category1.Slug)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Name, category2.Name)
	require.Equal(t, category1.Description, category2.Description)
	require.Equal(t, category1.ImagePath, category2.ImagePath)
	require.Equal(t, category1.Slug, category2.Slug)

	require.Equal(t, category1.CreatedAt[:20], category2.CreatedAt[:20])
	require.Equal(t, category1.UpdatedAt[:20], category2.UpdatedAt[:20])

	// notExistCategory, err := testQueries.Category.GetCategoryBySlug("not-exist")
	// require.Error(t, err)
	// require.Empty(t, notExistCategory)
}

func TestUpdateCategoryById(t *testing.T) {
	category1 := createRandomCategory(t)
	arg := &category.UpdateCategoryRequest{
		Name:     random.RandomOwner(),
		FilePath: "https://picsum.photos/300",
	}

	category2, err := testQueries.Category.UpdateCategoryById(int(category1.ID), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Description, category2.Description)

	require.NotEqual(t, category1.Name, category2.Name)
	require.NotEqual(t, category1.ImagePath, category2.ImagePath)
	require.NotEqual(t, category1.Slug, category2.Slug)

	require.Equal(t, category1.CreatedAt[:20], category2.CreatedAt[:20])
	require.NotEqual(t, category1.UpdatedAt, category2.UpdatedAt)
}

func TestDeleteCategoryById(t *testing.T) {
	category1 := createRandomCategory(t)

	_, err := testQueries.Category.DeleteCategoryById(int(category1.ID))
	require.NoError(t, err)
}
