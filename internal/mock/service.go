// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/interfaces.go
//
// Generated by this command:
//
//	mockgen -package mockdb -source internal/service/interfaces.go -destination internal/mock/service.go ...
//

// Package mockdb is a generated GoMock package.
package mockdb

import (
	reflect "reflect"

	auth "github.com/alifdwt/haiwan-go/internal/domain/requests/auth"
	cart "github.com/alifdwt/haiwan-go/internal/domain/requests/cart"
	category "github.com/alifdwt/haiwan-go/internal/domain/requests/category"
	order "github.com/alifdwt/haiwan-go/internal/domain/requests/order"
	product "github.com/alifdwt/haiwan-go/internal/domain/requests/product"
	user "github.com/alifdwt/haiwan-go/internal/domain/requests/user"
	responses "github.com/alifdwt/haiwan-go/internal/domain/responses"
	category0 "github.com/alifdwt/haiwan-go/internal/domain/responses/category"
	order0 "github.com/alifdwt/haiwan-go/internal/domain/responses/order"
	product0 "github.com/alifdwt/haiwan-go/internal/domain/responses/product"
	models "github.com/alifdwt/haiwan-go/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockAuthService is a mock of AuthService interface.
type MockAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService.
type MockAuthServiceMockRecorder struct {
	mock *MockAuthService
}

// NewMockAuthService creates a new mock instance.
func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
	mock := &MockAuthService{ctrl: ctrl}
	mock.recorder = &MockAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockAuthService) Login(input *auth.LoginRequest) (*responses.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", input)
	ret0, _ := ret[0].(*responses.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthServiceMockRecorder) Login(input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthService)(nil).Login), input)
}

// RefreshToken mocks base method.
func (m *MockAuthService) RefreshToken(req auth.RefreshTokenRequest) (*responses.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshToken", req)
	ret0, _ := ret[0].(*responses.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshToken indicates an expected call of RefreshToken.
func (mr *MockAuthServiceMockRecorder) RefreshToken(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshToken", reflect.TypeOf((*MockAuthService)(nil).RefreshToken), req)
}

// Register mocks base method.
func (m *MockAuthService) Register(input *auth.RegisterRequest) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", input)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockAuthServiceMockRecorder) Register(input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuthService)(nil).Register), input)
}

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserService) CreateUser(request *auth.RegisterRequest) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", request)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserServiceMockRecorder) CreateUser(request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserService)(nil).CreateUser), request)
}

// DeleteUserById mocks base method.
func (m *MockUserService) DeleteUserById(id int) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserById", id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUserById indicates an expected call of DeleteUserById.
func (mr *MockUserServiceMockRecorder) DeleteUserById(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserById", reflect.TypeOf((*MockUserService)(nil).DeleteUserById), id)
}

// GetUserAll mocks base method.
func (m *MockUserService) GetUserAll() (*[]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAll")
	ret0, _ := ret[0].(*[]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAll indicates an expected call of GetUserAll.
func (mr *MockUserServiceMockRecorder) GetUserAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAll", reflect.TypeOf((*MockUserService)(nil).GetUserAll))
}

// GetUserById mocks base method.
func (m *MockUserService) GetUserById(id int) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserServiceMockRecorder) GetUserById(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserService)(nil).GetUserById), id)
}

// UpdateUserById mocks base method.
func (m *MockUserService) UpdateUserById(id int, request *user.UpdateUserRequest) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserById", id, request)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserById indicates an expected call of UpdateUserById.
func (mr *MockUserServiceMockRecorder) UpdateUserById(id, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserById", reflect.TypeOf((*MockUserService)(nil).UpdateUserById), id, request)
}

// MockCategoryService is a mock of CategoryService interface.
type MockCategoryService struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryServiceMockRecorder
}

// MockCategoryServiceMockRecorder is the mock recorder for MockCategoryService.
type MockCategoryServiceMockRecorder struct {
	mock *MockCategoryService
}

// NewMockCategoryService creates a new mock instance.
func NewMockCategoryService(ctrl *gomock.Controller) *MockCategoryService {
	mock := &MockCategoryService{ctrl: ctrl}
	mock.recorder = &MockCategoryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryService) EXPECT() *MockCategoryServiceMockRecorder {
	return m.recorder
}

