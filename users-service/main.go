package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"users-service/configs"
	"users-service/handlers"
	pb "users-service/pb/generated"
	"users-service/services"

	"google.golang.org/grpc"
)

func main() {
	db := configs.CreateDBInstance()

	//instantiate services
	logService := services.NewLogService()
	userHandler := handlers.NewUserHandler(db, logService)

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, userHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "50001"
	}

	// start server
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %s", listen.Addr().String())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
