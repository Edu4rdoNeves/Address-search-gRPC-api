package server

import (
	"log"
	"net"
	"os"

	"github.com/Edu4rdoNeves/Address-search-gRPC-api/business"
	"github.com/Edu4rdoNeves/Address-search-gRPC-api/pb"
	"github.com/Edu4rdoNeves/Address-search-gRPC-api/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func GrpcServer() error {
	business := business.NewAddressBusiness()
	addressService := service.NewAddressService(business)

	grpcServer := grpc.NewServer()
	pb.RegisterAddressServicesServer(grpcServer, addressService)
	reflection.Register(grpcServer)

	port := os.Getenv("SERVER_PORT")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listening on Port: %s", port)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to inicialize servidor. Error: %v", err)
	}
	return nil
}
