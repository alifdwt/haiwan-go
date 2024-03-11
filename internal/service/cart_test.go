package service

import (
	"fmt"
	"testing"

	cartreq "github.com/alifdwt/haiwan-go/internal/domain/requests/cart"
	"github.com/alifdwt/haiwan-go/internal/models"
	"github.com/alifdwt/haiwan-go/pkg/random"
	"github.com/stretchr/testify/require"
)

func createRandomCart(t *testing.T) (cartData models.Cart, userId int) {
	user, _ := createRandomUser(t)
	category := createRandomCategory(t)
	product := createRandomProduct(t, category.ID)

	quantity := int(random.RandomInt(1, 5))

	arg := &cartreq.CartCreateRequest{
		Name:         product.Name,
		Price:        fmt.Sprint(product.Price),
		ImageProduct: product.ImagePath,
		Quantity:     quantity,
		Weight:       int(product.Weight) * quantity,
		ProductID:    product.ID,
		UserID:       int(user.ID),
	}

	cart, err := testQueries.Cart.CreateCart(arg)
	require.NoError(t, err)
	require.NotEmpty(t, cart)

	require.Equal(t, arg.Name, cart.Name)
	require.Equal(t, arg.Price, cart.Price)
	require.Equal(t, arg.ImageProduct, cart.Image)
	require.Equal(t, arg.Quantity, cart.Quantity)
	require.Equal(t, arg.Weight, cart.Weight)
	require.Equal(t, arg.ProductID, int(cart.ProductID))
	require.Equal(t, arg.UserID, int(cart.UserID))

	// Failed scenario
	failedCart, err := testQueries.Cart.CreateCart(&cartreq.CartCreateRequest{})
	require.Error(t, err)
	require.Empty(t, failedCart)

	return *cart, int(user.ID)
}

func TestCreateCart(t *testing.T) {
	createRandomCart(t)
}

func TestGetCartAll(t *testing.T) {
	_, userId := createRandomCart(t)

	cart2, err := testQueries.Cart.GetCartAll(userId)
	require.NoError(t, err)
	require.NotEmpty(t, cart2)

	// Failed scenario
	failedCart, err := testQueries.Cart.GetCartAll(123456)
	require.Error(t, err)
	require.Empty(t, failedCart)
}

func TestDeleteCart(t *testing.T) {
	cart, _ := createRandomCart(t)
	deletedCart, err := testQueries.Cart.DeleteCart(int(cart.ID))
	require.NoError(t, err)
	require.NotEmpty(t, deletedCart)

	require.NotNil(t, deletedCart.DeletedAt)

	// Failed scenario
	failedDeletedCart, err := testQueries.Cart.DeleteCart(123456)
	require.Error(t, err)
	require.Empty(t, failedDeletedCart)
}
