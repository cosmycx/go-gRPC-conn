package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

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

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {

	fmt.Printf("Greet Many Times function invoked with request: %v", req)
	fn := req.GetGreeting().GetFirstName()

	for i := 0; i < 10; i++ {
		result := "Hello " + fn + ", number: " + strconv.Itoa(i)
		resp := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(resp)
		time.Sleep(500 * time.Millisecond)
	} // .for

	return nil
} // .GreetManyTimes

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Println("Long Greet has been invoked with streaming request")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// finished reading the streaming client
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		fName := req.GetGreeting().GetFirstName()
		result += "Hello " + fName + "! "

	}
} // .LongGreet

func main() {
	fmt.Println("Hello, starting Greet RPC server...")

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
