package repository

import (
	"errors"

	"github.com/alifdwt/haiwan-go/internal/domain/requests/slider"
	"github.com/alifdwt/haiwan-go/internal/models"
	"gorm.io/gorm"
)

type sliderRepository struct {
	db *gorm.DB
}

func NewSliderRepository(db *gorm.DB) *sliderRepository {
	return &sliderRepository{db: db}
}

func (r *sliderRepository) GetSliderAll() (*[]models.Slider, error) {
	var sliders []models.Slider

	db := r.db.Model(&sliders)

	checkSlider := db.Debug().Find(&sliders)
	if checkSlider.RowsAffected < 0 {
		return nil, errors.New("there is no slider yet")
	}

	return &sliders, nil
}

func (r *sliderRepository) GetSliderById(sliderId int) (*models.Slider, error) {
	var slider models.Slider

	db := r.db.Model(slider)

	checkSliderById := db.Debug().Where("id = ?", sliderId).First(&slider)
	if checkSliderById.RowsAffected < 0 {
		return nil, errors.New("failed to get slider by id")
	}

	return &slider, nil
}

func (r *sliderRepository) CreateSlider(sliderRequest *slider.CreateSliderRequest) (*models.Slider, error) {
	var slider models.Slider

	db := r.db.Model(&slider)

	mySlider := models.Slider{
		Name:  sliderRequest.Nama,
		Image: sliderRequest.FilePath,
	}

	addSlider := db.Debug().Create(&mySlider).Commit()
	if addSlider.RowsAffected < 1 {
		return &mySlider, errors.New("error while creating slider")
	}

	return &mySlider, nil
}

func (r *sliderRepository) UpdateSliderById(sliderId int, updatedSlider *slider.UpdateSliderRequest) (*models.Slider, error) {
	mySlider, err := r.GetSliderById(sliderId)
	if err != nil {
		return nil, err
	}

	if mySlider != nil {
		mySlider.Name = updatedSlider.Nama
		mySlider.Image = updatedSlider.FilePath

		if err := r.db.Debug().Save(mySlider).Error; err != nil {
			return nil, err
		}

		return mySlider, nil
	}
	return nil, errors.New("slider not found")
}

func (r *sliderRepository) DeleteSliderById(sliderId int) (*models.Slider, error) {
	var slider models.Slider

	mySlider, err := r.GetSliderById(sliderId)
	if err != nil {
		return nil, err
	}

	if mySlider == nil {
		return nil, errors.New("slider not found")
	}

	db := r.db.Model(&slider)

	if err := db.Debug().Delete(mySlider).Error; err != nil {
		return nil, err
	}

	return mySlider, nil
}
