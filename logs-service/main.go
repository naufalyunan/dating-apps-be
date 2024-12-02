package main

import (
	"fmt"
	"log"
	"logs-service/configs"
	"logs-service/handlers"
	pb "logs-service/pb/generated"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	db := configs.CreateDBInstance()

	//instantiate services
	logHandler := handlers.NewLogHandler(db)

	grpcServer := grpc.NewServer()

	pb.RegisterLogServiceServer(grpcServer, logHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "50002"
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
