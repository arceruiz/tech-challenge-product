package main

import (
	"tech-challenge-product/internal/channels/rest"
	"tech-challenge-product/internal/config"
	"tech-challenge-product/internal/repository"
	"tech-challenge-product/internal/service"

	"github.com/sirupsen/logrus"
)

var (
	cfg = &config.Cfg
)

func main() {
	config.ParseFromFlags()
	restChannel := rest.NewProductChannel(service.NewProductService(repository.NewProductRepo(repository.NewMongo())))
	if err := rest.New(restChannel).Start(); err != nil {
		logrus.Panic()
	}

}
