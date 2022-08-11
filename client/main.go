package main

import (
	"context"
	"first-grpc/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	var conn *grpc.ClientConn

	// instantiate a client connection, on the TCP port the server is bound to
	conn, err := grpc.Dial(":7777", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("error when closing the connection: %s", err)
		}
	}(conn)

	// create a client for the Ping service
	c := api.NewPingClient(conn)

	resp, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "foo"})
	if err != nil {
		log.Fatalf("error when calling SayHello: %s", err)
	}
	log.Printf("response from server: %s", resp.Greeting)
}
