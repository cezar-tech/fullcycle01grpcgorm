package server

import (
	"context"
	"fmt"
	"log"
	"net"

	db "github.com/cezar-tech/fullcycle01grpcgorm/go/src/db"
	"github.com/cezar-tech/fullcycle01grpcgorm/go/src/repository"
	fullcycle01grpcgorm "github.com/cezar-tech/fullcycle01grpcgorm/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	repo *repository.ProductRepository
	fullcycle01grpcgorm.UnimplementedProductServiceServer
}

func (s *Server) CreateProduct(ctx context.Context, in *fullcycle01grpcgorm.CreateProductRequest) (*fullcycle01grpcgorm.CreateProductResponse, error) {
	return s.repo.CreateProduct(ctx, in)
}

func (s *Server) FindProducts(ctx context.Context, in *fullcycle01grpcgorm.FindProductsRequest) (*fullcycle01grpcgorm.FindProductsResponse, error) {
	return s.repo.FindProducts(ctx, in)
}

func Start() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("Listening on port %v", address)
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	server := &Server{
		repo: &repository.ProductRepository{
			Db: db.ConnectDB("dev"),
		},
	}

	reflection.Register(grpcServer)
	fullcycle01grpcgorm.RegisterProductServiceServer(grpcServer, server)
	grpcServer.Serve(lis)
}
