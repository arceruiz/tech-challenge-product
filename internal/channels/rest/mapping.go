package rest

import "tech-challenge-product/internal/canonical"

func (p *ProductRequest) toCanonical() *canonical.Product {
	return &canonical.Product{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Category:    p.Category,
		Status:      canonical.BaseStatus(p.Status),
		ImagePath:   p.ImagePath,
	}
}

func productToResponse(p *canonical.Product) ProductResponse {
	return ProductResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Category:    p.Category,
		Status:      int(p.Status),
		ImagePath:   p.ImagePath,
	}
}
