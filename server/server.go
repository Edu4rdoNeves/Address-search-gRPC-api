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

	lis, err := net.Listen("tcp", os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listening on Port: %s", os.Getenv("SERVER_PORT"))

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to inicialize servidor. Error: %v", err)
	}

	return nil
}
