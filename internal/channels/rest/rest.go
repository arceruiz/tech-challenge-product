package rest

import (
	"tech-challenge-product/internal/config"
	"tech-challenge-product/internal/middlewares"

	"github.com/labstack/echo/v4"
)

var (
	cfg = &config.Cfg
)

type rest struct {
	product Product
}

func New(product Product) rest {
	return rest{
		product: product,
	}
}

func (r rest) Start() error {
	router := echo.New()

	router.Use(middlewares.Logger)

	mainGroup := router.Group("/api")

	mainGroup.GET("/healthz", r.product.HealthCheck)
	productGroup := mainGroup.Group("/product")
	r.product.RegisterGroup(productGroup)
	//productGroup.Use(middlewares.Authorization)

	return router.Start(":" + cfg.Server.Port)
}
