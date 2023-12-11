package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

<<<<<<< HEAD
	"github.com/epsilondylan/blockchain/common"
	"github.com/epsilondylan/blockchain/models"
	pto "github.com/epsilondylan/blockchain/protocal"
	create "github.com/epsilondylan/blockchain/server/create"
	"github.com/epsilondylan/service"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
=======
	"google.golang.org/grpc"
	pto "github.com/epsilondylan/blockchain/protocal"
	"github.com/epsilondylan/service"
	"github.com/golang/protobuf/proto"
>>>>>>> 4770628777744808942dfe386f04d0f1853a991d
)

var wg sync.WaitGroup

<<<<<<< HEAD
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
=======
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

>>>>>>> 4770628777744808942dfe386f04d0f1853a991d
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	wg.Wait()
}

func (s *blockchainServer) CreateBlock(ctx context.Context, request *service.CRequest) (*service.CResponse, error) {
	// Processing the request
<<<<<<< HEAD
	idl := create.GenIdl()
=======
	idl := ctrl.GenIdl()
>>>>>>> 4770628777744808942dfe386f04d0f1853a991d
	// Assuming request.Data is a serialized protobuf message, decode it into idl
	err := proto.Unmarshal([]byte(request.Data), idl)
	if err != nil {
		fmt.Println("Error decoding request data:", err)
		return nil, err
	}

	// Call the controller to process the request
<<<<<<< HEAD
	resp := create.Do(idl)
=======
	resp := ctrl.Do(idl)
>>>>>>> 4770628777744808942dfe386f04d0f1853a991d

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
<<<<<<< HEAD
	resp := &service.JResponse{
		Errno: common.Success,
		Msg:   common.ErrMap[common.Success],
	}
=======
	resp := service.NewJResponse()
	resp.Errno = common.Success
	resp.Msg = common.ErrMap[common.Success]
>>>>>>> 4770628777744808942dfe386f04d0f1853a991d
	err := pto.AddPeer(request.PeerAddr)
	if err != nil {
		resp.Errno = common.JoinPeerFail
		resp.Msg = common.ErrMap[common.JoinPeerFail]
	}
	return resp, nil
}

func (s *blockchainServer) ShowBlockchain(ctx context.Context, request *service.SRequest) (*service.SResponse, error) {
<<<<<<< HEAD
	resp := &service.SResponse{}
=======
	resp := service.NewSResponse()
>>>>>>> 4770628777744808942dfe386f04d0f1853a991d
	single := pto.GetProtocal()
	if request.Chain {
		resp.Chain = models.FetchChain()
	}
	if request.Peer {
		resp.Peer = single.GetRouter().FetchPeers()
	}
	return resp, nil
}
