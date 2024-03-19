package service

import (
	"context"
	"errors"
	"tech-challenge-product/internal/canonical"
	"tech-challenge-product/internal/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/undefinedlabs/go-mpatch"
)

func TestProductService_GetByID(t *testing.T) {

	type Given struct {
		id          string
		productRepo func() repository.ProductRepository
	}
	type Expected struct {
		err assert.ErrorAssertionFunc
	}
	tests := map[string]struct {
		given    Given
		expected Expected
	}{

		"given product with main fields filled, must return created paymend with all fields filled": {
			given: Given{
				id: "1234",
				productRepo: func() repository.ProductRepository {
					repoMock := ProductRepositoryMock{}

					repoMock.On("GetByID", mock.Anything, "1234").Return(&canonical.Product{
						ID:          "product_valid_id",
						Name:        "product_valid_name",
						Description: "product_valid_desc",
						Price:       10,
						Category:    "product_valid_category",
						Status:      0,
						ImagePath:   "product_valid_imgpath",
					}, nil)

					return &repoMock
				},
			},
			expected: Expected{
				err: assert.NoError,
			},
		},
	}

	for _, tc := range tests {
		svc := productService{
			repo: tc.given.productRepo(),
		}
		_, err := svc.GetByID(context.Background(), tc.given.id)

		tc.expected.err(t, err)
	}
}

func TestProductService_GetAll(t *testing.T) {

	type Given struct {
		productRepo func() repository.ProductRepository
	}
	type Expected struct {
		err assert.ErrorAssertionFunc
	}
	tests := map[string]struct {
		given    Given
		expected Expected
	}{

		"given product with main fields filled, must return created paymend with all fields filled": {
			given: Given{
				productRepo: func() repository.ProductRepository {
					repoMock := &ProductRepositoryMock{}
					repoMock.On("GetAll", mock.Anything).Return([]canonical.Product{
						{
							ID:          "product_valid_id",
							Name:        "product_valid_name",
							Description: "product_valid_desc",
							Price:       10,
							Category:    "product_valid_category",
							Status:      0,
							ImagePath:   "product_valid_imgpath",
						},
						{
							ID:          "product_valid_id",
							Name:        "product_valid_name",
							Description: "product_valid_desc",
							Price:       10,
							Category:    "product_valid_category",
							Status:      0,
							ImagePath:   "product_valid_imgpath",
						},
					}, nil)
					return repoMock
				},
			},
			expected: Expected{
				err: assert.NoError,
			},
		},
	}

	for _, tc := range tests {
		svc := productService{
			repo: tc.given.productRepo(),
		}

		_, err := svc.GetAll(context.Background())

		tc.expected.err(t, err)
	}
}

func TestProductService_GetByCategory(t *testing.T) {

	type Given struct {
		category    string
		productRepo func() repository.ProductRepository
	}
	type Expected struct {
		err assert.ErrorAssertionFunc
	}
	tests := map[string]struct {
		given    Given
		expected Expected
	}{

		"given product with main fields filled, must return created paymend with all fields filled": {
			given: Given{
				category: "product_valid_category",
				productRepo: func() repository.ProductRepository {
					repoMock := &ProductRepositoryMock{}
					repoMock.On("GetByCategory", mock.Anything, "product_valid_category").Return([]canonical.Product{
						{
							ID:          "product_valid_id",
							Name:        "product_valid_name",
							Description: "product_valid_desc",
							Price:       10,
							Category:    "product_valid_category",
							Status:      0,
							ImagePath:   "product_valid_imgpath",
						},
						{
							ID:          "product_valid_id",
							Name:        "product_valid_name",
							Description: "product_valid_desc",
							Price:       10,
							Category:    "product_valid_category",
							Status:      0,
							ImagePath:   "product_valid_imgpath",
						},
					}, nil)
					return repoMock
				},
			},
			expected: Expected{
				err: assert.NoError,
			},
		},
	}

	for _, tc := range tests {
		svc := productService{
			repo: tc.given.productRepo(),
		}
		_, err := svc.GetByCategory(context.Background(), tc.given.category)

		tc.expected.err(t, err)
	}
}

