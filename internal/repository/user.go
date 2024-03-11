package repository

import (
	"errors"

	"github.com/alifdwt/haiwan-go/internal/domain/requests/auth"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/user"
	"github.com/alifdwt/haiwan-go/internal/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(registerReq *auth.RegisterRequest) (*models.User, error) {
	var user models.User

	user.Name = registerReq.Name
	user.Email = registerReq.Email
	user.Password = registerReq.Password

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().Where("email = ?", registerReq.Email)
	if checkEmailExist.RowsAffected > 1 {
		return &user, errors.New("error")
	}

	result := db.Debug().Create(&user).Commit()

	if result.RowsAffected < 1 {
		return nil, errors.New("error on create")
	}

	return &user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	db := r.db.Model(user)

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("email not found")
		}
		return nil, errors.New("error fetching email: " + err.Error())
	}

	return &user, nil
}

func (r *userRepository) GetUserAll() (*[]models.User, error) {
	var user []models.User

	db := r.db.Model(user)

	checkUser := db.Debug().Find(&user)
	if checkUser.RowsAffected < 1 {
		return nil, errors.New("failed to retrieve users")
	}

	return &user, nil
}

func (r *userRepository) GetUserById(id int) (*models.User, error) {
	var user models.User

	db := r.db.Model(user)

	checkUserById := db.Debug().Where("id = ?", id).First(&user)
	if checkUserById.RowsAffected < 0 {
		return &user, errors.New("user not found")
	}

	return &user, nil
}

func (r *userRepository) UpdateUserById(id int, updatedUser *user.UpdateUserRequest) (*models.User, error) {
	var user models.User

	db := r.db.Model(user)

	res, err := r.GetUserById(id)
	if err != nil {
		return nil, err
	}

	res.Name = updatedUser.Name
	res.Email = updatedUser.Email
	res.Password = updatedUser.Password

	// if err := db.Save(&res).Error; err != nil {
	// 	return nil, errors.New("error updating user: " + err.Error())
	// }

	updateUser := db.Debug().Updates(&res)
	if updateUser.RowsAffected > 1 {
		return res, errors.New("error updating user")
	}

	return res, nil
}

func (r *userRepository) DeleteUserById(id int) (*models.User, error) {
	var user models.User

	db := r.db.Model(user)

	res, err := r.GetUserById(id)
	if err != nil {
		return nil, err
	}

	if err := db.Delete(&res).Error; err != nil {
		return nil, errors.New("error on deleting user: " + err.Error())
	}

	return res, nil
}

func (r *userRepository) CountUser() (int, error) {
	var user models.User

	db := r.db.Model(user)

	var totalUser int64

	db.Debug().Model(&user).Count(&totalUser)

	return int(totalUser), nil
}
