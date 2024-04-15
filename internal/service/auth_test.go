package service

import (
	"testing"

	"github.com/alifdwt/haiwan-go/internal/domain/requests/auth"
	"github.com/alifdwt/haiwan-go/internal/models"
	"github.com/alifdwt/haiwan-go/pkg/random"
	"github.com/gosimple/slug"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) (userData models.User, userPassword string) {
	password := random.RandomString(6)
	name := random.RandomOwner()
	arg := &auth.RegisterRequest{
		Name:             name,
		Email:            random.RandomEmail(slug.Make(name)),
		Password:         password,
		Confirm_password: password,
	}

	user, err := testQueries.Auth.Register(arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.Password)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)

	require.Zero(t, user.DeletedAt)

	return *user, password
}

func TestRegisterUser(t *testing.T) {
	createRandomUser(t)
}

func TestLogin(t *testing.T) {
	// Skenario berhasil login
	user, password := createRandomUser(t)
	loginReq := &auth.LoginRequest{
		Email:    user.Email,
		Password: password,
	}
	token, err := testQueries.Auth.Login(loginReq)
	require.NoError(t, err)
	require.NotNil(t, token)

	// Skenario user tidak ditemukan
	invalidUserReq := &auth.LoginRequest{
		Email:    "nonexistent@example.com",
		Password: "password",
	}
	_, err = testQueries.Auth.Login(invalidUserReq)
	require.Error(t, err)

	// Skenario password salah
	invalidPasswordReq := &auth.LoginRequest{
		Email:    user.Email,
		Password: "wrongpassword",
	}
	_, err = testQueries.Auth.Login(invalidPasswordReq)
	require.Error(t, err)
}
