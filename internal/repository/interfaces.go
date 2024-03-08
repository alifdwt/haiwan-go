package repository

import (
	"github.com/alifdwt/haiwan-go/internal/domain/requests/auth"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/user"
	"github.com/alifdwt/haiwan-go/internal/models"
)

type UserRepository interface {
	CreateUser(request *auth.RegisterRequest) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserAll() (*[]models.User, error)
	GetUserById(id int) (*models.User, error)
	UpdateUserById(id int, updatedUser *user.UpdateUserRequest) (*models.User, error)
	DeleteUserById(id int) (*models.User, error)
	CountUser() (int, error)
}
