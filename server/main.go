package main

import (
	"first-grpc/api"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

// start a gRPC server and waits for connection
func main() {

	// create a listener on TCP port 7777 to bind the gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a server instance
	s := api.Server{}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the Ping service to the server, register the serve
	api.RegisterPingServer(grpcServer, &s)

	// start the gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
