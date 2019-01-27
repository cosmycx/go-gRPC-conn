package main

import (
	"context"
	"fmt"
	"log"
	"net"

	calc_pb "./calculator_pb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) SumService(ctx context.Context, req *calc_pb.SumRequest) (*calc_pb.SumResponse, error) {
	fmt.Printf("Sum service func invoked with request: %v", req)

	a := req.GetSum().GetA()
	b := req.GetSum().GetB()

	result := a + b
	resp := &calc_pb.SumResponse{
		Result: result,
	}
	return resp, nil
} // .Greet

func main() {

	fmt.Println("Starting RPC server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	calc_pb.RegisterSumServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
