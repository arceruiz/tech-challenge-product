package main

import (
	"tech-challenge-product/internal/channels/grpc"
	"tech-challenge-product/internal/channels/rest"
	"tech-challenge-product/internal/config"

	"github.com/sirupsen/logrus"
)

func main() {
	config.ParseFromFlags()
	go func() {
		logrus.Fatal(grpc.Listen())
	}()

	if err := rest.New(rest.NewProductChannel()).Start(); err != nil {
		logrus.Panic()
	}
}
