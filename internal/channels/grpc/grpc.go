package grpc

import (
	"context"
	"net"
	"tech-challenge-product/internal/service"

	protocol "google.golang.org/grpc"
)

type productGRPCServer struct {
	service.ProductService
	UnimplementedProductServiceServer
}

func New() ProductServiceServer {
	return &productGRPCServer{
		ProductService: service.NewProductService(),
	}
}

func Listen(port int) error {
	server := protocol.NewServer()
	listener, err := net.Listen("tcp", ":3004")
	if err != nil {
		return err
	}

	RegisterProductServiceServer(server, New())

	return server.Serve(listener)
}

func (p *productGRPCServer) GetProduct(ctx context.Context, ids *Ids) (*Products, error) {
	products, err := p.ProductService.GetProductsWithId(context.Background(), ids.Ids)
	if err != nil {
		return nil, err
	}

	return toResult(products), nil
}
