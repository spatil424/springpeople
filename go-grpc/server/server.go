package main

import (
	"context"
	pb "go-grpc/helloworld/helloworld"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreetingMessageServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloworldRequest) (*pb.HelloworldResponse, error) {
	log.Printf("Received %v", in.GetName())
	return &pb.HelloworldResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	log.Println("inside main")

	s := grpc.NewServer()
	pb.RegisterGreetingMessageServer(s, &server{})
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("got following error %v", err)
	}

	log.Printf("Listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
