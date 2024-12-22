package main

import (
	"fmt"
	"log"
	"net"

	pb "cinema/api/gen/cinema"
	cinemaservice "cinema/internal/cinema-service"
	"cinema/internal/config"
	"cinema/pkg/mpostgres"

	"google.golang.org/grpc"
)

func main() {
	// Load global configuration
	cfg, err := config.LoadConfig("config/", "cinema")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Connect to the database
	_, close := mpostgres.Connect(cfg.Postgres.ConnectionString, cfg.Postgres.MaxIdleConns, cfg.Postgres.MaxOpenConns)
	defer close()

	// Create gRPC server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Service.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	cinemaService := cinemaservice.New()
	pb.RegisterCinemaServiceServer(grpcServer, cinemaService)

	log.Printf("Server is running on port %d...\n", cfg.Service.Port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
