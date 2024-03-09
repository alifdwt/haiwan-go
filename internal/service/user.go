package service

import (
	"errors"

	"github.com/alifdwt/haiwan-go/internal/domain/requests/auth"
	"github.com/alifdwt/haiwan-go/internal/domain/requests/user"
	"github.com/alifdwt/haiwan-go/internal/models"
	"github.com/alifdwt/haiwan-go/internal/repository"
	"github.com/alifdwt/haiwan-go/pkg/hashing"
	"github.com/alifdwt/haiwan-go/pkg/logger"
)

type userService struct {
	Repository repository.UserRepository
	hash       hashing.Hashing
	log        logger.Logger
}

func NewUserService(auth repository.UserRepository, hash hashing.Hashing, logger logger.Logger) *userService {
	return &userService{
		Repository: auth,
		hash:       hash,
		log:        logger,
	}
}

func (s *userService) GetUserAll() (*[]models.User, error) {
	res, err := s.Repository.GetUserAll()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *userService) CreateUser(request *auth.RegisterRequest) (*models.User, error) {
	hashing, err := s.hash.HashPassword(request.Password)
	if err != nil {
		return nil, errors.New("error while hashing password")
	}

	request.Password = hashing

	res, err := s.Repository.CreateUser(request)
	if err != nil {
		return nil, errors.New("error while creating user")
	}

	return res, nil
}

func (s *userService) GetUserById(id int) (*models.User, error) {
	res, err := s.Repository.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *userService) UpdateUserById(id int, request *user.UpdateUserRequest) (*models.User, error) {
	hashing, err := s.hash.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	request.Password = hashing
	res, err := s.Repository.UpdateUserById(id, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *userService) DeleteUserById(id int) (*models.User, error) {
	res, err := s.Repository.DeleteUserById(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