func TestProductService_Create(t *testing.T) {

	mpatch.PatchMethod(time.Now, func() time.Time {
		return time.Date(2020, 11, 01, 00, 00, 00, 0, time.UTC)
	})
	mpatch.PatchMethod(canonical.NewUUID, func() string {
		return "product_valid_id"
	})

	type Given struct {
		product     *canonical.Product
		productRepo func() repository.ProductRepository
	}
	type Expected struct {
		err assert.ErrorAssertionFunc
	}
	tests := map[string]struct {
		given    Given
		expected Expected
	}{
		"given product with main fields filled, must return created paymend with all fields filled": {
			given: Given{
				product: &canonical.Product{
					Name:        "product_valid_name",
					Description: "product_valid_desc",
					Price:       10,
					Category:    "product_valid_category",
					Status:      0,
					ImagePath:   "product_valid_imgpath",
				},
				productRepo: func() repository.ProductRepository {
					product := &canonical.Product{
						ID:          "product_valid_id",
						Name:        "product_valid_name",
						Description: "product_valid_desc",
						Price:       10,
						Category:    "product_valid_category",
						Status:      0,
						ImagePath:   "product_valid_imgpath",
					}
					repoMock := &ProductRepositoryMock{}
					repoMock.On("Create", mock.Anything, product).Return(product, nil)
					return repoMock
				},
			},
			expected: Expected{
				err: assert.NoError,
			},
		},
		"given error creating, must return error": {
			given: Given{
				product: &canonical.Product{
					Name:        "product_valid_name",
					Description: "product_valid_desc",
					Price:       10,
					Category:    "product_valid_category",
					Status:      0,
					ImagePath:   "product_valid_imgpath",
				},
				productRepo: func() repository.ProductRepository {
					repoMock := &ProductRepositoryMock{}
					repoMock.On("Create", mock.Anything, mock.Anything).Return(&canonical.Product{}, errors.New("error creating product"))
					return repoMock
				},
			},
			expected: Expected{
				err: assert.Error,
			},
		},
	}

	for _, tc := range tests {
		svc := productService{
			repo: tc.given.productRepo(),
		}
		_, err := svc.Create(context.Background(), tc.given.product)

		tc.expected.err(t, err)
	}
}

func TestProductService_Update(t *testing.T) {

	mpatch.PatchMethod(time.Now, func() time.Time {
		return time.Date(2020, 11, 01, 00, 00, 00, 0, time.UTC)
	})
	mpatch.PatchMethod(canonical.NewUUID, func() string {
		return "product_valid_id"
	})

	type Given struct {
		product     canonical.Product
		productID   string
		productRepo func() repository.ProductRepository
	}
	type Expected struct {
		err assert.ErrorAssertionFunc
	}
	tests := map[string]struct {
		given    Given
		expected Expected
	}{
		"given product with main fields filled, must return created paymend with all fields filled": {
			given: Given{
				productID: "product_valid_id",
				product: canonical.Product{
					Name:        "product_valid_name",
					Description: "product_valid_desc",
					Price:       10,
					Category:    "product_valid_category",
					Status:      0,
					ImagePath:   "product_valid_imgpath",
				},
				productRepo: func() repository.ProductRepository {
					product := canonical.Product{
						ID:          "product_valid_id",
						Name:        "product_valid_name",
						Description: "product_valid_desc",
						Price:       10,
						Category:    "product_valid_category",
						Status:      0,
						ImagePath:   "product_valid_imgpath",
					}
					repoMock := &ProductRepositoryMock{}
					repoMock.On("Update", mock.Anything, "product_valid_id", product).Return(nil)
					return repoMock
				},
			},
			expected: Expected{
				err: assert.NoError,
			},
		},
		"given error creating, must return error": {
			given: Given{
				productID: "product_valid_id",
				product: canonical.Product{
					Name:        "product_valid_name",
					Description: "product_valid_desc",
					Price:       10,
					Category:    "product_valid_category",
					Status:      0,
					ImagePath:   "product_valid_imgpath",
				},
				productRepo: func() repository.ProductRepository {
					repoMock := &ProductRepositoryMock{}
					repoMock.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error creating product"))
					return repoMock
				},
			},
			expected: Expected{
				err: assert.Error,
			},
		},
	}

	for _, tc := range tests {
		svc := productService{
			repo: tc.given.productRepo(),
		}

		err := svc.Update(context.Background(), tc.given.productID, tc.given.product)

		tc.expected.err(t, err)
	}
}

