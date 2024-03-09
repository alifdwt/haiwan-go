package repository

import (
	"testing"

	"github.com/alifdwt/haiwan-go/internal/domain/requests/auth"
	"github.com/alifdwt/haiwan-go/internal/models"
	"github.com/alifdwt/haiwan-go/pkg/random"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) *models.User {
	password := random.RandomString(6)
	arg := &auth.RegisterRequest{
		Name:             random.RandomOwner(),
		Email:            random.RandomEmail(),
		Password:         password,
		Confirm_password: password,
	}

	userRepository := NewUserRepository(setDatabase())
	user, err := UserRepository.CreateUser(userRepository, arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.Password)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)

	require.Zero(t, user.DeletedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
