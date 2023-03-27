package main

import (
	"context"
	"io"
	"log"

	p "grpc_streaming/client/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can't connect with server: %v", err)
	}
	client := p.NewStreamServiceClient(conn)
	in := &p.Request{Id: 1}
	stream, err := client.FetchResponse(context.Background(), in)
	if err != nil {
		log.Fatalf("open stream error: %v", err)
	}
	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				log.Fatalf("can't receivce: %v", err)
			}
			log.Printf("Resp received: %s", resp.Result)
		}
	}()

	<-done
	log.Printf("finished")
}
