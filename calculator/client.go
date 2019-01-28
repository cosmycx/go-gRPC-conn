package main

import (
	"context"
	"fmt"
	"log"

	calc_pb "./calculator_pb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("from client ...")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := calc_pb.NewCalculatorServiceClient(conn)

	rpcCalcSum(c, int32(10), int32(15))

	// rpcStreamFiboNums(c, int32(30))

} // .main

func rpcCalcSum(cc calc_pb.CalculatorServiceClient, a int32, b int32) int32 {
	fmt.Println("Starting sum request ...")

	req := &calc_pb.SumRequest{
		Sum: &calc_pb.Sum{
			A: a,
			B: b,
		},
	} // .req
	res, err := cc.SumService(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}

	log.Printf("Sum: %v", res.Result)
	return res.Result
} // .rpcCalcSum
