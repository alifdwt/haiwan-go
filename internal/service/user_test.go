package service

import (
	"testing"
	"time"

	"github.com/alifdwt/haiwan-go/internal/domain/requests/user"
	"github.com/alifdwt/haiwan-go/pkg/random"
	"github.com/gosimple/slug"
	"github.com/stretchr/testify/require"
)

// func TestCreateUser(t *testing.T) {
// 	createRandomUser(t)
// }

func TestGetUserAll(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomUser(t)
	}

	users, err := testQueries.User.GetUserAll()
	require.NoError(t, err)
	require.NotEmpty(t, users)

	for _, user := range *users {
		require.NotEmpty(t, user)
	}
}

func TestGetUserById(t *testing.T) {
	user1, _ := createRandomUser(t)
	user2, err := testQueries.User.GetUserById(int(user1.ID))
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Password, user2.Password)

	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt, time.Second)
}

func TestUpdateUserById(t *testing.T) {
	user1, _ := createRandomUser(t)
	name := random.RandomOwner()
	arg := &user.UpdateUserRequest{
		Name:     name,
		Email:    random.RandomEmail(slug.Make(name)),
		Password: random.RandomString(6),
	}

	user2, err := testQueries.User.UpdateUserById(int(user1.ID), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.NotEqual(t, user1.Name, user2.Name)
	require.NotEqual(t, user1.Email, user2.Email)
	require.NotEqual(t, user1.Password, user2.Password)
}

func TestDeleteUserById(t *testing.T) {
	user1, _ := createRandomUser(t)

	userDeleted, err := testQueries.User.DeleteUserById(int(user1.ID))
	require.NoError(t, err)
	require.NotEmpty(t, userDeleted)

	// user2, err := testQueries.User.GetUserById(int(user1.ID))
	// require.Error()
}
