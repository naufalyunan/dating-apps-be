package main

import (
	"date-service/configs"
	"date-service/handlers"
	pb "date-service/pb/generated"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	db := configs.CreateDBInstance()

	//instantiate services
	swipeHandler := handlers.NewSwipeHandler(db)

	grpcServer := grpc.NewServer()

	pb.RegisterSwipeServiceServer(grpcServer, swipeHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "50003"
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
