package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	greetpb "../greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("from client ...")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	//cfmt.Printf("created client: %f", c)

	// doUnary(c)
	// doServerStreaming(c)
	doClientStreaming(c)

} // .main

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting client streaming RPC...")

	req := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Jack",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Jill",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "John",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Piper",
			},
		},
	} // .req

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongGreet: %v", err)
	}

	for _, req := range req {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(500 * time.Millisecond)
	} // .for

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v", err)
	}
	fmt.Printf("LongGreet Response: %v\n", res)
} // .doClientStreaming

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting server streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Android",
			LastName:  "Apple",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()

		if err == io.EOF {
			// end of stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	} // .for
} // .doServerStreaming

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting an Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Andy",
			LastName:  "McDonnald",
		},
	} // .req
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet RPC: %v", res.Result)

} // .doUnary
