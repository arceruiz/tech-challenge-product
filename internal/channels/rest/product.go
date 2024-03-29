package rest

import (
	"context"
	"encoding/json"
	"tech-challenge-product/internal/service"

	"net/http"

	"github.com/labstack/echo/v4"
)

type Product interface {
	RegisterGroup(g *echo.Group)
	Get(c echo.Context) error
	Add(c echo.Context) error
	Update(c echo.Context) error
	Remove(c echo.Context) error
	HealthCheck(c echo.Context) error
}

type productChannel struct {
	service service.ProductService
}

func NewProductChannel() Product {
	return &productChannel{
		service: service.NewProductService(),
	}
}

func (p *productChannel) RegisterGroup(g *echo.Group) {
	indexPath := "/"
	g.GET("", p.Get)
	g.GET(indexPath, p.Get)
	g.POST(indexPath, p.Add)
	g.PUT(indexPath+":id", p.Update)
	g.DELETE(indexPath+":id", p.Remove)
}

func (r *productChannel) HealthCheck(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (p *productChannel) Get(ctx echo.Context) error {
	productID := ctx.QueryParam("id")
	category := ctx.QueryParam("category")

	response, err := p.get(ctx.Request().Context(), productID, category)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	if len(response) == 0 {
		return ctx.JSON(http.StatusNotFound, nil)
	}
	if len(response) == 1 {
		return ctx.JSON(http.StatusOK, response[0])
	}
	return ctx.JSON(http.StatusOK, response)
}

func (p *productChannel) get(ctx context.Context, productID string, category string) ([]ProductResponse, error) {

	var response []ProductResponse

	if productID != "" {
		product, err := p.service.GetByID(ctx, productID)
		if err != nil {
			return nil, err
		}
		return []ProductResponse{productToResponse(product)}, nil
	}

	if category != "" {
		products, err := p.service.GetByCategory(ctx, category)
		if err != nil {
			return nil, err
		}

		for _, product := range products {
			response = append(response, productToResponse(&product))
		}
		return response, nil
	}

	products, err := p.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		response = append(response, productToResponse(&product))
	}

	return response, nil
}

func (p *productChannel) Add(c echo.Context) error {
	var newProduct ProductRequest
	err := c.Bind(&newProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	insertedProduct, err := p.service.Create(c.Request().Context(), newProduct.toCanonical())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, productToResponse(insertedProduct))
}

func (p *productChannel) Update(c echo.Context) error {
	productID := c.Param("id")

	var updatedProduct *ProductRequest
	err := json.NewDecoder(c.Request().Body).Decode(&updatedProduct)
	if err != nil || updatedProduct == nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	err = p.service.Update(c.Request().Context(), productID, *updatedProduct.toCanonical())
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	return c.JSON(http.StatusOK, nil)
}

func (p *productChannel) Remove(c echo.Context) error {
	productID := c.Param("id")

	err := p.service.Remove(c.Request().Context(), productID)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	return c.NoContent(http.StatusOK)
}
