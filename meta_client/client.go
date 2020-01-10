package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "github.com/bruteforce1414/testGRPC/metainfo"
)

const (
	address        = "localhost:50051"
	defaultID      = "nori/session"
	defaultVersion = "1.0.0"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMetaInfoClient(conn)

	// Contact the server and print out its response.
	id := defaultID

	if len(os.Args) > 1 {
		id = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetMetaInfo(ctx, &pb.MetaDataRequest{
		Id:      id,
		Version: defaultVersion,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Plugin: %s", r)
}
