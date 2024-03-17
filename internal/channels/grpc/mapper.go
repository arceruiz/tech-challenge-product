package grpc

import (
	"fmt"
	"tech-challenge-product/internal/canonical"
)

func toResult(products []canonical.Product) *Products {
	var result []*Product

	for _, product := range products {
		productResult := Product{
			Id:       product.ID,
			Name:     product.Name,
			Price:    fmt.Sprintf("%.2f", product.Price),
			Category: product.Category,
		}

		result = append(result, &productResult)
	}

	return &Products{
		Products: result,
	}
}
