package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "go-grpc/helloworld/helloworld"
)

func main() {
	log.Println("inside client grpc")
	conn, err := grpc.NewClient("localhost:3000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Error in client %v", err)
	}

	defer conn.Close()
	client := pb.NewGreetingMessageClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.SayHello(ctx, &pb.HelloworldRequest{Name: "Sumeet"})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Greeting: %s", response.GetMessage())
}
