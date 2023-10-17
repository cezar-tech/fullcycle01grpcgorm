package server

import (
	"context"
	"log"
	"net"

	db "github.com/cezar-tech/fullcycle01grpcgorm/db"
	fullcycle01grpcgorm "github.com/cezar-tech/fullcycle01grpcgorm/proto"
	"github.com/cezar-tech/fullcycle01grpcgorm/repository"

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
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
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
