package main

import (
	"grpc_streaming/server/controller"
	p "grpc_streaming/server/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	s := grpc.NewServer()
	p.RegisterStreamServiceServer(s, &controller.ServerService{})
	log.Println("Start server")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
