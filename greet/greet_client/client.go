package main

import (
	"context"
	"fmt"
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

} // .main

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
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
