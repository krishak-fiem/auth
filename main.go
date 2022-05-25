package main    

import (
	"github.com/krishak-fiem/auth/proto/pb"
	"github.com/krishak-fiem/auth/service"
	"github.com/krishak-fiem/db/go/cassandra"
	"google.golang.org/grpc"
	"log"
	"net"
	"fmt"
)

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Failed to start the auth products on port 5000: %v\n", err)
	}

	grpcServer := grpc.NewServer()

	auth := service.Server{}
	pb.RegisterAuthServiceServer(grpcServer, &auth)

	cassandra.Init(9042)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start the auth products on port 5000: %v\n", err)
	}
	fmt.Println("Auth server started on port 5000")
}
        