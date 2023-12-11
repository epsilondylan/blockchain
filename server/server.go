package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	pto "github.com/epsilondylan/blockchain/protocal"
	"github.com/epsilondylan/service"
	"github.com/golang/protobuf/proto"
)

var wg sync.WaitGroup

type blockchainServer struct {
	pto *protocal.Protocal
}

func main() {
	p2pAddress := "p2p_address" // Replace with your actual P2P address
	grpcAddress := "grpc_address" // Replace with your actual gRPC address

	// Initialize P2P and gRPC
	pto.InitPto(p2pAddress, grpcAddress, common.P2PTimeOut)

	// Start gRPC server
	go func() {
		listen, err := net.Listen("tcp", grpcAddress)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		ServiceServer := &blockchainServer{pto: pto}
		service.RegisterBlockchainServiceServer(grpcServer, ServiceServer)

		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	wg.Wait()
}

func (s *blockchainServer) CreateBlock(ctx context.Context, request *service.CRequest) (*service.CResponse, error) {
	// Processing the request
	idl := ctrl.GenIdl()
	// Assuming request.Data is a serialized protobuf message, decode it into idl
	err := proto.Unmarshal([]byte(request.Data), idl)
	if err != nil {
		fmt.Println("Error decoding request data:", err)
		return nil, err
	}

	// Call the controller to process the request
	resp := ctrl.Do(idl)

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
	resp := service.NewJResponse()
	resp.Errno = common.Success
	resp.Msg = common.ErrMap[common.Success]
	err := pto.AddPeer(request.PeerAddr)
	if err != nil {
		resp.Errno = common.JoinPeerFail
		resp.Msg = common.ErrMap[common.JoinPeerFail]
	}
	return resp, nil
}

func (s *blockchainServer) ShowBlockchain(ctx context.Context, request *service.SRequest) (*service.SResponse, error) {
	resp := service.NewSResponse()
	single := pto.GetProtocal()
	if request.Chain {
		resp.Chain = models.FetchChain()
	}
	if request.Peer {
		resp.Peer = single.GetRouter().FetchPeers()
	}
	return resp, nil
}
