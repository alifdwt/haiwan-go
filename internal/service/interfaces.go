package service

import (
	"github.com/alifdwt/haiwan-go/internal/domain/requests/auth"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/user"
	"github.com/alifdwt/haiwan-go/internal/domain/responses"
	"github.com/alifdwt/haiwan-go/internal/models"
)

type AuthService interface {
	Register(input *auth.RegisterRequest) (*models.User, error)
	Login(input *auth.LoginRequest) (*responses.Token, error)
	RefreshToken(req auth.RefreshTokenRequest) (*responses.Token, error)
}

type UserService interface {
	GetUserAll() (*[]models.User, error)
	CreateUser(request *auth.RegisterRequest) (*models.User, error)
	GetUserById(id int) (*models.User, error)
	UpdateUserById(id int, request *user.UpdateUserRequest) (*models.User, error)
	DeleteUserById(id int) (*models.User, error)
}
