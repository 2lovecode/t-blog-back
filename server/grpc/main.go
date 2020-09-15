package grpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"t-blog-back/proto/search"
)

type SearchService struct {}

func (s *SearchService) Search(ctx context.Context, r *search.SearchRequest) (*search.SearchResponse, error) {
	return &search.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9001"

func Start() {
	server := grpc.NewServer()
	search.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", ":"+PORT)

	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	err = server.Serve(lis)

	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

}
