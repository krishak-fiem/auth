package main

import (
	"github.com/krishak-fiem/auth/proto/pb"
	"github.com/krishak-fiem/auth/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Failed to start the auth service on port 5000: %v\n", err)
	}

	grpcServer := grpc.NewServer()

	auth := service.Server{}
	pb.RegisterAuthServiceServer(grpcServer, &auth)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start the auth service on port 5000: %v\n", err)
	}
}
