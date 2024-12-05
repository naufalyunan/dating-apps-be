package clients

import (
	pb "api-gateway/pb/generated"
	"log"
	"os"

	"google.golang.org/grpc"
)

func NewProfileClient() pb.ProfileServiceClient {
	addr := os.Getenv("PROFILE_SERVICE_URL")
	log.Printf("profile service url: %s", addr)
	// Set up a connection to the server.
	// opts := []grpc.DialOption{}
	// systemRoots, err := x509.SystemCertPool()
	// if err != nil {
	// 	log.Fatalf("filed to get certs: %v", err)
	// }
	// cred := credentials.NewTLS(&tls.Config{
	// 	RootCAs: systemRoots,
	// })
	// opts = append(opts, grpc.WithTransportCredentials(cred))
	// conn, err := grpc.NewClient(addr, opts...)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := pb.NewProfileServiceClient(conn)

	return client
}
