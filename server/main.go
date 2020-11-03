package main

import (
	"context"
	"fmt"
	"log"
	"net"

	_ "time/tzdata"

	"google.golang.org/grpc"
	"h2o3.com/demo/protobuf"
)

type helloServer struct {
}

func (h *helloServer) SayHello(ctx context.Context, hello *protobuf.Hello) (*protobuf.Message, error) {
	log.Printf("saying hello to %s", hello.Name)

	return &protobuf.Message{
		Message: fmt.Sprintf("hello, %s.", hello.Name),
	}, nil
}

func main() {
	server := grpc.NewServer()

	protobuf.RegisterHelloServiceServer(server, &helloServer{})

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on addr :9000, err = %+v", err)
	}

	log.Println("starting gRPC server.")

	log.Fatal(server.Serve(listener))
}
