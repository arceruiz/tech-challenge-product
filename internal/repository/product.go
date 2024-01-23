package repository

import (
	"context"
	"tech-challenge-product/internal/canonical"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	productCollection = "product"
)

type ProductRepository interface {
	GetAll(context.Context) ([]canonical.Product, error)
	Create(ctx context.Context, product *canonical.Product) (*canonical.Product, error)
	Update(context.Context, string, canonical.Product) error
	GetByID(context.Context, string) (*canonical.Product, error)
	GetByCategory(context.Context, string) ([]canonical.Product, error)
}

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepo(db *mongo.Database) ProductRepository {
	return &productRepository{collection: db.Collection(productCollection)}
}

func (r *productRepository) GetAll(ctx context.Context) ([]canonical.Product, error) {
	filter := bson.D{{Key: "status", Value: 0}}
	cursor, err := r.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	var results []canonical.Product
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *productRepository) Create(ctx context.Context, product *canonical.Product) (*canonical.Product, error) {
	_, err := r.collection.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) Update(ctx context.Context, id string, product canonical.Product) error {
	filter := bson.M{"_id": id}
	fields := bson.M{"$set": product}

	_, err := r.collection.UpdateOne(ctx, filter, fields)
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepository) GetByID(ctx context.Context, id string) (*canonical.Product, error) {

	var roduct canonical.Product

	err := r.collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&roduct)
	if err != nil {
		return nil, err
	}

	return &roduct, nil
}

func (r *productRepository) GetByCategory(ctx context.Context, category string) ([]canonical.Product, error) {
	filter := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{{Key: "category", Value: category}},
				bson.D{{Key: "status", Value: 0}},
			},
		},
	}
	cursor, err := r.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	var results []canonical.Product
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil
}
