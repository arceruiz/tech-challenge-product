package rest

import (
	"errors"
	"net/http"
	"tech-challenge-product/internal/canonical"
)

type Response struct {
	Message string `json:"message"`
}

type ProductResponse struct {
	ID          string  `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Category    string  `json:"category,omitempty"`
	Status      int     `json:"status"`
	ImagePath   string  `json:"image_path,omitempty"`
}

type ProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Status      int     `json:"status"`
	ImagePath   string  `json:"image_path"`
}

func HandleError(err error) int {
	if errors.Is(err, canonical.ErrorNotFound) {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
