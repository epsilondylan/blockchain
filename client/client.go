package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	service "github.com/epsilondylan/service"
)

func main() {
	serverAddr := "127.0.0.1:8888"
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := service.NewBlockchainServiceClient(conn)

	// Example: Make gRPC requests
	createBlockRequest := &service.CRequest{Name: "luda2", Data: "second blockchain"}
	createBlockResponse, err := client.CreateBlock(context.Background(), createBlockRequest)
	if err != nil {
		log.Fatalf("Failed to create block: %v", err)
	}
	log.Printf("Create Block Response: %v", createBlockResponse)

	// Add more gRPC requests as needed...
     
	joinNodeRequest := &service.JoinRequest{
		PeerAddr: "127.0.0.1:12345", // Replace with the target peer's address
	}

	joinNodeResponse, err := client.JoinNode(context.Background(), joinNodeRequest)
	if err != nil {
		log.Fatalf("Failed to join peer: %v", err)
	}
	log.Printf("Join Peer Response: %v", joinNodeResponse)

	Sreq := &service.SRequest{
		Chain: true,
		Peer:  true,
	}

	showBlockchainResponse, err := client.ShowBlockchain(context.Background(), Sreq)
	if err != nil {
		log.Fatalf("Failed to show blockchain: %v", err)
	}
	log.Printf("Show Blockchain Response: %+v", showBlockchainResponse)
	// Remember to handle errors appropriately in a real application
}