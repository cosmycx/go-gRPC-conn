package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"sync"
	"time"

	calc_pb "../calculator_pb"
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

	// doSumServ(c, int32(10), int32(15))

	// doFibonacciServ(c, int32(15))

	// doMeanServ(c)

	doMaxServ(c)

} // .main

func doMaxServ(c calc_pb.CalculatorServiceClient) {

	stream, err := c.MaxService(context.Background())
	if err != nil {
		log.Fatalf("error while calling MaxService: %v", err)
	}

	// waitchan := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for i := 0; i < 15; i++ {

			rand.Seed(time.Now().UnixNano())
			num := rand.Intn(1000) - 500

			fmt.Printf("Sending number: %v\n", num)
			stream.Send(&calc_pb.MaxRequest{
				Num: int64(num),
			})

			time.Sleep(time.Second)
		} // .for
		stream.CloseSend()

	}() // .go func

	go func() {
		for {

			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			fmt.Printf("Received, max: %v\n", res.GetMax())
		} // .for

		// close(waitchan)
		wg.Done()
	}() // .go func

	// <-waitchan
	wg.Wait()
} // .doMaxServ

func doMeanServ(c calc_pb.CalculatorServiceClient) {
	fmt.Println("Starting client streaming of numbers...")

	//MeanService(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_MeanServiceClient, error)

	numbers := []int64{1, 3, 9, 91, 504, 1023}

	stream, err := c.MeanService(context.Background())
	if err != nil {
		log.Fatalf("error while calling MeanService: %v", err)
	}

	for _, number := range numbers {
		fmt.Printf("Sending number: %v\n", number)
		stream.Send(&calc_pb.MeanRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}

	fmt.Printf("The Mean is: %v\n", res.GetMean())

} // .doMeanServ

func doFibonacciServ(cc calc_pb.CalculatorServiceClient, endNum int32) {
	fmt.Println("Starting fibonaci request...")
	req := &calc_pb.FibonacciRequest{
		Num: endNum,
	}

	resStream, err := cc.FibonacciService(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Fibonacci RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from Fibonacci: %v", msg.GetResultNum())
	} // .for

} // .doFibonacciServ

func doSumServ(cc calc_pb.CalculatorServiceClient, a int32, b int32) int32 {
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
} // .doSumServ
