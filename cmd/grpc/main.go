package main

import (
	"context"
	"database/sql"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/lib/pq"

	"clean-arch/internal/infra/database"
	"clean-arch/internal/infra/grpc/pb"
	usecase "clean-arch/internal/order"
)

type server struct {
	pb.UnimplementedOrderServiceServer
	service *usecase.ListOrderUseCase
}

func (s *server) ListOrders(ctx context.Context, in *pb.Empty) (*pb.OrderList, error) {
	orders, err := s.service.ListOrders()
	if err != nil {
		return nil, err
	}

	var grpcOrders []*pb.Order
	for _, order := range orders {
		grpcOrders = append(grpcOrders, &pb.Order{
			Id:          int32(order.ID),
			Description: order.Description,
			Price:       float32(order.Price),
			CreatedAt:   order.CreatedAt.String(),
		})
	}
	return &pb.OrderList{Orders: grpcOrders}, nil
}

func main() {
	connStr := "postgres://user:password@db:5432/orders_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	repo := database.NewOrderRepository(db)
	service := usecase.NewListOrderUseCase(repo)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, &server{service: service})
	reflection.Register(grpcServer)

	log.Println("gRPC server running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
