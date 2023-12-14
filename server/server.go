package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"encoding/json"
	"github.com/epsilondylan/blockchain/common"
	"github.com/epsilondylan/blockchain/models"
	pto "github.com/epsilondylan/blockchain/protocal"
	create "github.com/epsilondylan/blockchain/server/create"
	"github.com/epsilondylan/service"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)


var wg sync.WaitGroup

type blockchainServer struct {
	service.UnsafeBlockchainServiceServer
}

func main() {
	p2pAddress := "127.0.0.1:12345"   // Replace with your actual P2P address
	grpcAddress := "127.0.0.1:8888" // Replace with your actual gRPC address

	// Initialize P2P and gRPC
	pto.InitPto(p2pAddress, common.P2PTimeOut)

	// Start gRPC server
	listen, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	service.RegisterBlockchainServiceServer(grpcServer, &blockchainServer{})
    wg.Add(1)
	go func() {
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	wg.Wait()
}

type CResponse struct {
	Errno int    `json:"errno"`
	Msg   string `json:"msg"`
	Data  string `json:"data"`
}

type CRequest struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func NewCResponseIDL() *CResponse {
	return &CResponse{}
}

func GenerateBlock(req *CRequest) *CResponse {
	resp := NewCResponseIDL()
	resp.Errno = common.Success
	resp.Msg = common.ErrMap[common.Success]
	pto.DataQueueAppend((*pto.CRequest)(req))
	return resp
}

func (s *blockchainServer) CreateBlock(ctx context.Context, request *service.CRequest) (*service.CResponse, error) {
	// Processing the request
	req := &CRequest{
		Name: request.Name,
		Data: request.Data,
	}
	// Call the controller to process the request
	resp := GenerateBlock(req)

	// Return the response
	return &service.CResponse{
		Errno: int32(resp.Errno),
		Msg:   resp.Msg,
		Data:  resp.Data,
	}, nil
}
type JResponse struct {
	Errno int    `json:"errno"`
	Msg   string `json:"msg"`
}

type JRequest struct {
	PeerAddr string `json:"peer_addr"`
}

func AddPeer(req *JRequest) *JResponse {
	resp := &JResponse{}
	resp.Errno = common.Success
	resp.Msg = common.ErrMap[common.Success]
	err := pto.AddPeer(req.PeerAddr)
	if err != nil {
		resp.Errno = common.JoinPeerFail
		resp.Msg = common.ErrMap[common.JoinPeerFail]
	}
	return resp
}

func (s *blockchainServer) JoinNode(ctx context.Context, request *service.JoinRequest) (*service.JResponse, error) {
	midrequest := &JRequest{
		PeerAddr: request.PeerAddr,
	}
	resp := AddPeer(midrequest)
	return &service.JResponse{
		Errno: int32(resp.Errno),
		Msg:   resp.Msg,
	}, nil
}
type SRequest struct {
	Chain	bool	`json:"chain"`
	Peer 	bool	`json:"peer"`
}

type SResponse struct {
	Chain  interface{}	`json:"chain"`
	Peer   interface{}	`json:"peer"`
}

// NewJResponse ...
func NewJResponse() *SResponse {
	return &SResponse{}
}

func Show(req *SRequest) *SResponse {
	resp := NewJResponse()
	single := pto.GetProtocal()
	if req.Chain {
		resp.Chain = models.FetchChain()
	}
	if req.Peer {
		resp.Peer = single.GetRouter().FetchPeers()
	}
	return resp
}

func (s *blockchainServer) ShowBlockchain(ctx context.Context, request *service.SRequest) (*service.SResponse, error) {
	midrequest := &SRequest{
		Chain: request.Chain,
		Peer:  request.Peer,
	}
	resp := Show(midrequest)

	// Handle the correct type conversion based on the actual types returned by Show
	var chainResp interface{}
	var peerResp interface{}

	if request.Chain {
		chainResp = resp.Chain // Update this based on the actual type returned by Show
	}
	if request.Peer {
		peerResp = resp.Peer // Update this based on the actual type returned by Show
	}
    rspchain, err := json.Marshal(chainResp)
	if err != nil {	
		log.Fatalf("failed to marshal: %v", err)
	}

	rspperr, err := json.Marshal(peerResp)
	if err != nil {
		log.Fatalf("failed to marshal: %v", err)
	}
	return &service.SResponse{
		Chain: rspchain,
		Peer:  rspperr,
	}, nil
}

