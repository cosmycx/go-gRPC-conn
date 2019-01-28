package main

import (
	"context"
	"fmt"
	"io"
	"log"

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

	doUnary(c)
	doServerStreaming(c)

} // .main

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting streaming RPC...")

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
