package repository

import (
	"errors"

	"github.com/alifdwt/haiwan-go/internal/domain/requests/cart"
	"github.com/alifdwt/haiwan-go/internal/models"
	"gorm.io/gorm"
)

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *cartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) GetCartAll(userId int) (*[]models.Cart, error) {
	var carts []models.Cart

	db := r.db.Model(&carts)

	checkCartByUserId := db.Debug().First(&carts, "user_id", userId)
	if checkCartByUserId.RowsAffected < 1 {
		return nil, errors.New("cart not found")
	}

	return &carts, nil
}

func (r *cartRepository) CreateCart(cartRequest *cart.CartCreateRequest) (*models.Cart, error) {
	var cartModel models.Cart

	db := r.db.Model(&cartModel)

	cartModel.Name = cartRequest.Name
	cartModel.Price = cartRequest.Price
	cartModel.Image = cartRequest.ImageProduct
	cartModel.Quantity = cartRequest.Quantity
	cartModel.Weight = cartRequest.Weight
	cartModel.ProductID = uint(cartRequest.ProductID)
	cartModel.UserID = uint(cartRequest.UserID)

	addCart := db.Debug().Create(&cartModel).Commit()
	if addCart.RowsAffected < 1 {
		return nil, errors.New("unable to create cart")
	}

	return &cartModel, nil
}

func (r *cartRepository) DeleteCart(cartId int) (*models.Cart, error) {
	var cart models.Cart

	db := r.db.Model(cart)

	checkCartById := db.Debug().First(&cart, "id=?", cartId)
	if checkCartById.RowsAffected < 0 {
		return &cart, errors.New("cart not found")
	}

	deleteCart := db.Debug().Delete(&cart)
	if deleteCart.Error != nil {
		return nil, deleteCart.Error
	}

	return &cart, nil
}

func (r *cartRepository) DeleteCartMany(cartIds cart.DeleteCartRequest) (int64, error) {
	var cart models.Cart

	db := r.db.Model(cart)

	result := db.Debug().Where("id IN (?)", cartIds).Delete(&cart)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
