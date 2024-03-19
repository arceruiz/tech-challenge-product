package grpc

import (
	"context"
	"log"
	"net"
	"tech-challenge-product/internal/canonical"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	"github.com/stretchr/testify/assert"
)

var (
	mockS = ProductServiceMock{}
)

var buff int = 10 * 1024

func server() (ProductServiceClient, func()) {
	lis := bufconn.Listen(buff)

	server := grpc.NewServer()

	RegisterProductServiceServer(server, &productGRPCServer{
		ProductService: &mockS,
	})

	go func() {
		if err := server.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(context.Background(), "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		server.Stop()
	}

	client := NewProductServiceClient(conn)

	return client, closer
}

func TestGetProduct(t *testing.T) {
	mockS.On("GetProductsWithId", []string{
		"123", "456",
	}).Return([]canonical.Product{
		{
			ID:          "123",
			Name:        "test",
			Description: "desc",
			Price:       123,
			Category:    "cat",
			Status:      canonical.STATUS_ACTIVE,
			ImagePath:   "path",
		},
		{
			ID:          "456",
			Name:        "test",
			Description: "desc",
			Price:       123,
			Category:    "cat",
			Status:      canonical.STATUS_ACTIVE,
			ImagePath:   "path",
		},
	}, nil)

	server, f := server()

	defer f()

	products, err := server.GetProduct(context.Background(), &Ids{
		Ids: []string{
			"123", "456",
		},
	})

	assert.Nil(t, err)
	assert.NotNil(t, products)
}
