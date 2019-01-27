package main

import (
	"context"
	"fmt"
	"log"
	"net"

	greetpb "../greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function invoked with request: %v", req)
	fn := req.GetGreeting().GetFirstName()
	result := "Hello " + fn
	resp := &greetpb.GreetResponse{
		Result: result,
	}
	return resp, nil
} // .Greet

func main() {
	fmt.Println("Hello, starting RPC server...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

} // .main
