package repository_test

import (
	"context"
	"tech-challenge-product/internal/canonical"
	"tech-challenge-product/internal/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/undefinedlabs/go-mpatch"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestProductRepository_GetByID(t *testing.T) {

	mpatch.PatchMethod(time.Now, func() time.Time {
		return time.Date(2020, 11, 01, 00, 00, 00, 0, time.UTC)
	})

	type Given struct {
		mtestFunc func(mt *mtest.T)
	}
	type Expected struct {
	}
	tests := map[string]struct {
		given    Given
		expected Expected
	}{
		"given valid search result, must return valid product": {
			given: Given{
				mtestFunc: func(mt *mtest.T) {
					repo := repository.NewProductRepo(mt.DB)
					mt.AddMockResponses(
						mtest.CreateCursorResponse(1,
							"product.product",
							mtest.FirstBatch,
							bson.D{
								{Key: "_id", Value: "product_valid_id"},
								{Key: "name", Value: "product_valid_name"},
								{Key: "description", Value: "product_valid_desc"},
								{Key: "price", Value: 10.0},
								{Key: "category", Value: "product_valid_category"},
								{Key: "status", Value: 0},
								{Key: "image_path", Value: "product_valid_imgpath"},
							},
						),
					)
					product, err := repo.GetByID(context.Background(), "product_valid_id")
					assert.Nil(t, err)
					assert.Equal(t, product.ID, "product_valid_id")
					assert.Equal(t, product.Status, canonical.STATUS_ACTIVE)
				},
			},
		},
		"given entity not found must return error": {
			given: Given{
				mtestFunc: func(mt *mtest.T) {
					repo := repository.NewProductRepo(mt.DB)
					mt.AddMockResponses(mtest.CreateCursorResponse(0, "product.product", mtest.FirstBatch))
					product, err := repo.GetByID(context.Background(), "asd")
					assert.NotNil(t, err)
					assert.Equal(t, err.Error(), "mongo: no documents in result")
					assert.Nil(t, product)
				},
			},
		},
	}

	for _, tc := range tests {
		db := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
		db.Run("", tc.given.mtestFunc)
	}
}

func TestProductRepository_GetAll(t *testing.T) {

	mpatch.PatchMethod(time.Now, func() time.Time {
		return time.Date(2020, 11, 01, 00, 00, 00, 0, time.UTC)
	})

	type Given struct {
		mtestFunc func(mt *mtest.T)
	}
	type Expected struct {
	}
	tests := map[string]struct {
		given    Given
		expected Expected
	}{
		"given valid search result, must return valid product": {
			given: Given{
				mtestFunc: func(mt *mtest.T) {
					repo := repository.NewProductRepo(mt.DB)

					first := mtest.CreateCursorResponse(1, "product.product", mtest.FirstBatch, bson.D{
						{Key: "_id", Value: "product_valid_id"},
						{Key: "name", Value: "product_valid_name"},
						{Key: "description", Value: "product_valid_desc"},
						{Key: "price", Value: 10.0},
						{Key: "category", Value: "product_valid_category"},
						{Key: "status", Value: 0},
						{Key: "image_path", Value: "product_valid_imgpath"},
					})
					getMore := mtest.CreateCursorResponse(1, "product.product", mtest.NextBatch, bson.D{
						{Key: "_id", Value: "product_valid_id"},
						{Key: "name", Value: "product_valid_name"},
						{Key: "description", Value: "product_valid_desc"},
						{Key: "price", Value: 10.0},
						{Key: "category", Value: "product_valid_category"},
						{Key: "status", Value: 0},
						{Key: "image_path", Value: "product_valid_imgpath"},
					})
					lastCursor := mtest.CreateCursorResponse(0, "product.product", mtest.NextBatch)
					mt.AddMockResponses(first, getMore, lastCursor)

					products, err := repo.GetAll(context.Background())
					assert.Nil(t, err)
					for _, product := range products {
						assert.Equal(t, product.ID, "product_valid_id")
						assert.Equal(t, product.Status, canonical.STATUS_ACTIVE)
					}
				},
			},
		},
		"given entity not found must return error": {
			given: Given{
				mtestFunc: func(mt *mtest.T) {
					repo := repository.NewProductRepo(mt.DB)
					mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{Message: "mongo: no documents in result"}))
					product, err := repo.GetAll(context.Background())
					assert.NotNil(t, err)
					assert.Equal(t, err.Error(), "write command error: [{write errors: [{mongo: no documents in result}]}, {<nil>}]")
					assert.Nil(t, product)
				},
			},
		},
	}

	for _, tc := range tests {
		db := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
		db.Run("", tc.given.mtestFunc)
	}
}