// CreateCategory mocks base method.
func (m *MockCategoryService) CreateCategory(request *category.CreateCategoryRequest) (*category0.CategoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategory", request)
	ret0, _ := ret[0].(*category0.CategoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCategory indicates an expected call of CreateCategory.
func (mr *MockCategoryServiceMockRecorder) CreateCategory(request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategory", reflect.TypeOf((*MockCategoryService)(nil).CreateCategory), request)
}

// DeleteCategoryById mocks base method.
func (m *MockCategoryService) DeleteCategoryById(categoryId int) (*category0.CategoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategoryById", categoryId)
	ret0, _ := ret[0].(*category0.CategoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCategoryById indicates an expected call of DeleteCategoryById.
func (mr *MockCategoryServiceMockRecorder) DeleteCategoryById(categoryId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategoryById", reflect.TypeOf((*MockCategoryService)(nil).DeleteCategoryById), categoryId)
}

// GetCategoryAll mocks base method.
func (m *MockCategoryService) GetCategoryAll() ([]*category0.CategoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryAll")
	ret0, _ := ret[0].([]*category0.CategoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoryAll indicates an expected call of GetCategoryAll.
func (mr *MockCategoryServiceMockRecorder) GetCategoryAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryAll", reflect.TypeOf((*MockCategoryService)(nil).GetCategoryAll))
}

// GetCategoryById mocks base method.
func (m *MockCategoryService) GetCategoryById(categoryId int) (*category0.CategoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryById", categoryId)
	ret0, _ := ret[0].(*category0.CategoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoryById indicates an expected call of GetCategoryById.
func (mr *MockCategoryServiceMockRecorder) GetCategoryById(categoryId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryById", reflect.TypeOf((*MockCategoryService)(nil).GetCategoryById), categoryId)
}

// GetCategoryBySlug mocks base method.
func (m *MockCategoryService) GetCategoryBySlug(slug string) (*category0.CategoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryBySlug", slug)
	ret0, _ := ret[0].(*category0.CategoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoryBySlug indicates an expected call of GetCategoryBySlug.
func (mr *MockCategoryServiceMockRecorder) GetCategoryBySlug(slug any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryBySlug", reflect.TypeOf((*MockCategoryService)(nil).GetCategoryBySlug), slug)
}

// UpdateCategoryById mocks base method.
func (m *MockCategoryService) UpdateCategoryById(id int, updatedCategory *category.UpdateCategoryRequest) (*category0.CategoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCategoryById", id, updatedCategory)
	ret0, _ := ret[0].(*category0.CategoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCategoryById indicates an expected call of UpdateCategoryById.
func (mr *MockCategoryServiceMockRecorder) UpdateCategoryById(id, updatedCategory any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCategoryById", reflect.TypeOf((*MockCategoryService)(nil).UpdateCategoryById), id, updatedCategory)
}

// MockProductService is a mock of ProductService interface.
type MockProductService struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceMockRecorder
}

// MockProductServiceMockRecorder is the mock recorder for MockProductService.
type MockProductServiceMockRecorder struct {
	mock *MockProductService
}

// NewMockProductService creates a new mock instance.
func NewMockProductService(ctrl *gomock.Controller) *MockProductService {
	mock := &MockProductService{ctrl: ctrl}
	mock.recorder = &MockProductServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductService) EXPECT() *MockProductServiceMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockProductService) CreateProduct(request *product.CreateProductRequest) (*product0.ProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", request)
	ret0, _ := ret[0].(*product0.ProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockProductServiceMockRecorder) CreateProduct(request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockProductService)(nil).CreateProduct), request)
}

// DeleteProduct mocks base method.
func (m *MockProductService) DeleteProduct(productId int) (*product0.ProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", productId)
	ret0, _ := ret[0].(*product0.ProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockProductServiceMockRecorder) DeleteProduct(productId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockProductService)(nil).DeleteProduct), productId)
}

// GetProductAll mocks base method.
func (m *MockProductService) GetProductAll() ([]*product0.ProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductAll")
	ret0, _ := ret[0].([]*product0.ProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductAll indicates an expected call of GetProductAll.
func (mr *MockProductServiceMockRecorder) GetProductAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductAll", reflect.TypeOf((*MockProductService)(nil).GetProductAll))
}

// GetProductById mocks base method.
func (m *MockProductService) GetProductById(productId int) (*product0.ProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductById", productId)
	ret0, _ := ret[0].(*product0.ProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductById indicates an expected call of GetProductById.
func (mr *MockProductServiceMockRecorder) GetProductById(productId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductById", reflect.TypeOf((*MockProductService)(nil).GetProductById), productId)
}

// GetProductBySlug mocks base method.
func (m *MockProductService) GetProductBySlug(slug string) (*product0.ProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductBySlug", slug)
	ret0, _ := ret[0].(*product0.ProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductBySlug indicates an expected call of GetProductBySlug.
func (mr *MockProductServiceMockRecorder) GetProductBySlug(slug any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductBySlug", reflect.TypeOf((*MockProductService)(nil).GetProductBySlug), slug)
}

// UpdateProduct mocks base method.
func (m *MockProductService) UpdateProduct(productId int, request *product.UpdateProductRequest) (*product0.ProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", productId, request)
	ret0, _ := ret[0].(*product0.ProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockProductServiceMockRecorder) UpdateProduct(productId, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockProductService)(nil).UpdateProduct), productId, request)
}

// MockCartService is a mock of CartService interface.
type MockCartService struct {
	ctrl     *gomock.Controller
	recorder *MockCartServiceMockRecorder
}

// MockCartServiceMockRecorder is the mock recorder for MockCartService.
type MockCartServiceMockRecorder struct {
	mock *MockCartService
}

// NewMockCartService creates a new mock instance.
func NewMockCartService(ctrl *gomock.Controller) *MockCartService {
	mock := &MockCartService{ctrl: ctrl}
	mock.recorder = &MockCartServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCartService) EXPECT() *MockCartServiceMockRecorder {
	return m.recorder
}

// CreateCart mocks base method.
func (m *MockCartService) CreateCart(cartRequest *cart.CartCreateRequest) (*models.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCart", cartRequest)
	ret0, _ := ret[0].(*models.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCart indicates an expected call of CreateCart.
func (mr *MockCartServiceMockRecorder) CreateCart(cartRequest any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCart", reflect.TypeOf((*MockCartService)(nil).CreateCart), cartRequest)
}

// DeleteCart mocks base method.
func (m *MockCartService) DeleteCart(cartId int) (*models.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCart", cartId)
	ret0, _ := ret[0].(*models.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCart indicates an expected call of DeleteCart.
func (mr *MockCartServiceMockRecorder) DeleteCart(cartId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCart", reflect.TypeOf((*MockCartService)(nil).DeleteCart), cartId)
}

// DeleteCartMany mocks base method.
func (m *MockCartService) DeleteCartMany(cartIds cart.DeleteCartRequest) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCartMany", cartIds)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCartMany indicates an expected call of DeleteCartMany.
func (mr *MockCartServiceMockRecorder) DeleteCartMany(cartIds any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCartMany", reflect.TypeOf((*MockCartService)(nil).DeleteCartMany), cartIds)
}

// GetCartAll mocks base method.
func (m *MockCartService) GetCartAll(userId int) (*[]models.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartAll", userId)
	ret0, _ := ret[0].(*[]models.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartAll indicates an expected call of GetCartAll.
func (mr *MockCartServiceMockRecorder) GetCartAll(userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartAll", reflect.TypeOf((*MockCartService)(nil).GetCartAll), userId)
}

// MockOrderService is a mock of OrderService interface.
type MockOrderService struct {
	ctrl     *gomock.Controller
	recorder *MockOrderServiceMockRecorder
}

// MockOrderServiceMockRecorder is the mock recorder for MockOrderService.
type MockOrderServiceMockRecorder struct {
	mock *MockOrderService
}

// NewMockOrderService creates a new mock instance.
func NewMockOrderService(ctrl *gomock.Controller) *MockOrderService {
	mock := &MockOrderService{ctrl: ctrl}
	mock.recorder = &MockOrderServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderService) EXPECT() *MockOrderServiceMockRecorder {
	return m.recorder
}

// CreateOrder mocks base method.
func (m *MockOrderService) CreateOrder(userId int, request *order.CreateOrderRequest) (*models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", userId, request)
	ret0, _ := ret[0].(*models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrderServiceMockRecorder) CreateOrder(userId, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrderService)(nil).CreateOrder), userId, request)
}

// GetOrderAll mocks base method.
func (m *MockOrderService) GetOrderAll() (*[]order0.OrderResponses, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderAll")
	ret0, _ := ret[0].(*[]order0.OrderResponses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderAll indicates an expected call of GetOrderAll.
func (mr *MockOrderServiceMockRecorder) GetOrderAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderAll", reflect.TypeOf((*MockOrderService)(nil).GetOrderAll))
}

// GetOrderById mocks base method.
func (m *MockOrderService) GetOrderById(orderId int) (*order0.OrderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderById", orderId)
	ret0, _ := ret[0].(*order0.OrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderById indicates an expected call of GetOrderById.
func (mr *MockOrderServiceMockRecorder) GetOrderById(orderId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderById", reflect.TypeOf((*MockOrderService)(nil).GetOrderById), orderId)
}

// GetOrdersByUser mocks base method.
func (m *MockOrderService) GetOrdersByUser(userId int) (*[]models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrdersByUser", userId)
	ret0, _ := ret[0].(*[]models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrdersByUser indicates an expected call of GetOrdersByUser.
func (mr *MockOrderServiceMockRecorder) GetOrdersByUser(userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrdersByUser", reflect.TypeOf((*MockOrderService)(nil).GetOrdersByUser), userId)
}
