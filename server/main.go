package main

import (
	"context"
	trippb "godemo/proto/gen/go"
	trip "godemo/tripservice"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.Lshortfile)
	go startGPRCGateway()
	lis, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	trippb.RegisterTripServiceServer(s, &trip.Service{})
	log.Fatal(s.Serve(lis))
}

func startGPRCGateway() {
	log.SetFlags(log.Lshortfile)
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()
	mux := runtime.NewServeMux()
	err := trippb.RegisterTripServiceHandlerFromEndpoint(c, mux, "localhost:8081",
		[]grpc.DialOption{grpc.WithInsecure()})

	if err != nil {
		log.Fatalf("cannot start grpc gateway: %v", err)
	}

	err = http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatalf("cannot listen and server: %v", err)
	}
}