func TestProductRepository_GetByCategory(t *testing.T) {

	mpatch.PatchMethod(time.Now, func() time.Time {
		return time.Date(2020, 11, 01, 00, 00, 00, 0, time.UTC)
	})

	type Given struct {
		mtestFunc func(mt *mtest.T)
	}
	type Expected struct {
	}
	tests := map[string]struct {
		given    Given
		expected Expected
	}{
		"given valid search result, must return valid product": {
			given: Given{
				mtestFunc: func(mt *mtest.T) {
					repo := repository.NewProductRepo(mt.DB)

					first := mtest.CreateCursorResponse(1, "product.product", mtest.FirstBatch, bson.D{
						{Key: "_id", Value: "product_valid_id"},
						{Key: "name", Value: "product_valid_name"},
						{Key: "description", Value: "product_valid_desc"},
						{Key: "price", Value: 10.0},
						{Key: "category", Value: "product_valid_category"},
						{Key: "status", Value: 0},
						{Key: "image_path", Value: "product_valid_imgpath"},
					})
					getMore := mtest.CreateCursorResponse(1, "product.product", mtest.NextBatch, bson.D{
						{Key: "_id", Value: "product_valid_id"},
						{Key: "name", Value: "product_valid_name"},
						{Key: "description", Value: "product_valid_desc"},
						{Key: "price", Value: 10.0},
						{Key: "category", Value: "product_valid_category"},
						{Key: "status", Value: 0},
						{Key: "image_path", Value: "product_valid_imgpath"},
					})
					lastCursor := mtest.CreateCursorResponse(0, "product.product", mtest.NextBatch)
					mt.AddMockResponses(first, getMore, lastCursor)

					products, err := repo.GetByCategory(context.Background(), "product_valid_category")
					for _, product := range products {
						assert.Nil(t, err)
						assert.Equal(t, product.Category, "product_valid_category")
						assert.Equal(t, product.Status, canonical.STATUS_ACTIVE)
					}
				},
			},
		},
		"given entity not found must return error": {
			given: Given{
				mtestFunc: func(mt *mtest.T) {
					repo := repository.NewProductRepo(mt.DB)
					mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{Message: "mongo: no documents in result"}))
					product, err := repo.GetByCategory(context.Background(), "asd")
					assert.NotNil(t, err)
					assert.Equal(t, err.Error(), "write command error: [{write errors: [{mongo: no documents in result}]}, {<nil>}]")
					assert.Nil(t, product)
				},
			},
		},
	}

	for _, tc := range tests {
		db := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
		db.Run("", tc.given.mtestFunc)
	}
}

func TestCreate(t *testing.T) {

	mpatch.PatchMethod(time.Now, func() time.Time {
		return time.Date(2020, 11, 01, 00, 00, 00, 0, time.UTC)
	})

	type Given struct {
		mtestFunc func(mt *mtest.T)
	}
	type Expected struct {
	}
	tests := map[string]struct {
		given    Given
		expected Expected
	}{
		"given given no error saving must return correct entity": {
			given: Given{
				mtestFunc: func(mt *mtest.T) {
					repo := repository.NewProductRepo(mt.DB)
					mt.AddMockResponses(mtest.CreateSuccessResponse())

					product := &canonical.Product{
						ID:          "product_valid_id",
						Name:        "product_valid_name",
						Description: "product_valid_desc",
						Price:       10,
						Category:    "product_valid_category",
						Status:      0,
						ImagePath:   "product_valid_imgpath",
					}

					createdProduct, err := repo.Create(context.Background(), product)

					assert.Nil(t, err)
					assert.Equal(t, createdProduct, product)

				},
			},
		},
		"given given error saving must return error": {
			given: Given{
				mtestFunc: func(mt *mtest.T) {
					repo := repository.NewProductRepo(mt.DB)
					mt.AddMockResponses(
						bson.D{
							{Key: "ok", Value: -1},
						},
					)

					product := &canonical.Product{
						ID:          "product_valid_id",
						Name:        "product_valid_name",
						Description: "product_valid_desc",
						Price:       10,
						Category:    "product_valid_category",
						Status:      0,
						ImagePath:   "product_valid_imgpath",
					}

					createdProduct, err := repo.Create(context.Background(), product)

					assert.NotNil(t, err)
					assert.Nil(t, createdProduct)

				},
			},
		},
	}

	for _, tc := range tests {
		db := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
		db.Run("", tc.given.mtestFunc)
	}
}

func TestUpdate(t *testing.T) {

	mpatch.PatchMethod(time.Now, func() time.Time {
		return time.Date(2020, 11, 01, 00, 00, 00, 0, time.UTC)
	})

	type Given struct {
		mtestFunc func(mt *mtest.T)
	}
	type Expected struct {
	}
	tests := map[string]struct {
		given    Given
		expected Expected
	}{
		"given given no error updating must return no error": {
			given: Given{
				mtestFunc: func(mt *mtest.T) {
					repo := repository.NewProductRepo(mt.DB)
					mt.AddMockResponses(bson.D{
						{Key: "ok", Value: 1},
						{Key: "value", Value: bson.D{
							{Key: "_id", Value: "product_valid_id"},
							{Key: "name", Value: "product_valid_name"},
							{Key: "description", Value: "product_valid_desc"},
							{Key: "price", Value: 10.0},
							{Key: "category", Value: "product_valid_category"},
							{Key: "status", Value: 0},
							{Key: "image_path", Value: "product_valid_imgpath"},
						}},
					})

					product := canonical.Product{
						ID:          "product_valid_id",
						Name:        "product_valid_name",
						Description: "product_valid_desc",
						Price:       10,
						Category:    "product_valid_category",
						Status:      0,
						ImagePath:   "product_valid_imgpath",
					}

					err := repo.Update(context.Background(), "product_valid", product)

					assert.Nil(t, err)

				},
			},
		},
		"given error saving must return error": {
			given: Given{
				mtestFunc: func(mt *mtest.T) {
					repo := repository.NewProductRepo(mt.DB)
					mt.AddMockResponses(
						bson.D{
							{Key: "ok", Value: -1},
						},
					)
					product := canonical.Product{
						ID:          "",
						Name:        "",
						Description: "",
						Price:       0,
						Category:    "",
						Status:      0,
						ImagePath:   "",
					}

					err := repo.Update(context.Background(), "product_valid", product)

					assert.NotNil(t, err)

				},
			},
		},
	}

	for _, tc := range tests {
		db := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
		db.Run("", tc.given.mtestFunc)
	}
}
