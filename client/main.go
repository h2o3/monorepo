package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"h2o3.com/demo/protobuf"
)

var (
	serverTarget string
)

func init() {
	flag.StringVar(&serverTarget, "server.target", os.Getenv("server.target"), "")

	flag.Parse()
}

func main() {
	conn, err := grpc.Dial(serverTarget, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial, err = %+v", err)
	}
	defer func() {
		_ = conn.Close()
	}()

	client := protobuf.NewHelloServiceClient(conn)

	for {
		message, err := client.SayHello(context.Background(), &protobuf.Hello{Name: "Steven"})
		if err != nil {
			log.Printf("failed to request say hello service, err = %+v", err)

			time.Sleep(3 * time.Second)
			continue
		}

		log.Printf("get message: %s", message.Message)

		time.Sleep(1 * time.Second)
	}
}
