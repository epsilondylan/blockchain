package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/epsilondylan/blockchain/common"
	"github.com/epsilondylan/blockchain/models"
	pto "github.com/epsilondylan/blockchain/protocal"
	create "github.com/epsilondylan/blockchain/server/create"
	"github.com/epsilondylan/service"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

var wg sync.WaitGroup

type blockchainServer struct{}

func main() {
	p2pAddress := "p2p_address"   // Replace with your actual P2P address
	grpcAddress := "grpc_address" // Replace with your actual gRPC address

	// Initialize P2P and gRPC
	pto.InitPto(p2pAddress, common.P2PTimeOut)

	// Start gRPC server
	listen, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	service.RegisterBlockchainServiceServer(grpcServer, &blockchainServer{})

	go func() {
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	wg.Wait()
}

func (s *blockchainServer) CreateBlock(ctx context.Context, request *service.CRequest) (*service.CResponse, error) {
	// Processing the request
	idl := create.GenIdl()
	// Assuming request.Data is a serialized protobuf message, decode it into idl
	err := proto.Unmarshal([]byte(request.Data), idl)
	if err != nil {
		fmt.Println("Error decoding request data:", err)
		return nil, err
	}

	// Call the controller to process the request
	resp := create.Do(idl)

	// Assuming resp is a protobuf message, serialize it into bytes
	respData, err := proto.Marshal(resp)
	if err != nil {
		fmt.Println("Error encoding response data:", err)
		return nil, err
	}

	// Return the response
	return &service.CResponse{
		Errno: 0,
		Msg:   "OK",
		Data:  respData,
	}, nil
}

func (s *blockchainServer) JoinNode(ctx context.Context, request *service.JoinRequest) (*service.JResponse, error) {
	resp := &service.JResponse{
		Errno: common.Success,
		Msg:   common.ErrMap[common.Success],
	}
	err := pto.AddPeer(request.PeerAddr)
	if err != nil {
		resp.Errno = common.JoinPeerFail
		resp.Msg = common.ErrMap[common.JoinPeerFail]
	}
	return resp, nil
}

func (s *blockchainServer) ShowBlockchain(ctx context.Context, request *service.SRequest) (*service.SResponse, error) {
	resp := &service.SResponse{}
	single := pto.GetProtocal()
	if request.Chain {
		resp.Chain = models.FetchChain()
	}
	if request.Peer {
		resp.Peer = single.GetRouter().FetchPeers()
	}
	return resp, nil
}
