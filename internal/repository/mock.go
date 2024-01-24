package repository

import (
	"context"

	"tech-challenge-product/internal/canonical"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (m *ProductRepositoryMock) GetAll(ctx context.Context) ([]canonical.Product, error) {
	args := m.Called(ctx)
	return args.Get(0).([]canonical.Product), args.Error(1)
}

func (m *ProductRepositoryMock) Create(ctx context.Context, product *canonical.Product) (*canonical.Product, error) {
	args := m.Called(ctx, product)
	return args.Get(0).(*canonical.Product), args.Error(1)
}

func (m *ProductRepositoryMock) Update(ctx context.Context, id string, product canonical.Product) error {
	args := m.Called(ctx, id, product)
	return args.Error(0)
}

func (m *ProductRepositoryMock) GetByID(ctx context.Context, id string) (*canonical.Product, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*canonical.Product), args.Error(1)
}

func (m *ProductRepositoryMock) GetByCategory(ctx context.Context, category string) ([]canonical.Product, error) {
	args := m.Called(ctx, category)
	return args.Get(0).([]canonical.Product), args.Error(1)
}
