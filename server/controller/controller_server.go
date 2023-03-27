package controller

import (
	"fmt"
	p "grpc_streaming/server/proto"
	"log"
	"sync"
	"time"
)

type ServerService struct {
	p.UnimplementedStreamServiceServer
}

func (s *ServerService) FetchResponse(in *p.Request, srv p.StreamService_FetchResponseServer) error {
	log.Printf("fetch Response for id : %d", in.Id)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(count int64) {
			defer wg.Done()
			time.Sleep(time.Duration(count) * time.Second)
			resp := p.Result{Result: fmt.Sprintf("Request#%d for Id:%d", count, in.Id)}
			if err := srv.Send(&resp); err != nil {
				log.Printf("error when sending response: %v", err)
			}
			log.Printf("finishing request number %d", count)
		}(int64(i))
	}
	wg.Wait()
	return nil
}
