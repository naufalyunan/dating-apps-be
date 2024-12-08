package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"payment-service/configs"
	pb "payment-service/pb/generated"
	"payment-service/server"
	"payment-service/services"

	"google.golang.org/grpc"
)

func main() {
	db := configs.CreateDBInstance()

	// instantiate services
	invoiceService := services.NewInvoiceService()
	userService := services.NewUserService()
	logService := services.NewLogService()

	// subs-payments grpc server handler
	paymentServer := server.NewPaymentServer(
		db,
		invoiceService,
		userService,
		logService,
	)

	opts := []grpc.ServerOption{
		// The following grpc.ServerOption adds an interceptor for all unary
		// RPCs. To configure an interceptor for streaming RPCs, see:
		// https://godoc.org/google.golang.org/grpc#StreamInterceptor
		// grpc.UnaryInterceptor(jwtIntercept.EnsureValidToken),
	}
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSubPaymentServer(grpcServer, paymentServer)
	log.Printf("starting gRPC server on %s", port)
	log.Fatal(grpcServer.Serve(lis))
}
