package main

import (
	"context"
	"log"
	"t-blog-back/proto/search"

	"google.golang.org/grpc"
)

// PORT 端口号
const PORT = "9001"

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

	defer conn.Close()

	client := search.NewSearchServiceClient(conn)

	resp, err := client.Search(context.Background(), &search.SearchRequest{Request: "gggggg"})

	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", resp.GetResponse())
}
