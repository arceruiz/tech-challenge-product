package mocks

import (
	"context"
	"tech-challenge-product/internal/canonical"

	"github.com/stretchr/testify/mock"
)

type ProductServiceMock struct {
	mock.Mock
}

func (m *ProductServiceMock) GetAll(ctx context.Context) ([]canonical.Product, error) {
	args := m.Called(ctx)
	return args.Get(0).([]canonical.Product), args.Error(1)
}

func (m *ProductServiceMock) Create(ctx context.Context, product *canonical.Product) (*canonical.Product, error) {
	args := m.Called(ctx, product)
	return args.Get(0).(*canonical.Product), args.Error(1)
}

func (m *ProductServiceMock) Update(ctx context.Context, id string, product canonical.Product) error {
	args := m.Called(ctx, id, product)
	return args.Error(0)
}

func (m *ProductServiceMock) GetByID(ctx context.Context, id string) (*canonical.Product, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*canonical.Product), args.Error(1)
}

func (m *ProductServiceMock) GetByCategory(ctx context.Context, category string) ([]canonical.Product, error) {
	args := m.Called(ctx, category)
	return args.Get(0).([]canonical.Product), args.Error(1)
}

func (m *ProductServiceMock) Remove(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
