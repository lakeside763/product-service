// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/moses/Desktop/folder/projects/go-programs/product-service/internal/ports/interfaces/product_interface.go

// Package interfaces is a generated GoMock package.
package interfaces

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/lakeside763/product-service/internal/core/models"
)

// MockProducts is a mock of Products interface.
type MockProducts struct {
	ctrl     *gomock.Controller
	recorder *MockProductsMockRecorder
}

// CheckProductExistsByName implements interfaces.Products.
func (m *MockProducts) CheckProductExistsByName(name string) error {
	panic("unimplemented")
}

// CreateProduct implements interfaces.Products.
func (m *MockProducts) CreateProduct(p models.CreateProductInput) (*models.Product, error) {
	panic("unimplemented")
}

// MockProductsMockRecorder is the mock recorder for MockProducts.
type MockProductsMockRecorder struct {
	mock *MockProducts
}

// NewMockProducts creates a new mock instance.
func NewMockProducts(ctrl *gomock.Controller) *MockProducts {
	mock := &MockProducts{ctrl: ctrl}
	mock.recorder = &MockProductsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProducts) EXPECT() *MockProductsMockRecorder {
	return m.recorder
}

// GetMaxDiscount mocks base method.
func (m *MockProducts) GetMaxDiscount(category, sku string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxDiscount", category, sku)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMaxDiscount indicates an expected call of GetMaxDiscount.
func (mr *MockProductsMockRecorder) GetMaxDiscount(category, sku interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxDiscount", reflect.TypeOf((*MockProducts)(nil).GetMaxDiscount), category, sku)
}

// GetProducts mocks base method.
func (m *MockProducts) GetProducts(category string, priceLessThan int, cursorId string, pageSize int) ([]*models.Product, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducts", category, priceLessThan, cursorId, pageSize)
	ret0, _ := ret[0].([]*models.Product)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetProducts indicates an expected call of GetProducts.
func (mr *MockProductsMockRecorder) GetProducts(category, priceLessThan, cursorId, pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducts", reflect.TypeOf((*MockProducts)(nil).GetProducts), category, priceLessThan, cursorId, pageSize)
}
