package main

import (
	"context"
	"fmt"
	trippb "godemo/proto/gen/go"
	"log"

	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.Lshortfile)

	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("cannot connect server: %v", err)
	}

	tsClient := trippb.NewTripServiceClient(conn)
	r, err := tsClient.GetTrip(context.Background(), &trippb.GetTripRequest{
		Id: "trip456",
	})

	if err != nil {
		log.Fatalf("cannot call GetTrip: %v", err)
	}
	fmt.Println("++++++")
	fmt.Println(r)

}