func TestProductService_Remove(t *testing.T) {

	type Given struct {
		id          string
		productRepo func() repository.ProductRepository
	}
	type Expected struct {
		err assert.ErrorAssertionFunc
	}
	tests := map[string]struct {
		given    Given
		expected Expected
	}{
		"given valid product id, must update product correctly": {
			given: Given{
				id: "product_valid_id",
				productRepo: func() repository.ProductRepository {
					repoMock := &ProductRepositoryMock{}
					repoMock.On("GetByID", mock.Anything, "product_valid_id").Return(&canonical.Product{
						ID:          "product_valid_id",
						Name:        "product_valid_name",
						Description: "product_valid_desc",
						Price:       10,
						Category:    "product_valid_category",
						Status:      0,
						ImagePath:   "product_valid_imgpath",
					}, nil)

					repoMock.On("Update", mock.Anything, "product_valid_id", canonical.Product{
						ID:          "product_valid_id",
						Name:        "product_valid_name",
						Description: "product_valid_desc",
						Price:       10,
						Category:    "product_valid_category",
						Status:      1,
						ImagePath:   "product_valid_imgpath",
					}).Return(nil)
					return repoMock
				},
			},
			expected: Expected{
				err: assert.NoError,
			},
		},
		"given error getting product must return error": {
			given: Given{
				id: "product_valid_id",
				productRepo: func() repository.ProductRepository {
					repoMock := &ProductRepositoryMock{}
					repoMock.On("GetByID", mock.Anything, "product_valid_id").Return(&canonical.Product{
						ID:          "product_valid_id",
						Name:        "product_valid_name",
						Description: "product_valid_desc",
						Price:       10,
						Category:    "product_valid_category",
						Status:      0,
						ImagePath:   "product_valid_imgpath",
					}, errors.New("error getting product"))
					return repoMock
				},
			},
			expected: Expected{
				err: assert.Error,
			},
		},
		"given no product found, must return error": {
			given: Given{
				id: "product_valid_id",
				productRepo: func() repository.ProductRepository {
					repoMock := &ProductRepositoryMock{}
					repoMock.On("GetByID", mock.Anything, "product_valid_id").Return(nil, nil)
					return repoMock
				},
			},
			expected: Expected{
				err: assert.Error,
			},
		},
		"given error on product update, must return error": {
			given: Given{
				id: "product_valid_id",
				productRepo: func() repository.ProductRepository {
					repoMock := &ProductRepositoryMock{}
					repoMock.On("GetByID", mock.Anything, "product_valid_id").Return(&canonical.Product{
						ID:          "product_valid_id",
						Name:        "product_valid_name",
						Description: "product_valid_desc",
						Price:       10,
						Category:    "product_valid_category",
						Status:      0,
						ImagePath:   "product_valid_imgpath",
					}, nil)

					repoMock.On("Update", mock.Anything, "product_valid_id", canonical.Product{
						ID:          "product_valid_id",
						Name:        "product_valid_name",
						Description: "product_valid_desc",
						Price:       10,
						Category:    "product_valid_category",
						Status:      1,
						ImagePath:   "product_valid_imgpath",
					}).Return(errors.New("error updating product"))
					return repoMock
				},
			},
			expected: Expected{
				err: assert.Error,
			},
		},
	}

	for _, tc := range tests {
		svc := productService{
			repo: tc.given.productRepo(),
		}
		err := svc.Remove(context.Background(), tc.given.id)

		tc.expected.err(t, err)
	}
}

func TestGetProductsWithId(t *testing.T) {
	mock := &ProductRepositoryMock{}

	svc := productService{
		repo: mock,
	}

	mock.On("GetProductsWithId").Return([]canonical.Product{
		{
			ID:          "123",
			Name:        "x",
			Description: "y",
			Price:       1.0,
			Category:    "l",
			Status:      canonical.STATUS_ACTIVE,
			ImagePath:   "k",
		},
	}, nil)

	products, err := svc.GetProductsWithId(context.Background(), []string{
		"123",
	})

	assert.Nil(t, err)
	assert.NotNil(t, products)
}
