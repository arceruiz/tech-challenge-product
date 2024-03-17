package service

import (
	"context"
	"tech-challenge-product/internal/canonical"
	"tech-challenge-product/internal/repository"
)

type ProductService interface {
	GetAll(context.Context) ([]canonical.Product, error)
	Create(ctx context.Context, product *canonical.Product) (*canonical.Product, error)
	Update(context.Context, string, canonical.Product) error
	GetByID(context.Context, string) (*canonical.Product, error)
	GetByCategory(context.Context, string) ([]canonical.Product, error)
	Remove(context.Context, string) error
	GetProductsWithId(ctx context.Context, ids []string) ([]canonical.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService() ProductService {
	return &productService{
		repo: repository.NewProductRepo(),
	}
}

func (s *productService) GetProductsWithId(ctx context.Context, ids []string) ([]canonical.Product, error) {
	return s.repo.GetProductsWithId(ctx, ids)
}

func (s *productService) GetAll(ctx context.Context) ([]canonical.Product, error) {
	return s.repo.GetAll(ctx)
}

func (s *productService) Create(ctx context.Context, product *canonical.Product) (*canonical.Product, error) {
	product.ID = canonical.NewUUID()
	return s.repo.Create(ctx, product)
}

func (s *productService) Update(ctx context.Context, id string, updatedProduct canonical.Product) error {
	if updatedProduct.ID == "" {
		updatedProduct.ID = id
	}
	return s.repo.Update(ctx, id, updatedProduct)
}

func (s *productService) GetByID(ctx context.Context, id string) (*canonical.Product, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *productService) GetByCategory(ctx context.Context, id string) ([]canonical.Product, error) {
	return s.repo.GetByCategory(ctx, id)
}

func (s *productService) Remove(ctx context.Context, id string) error {
	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if product == nil {
		return canonical.ErrorNotFound
	}
	product.Status = 1
	err = s.repo.Update(ctx, id, *product)
	if err != nil {
		return err
	}
	return nil
}
