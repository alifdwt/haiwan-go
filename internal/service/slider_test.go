package service

import (
	"testing"
	"time"

	"github.com/alifdwt/haiwan-go/internal/domain/requests/slider"
	"github.com/alifdwt/haiwan-go/internal/models"
	"github.com/alifdwt/haiwan-go/pkg/random"
	"github.com/stretchr/testify/require"
)

func createRandomSlider(t *testing.T) models.Slider {
	arg := &slider.CreateSliderRequest{
		Nama:     random.RandomString(6),
		FilePath: random.RandomSliderPhotoUrl(),
	}

	slider, err := testQueries.Slider.CreateSlider(*arg)
	require.NoError(t, err)
	require.NotEmpty(t, slider)

	require.Equal(t, arg.Nama, slider.Name)
	require.Equal(t, arg.FilePath, slider.Image)

	require.NotZero(t, slider.ID)
	require.NotZero(t, slider.CreatedAt)

	return *slider
}

func TestCreateSlider(t *testing.T) {
	createRandomSlider(t)
}

func TestGetAllSlider(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomSlider(t)
	}

	sliders, err := testQueries.Slider.GetSliderAll()
	require.NoError(t, err)
	require.NotEmpty(t, sliders)

	for _, slider := range *sliders {
		require.NotEmpty(t, slider)
	}
}

func TestGetSliderById(t *testing.T) {
	slider1 := createRandomSlider(t)
	slider2, err := testQueries.Slider.GetSliderById(int(slider1.ID))
	require.NoError(t, err)
	require.NotEmpty(t, slider2)

	require.Equal(t, slider1.ID, slider2.ID)
	require.Equal(t, slider1.Name, slider2.Name)
	require.Equal(t, slider1.Image, slider2.Image)

	require.WithinDuration(t, slider1.CreatedAt, slider2.CreatedAt, time.Second)
	require.WithinDuration(t, slider1.UpdatedAt, slider2.UpdatedAt, time.Second)
}

func TestUpdateSliderById(t *testing.T) {
	slider1 := createRandomSlider(t)
	arg := &slider.UpdateSliderRequest{
		Nama:     random.RandomString(6),
		FilePath: random.RandomSliderPhotoUrl(),
	}

	slider2, err := testQueries.Slider.UpdateSliderById(int(slider1.ID), *arg)
	require.NoError(t, err)
	require.NotEmpty(t, slider2)

	require.Equal(t, slider1.ID, slider2.ID)
	require.NotEqual(t, slider1.Name, slider2.Name)
	require.NotEqual(t, slider1.Image, slider2.Image)

	// require.Equal(t, slider1.CreatedAt, slider2.CreatedAt)
	require.NotEqual(t, slider1.UpdatedAt, slider2.UpdatedAt)
}

func TestDeleteSliderById(t *testing.T) {
	slider1 := createRandomSlider(t)
	_, err := testQueries.Slider.DeleteSliderById(int(slider1.ID))
	require.NoError(t, err)
}
