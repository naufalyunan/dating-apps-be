package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"profiles-service/configs"
	"profiles-service/handlers"
	pb "profiles-service/pb/generated"
	"profiles-service/services"

	"google.golang.org/grpc"
)

func main() {
	db := configs.CreateDBInstance()

	//instantiate services
	logService := services.NewLogService()
	userService := services.NewUserService()
	profileHandler := handlers.NewProfileHandler(db, userService, logService)

	grpcServer := grpc.NewServer()

	pb.RegisterProfileServiceServer(grpcServer, profileHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "50004"
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
