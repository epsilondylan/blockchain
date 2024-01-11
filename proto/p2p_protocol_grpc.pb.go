// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: proto/p2p_protocol.proto

package p2p_proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	P2P_RequestTailBlock_FullMethodName = "/p2p.P2P/RequestTailBlock"
	P2P_UpdateBlockChain_FullMethodName = "/p2p.P2P/UpdateBlockChain"
	P2P_NewBlock_FullMethodName         = "/p2p.P2P/NewBlock"
	P2P_NewTransaction_FullMethodName   = "/p2p.P2P/NewTransaction"
	P2P_NewPeer_FullMethodName          = "/p2p.P2P/NewPeer"
)

// P2PClient is the client API for P2P service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type P2PClient interface {
	RequestTailBlock(ctx context.Context, in *TailBlockRequest, opts ...grpc.CallOption) (*Block, error)
	UpdateBlockChain(ctx context.Context, in *BlockChain, opts ...grpc.CallOption) (*UpdateBlockChainResponse, error)
	NewBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*NewBlockResponse, error)
	NewTransaction(ctx context.Context, in *Trans, opts ...grpc.CallOption) (*NewTransactionResponse, error)
	NewPeer(ctx context.Context, in *Peer, opts ...grpc.CallOption) (*PeerList, error)
}

type p2PClient struct {
	cc grpc.ClientConnInterface
}

func NewP2PClient(cc grpc.ClientConnInterface) P2PClient {
	return &p2PClient{cc}
}

func (c *p2PClient) RequestTailBlock(ctx context.Context, in *TailBlockRequest, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := c.cc.Invoke(ctx, P2P_RequestTailBlock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2PClient) UpdateBlockChain(ctx context.Context, in *BlockChain, opts ...grpc.CallOption) (*UpdateBlockChainResponse, error) {
	out := new(UpdateBlockChainResponse)
	err := c.cc.Invoke(ctx, P2P_UpdateBlockChain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2PClient) NewBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*NewBlockResponse, error) {
	out := new(NewBlockResponse)
	err := c.cc.Invoke(ctx, P2P_NewBlock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2PClient) NewTransaction(ctx context.Context, in *Trans, opts ...grpc.CallOption) (*NewTransactionResponse, error) {
	out := new(NewTransactionResponse)
	err := c.cc.Invoke(ctx, P2P_NewTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2PClient) NewPeer(ctx context.Context, in *Peer, opts ...grpc.CallOption) (*PeerList, error) {
	out := new(PeerList)
	err := c.cc.Invoke(ctx, P2P_NewPeer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// P2PServer is the server API for P2P service.
// All implementations must embed UnimplementedP2PServer
// for forward compatibility
type P2PServer interface {
	RequestTailBlock(context.Context, *TailBlockRequest) (*Block, error)
	UpdateBlockChain(context.Context, *BlockChain) (*UpdateBlockChainResponse, error)
	NewBlock(context.Context, *Block) (*NewBlockResponse, error)
	NewTransaction(context.Context, *Trans) (*NewTransactionResponse, error)
	NewPeer(context.Context, *Peer) (*PeerList, error)
	mustEmbedUnimplementedP2PServer()
}

// UnimplementedP2PServer must be embedded to have forward compatible implementations.
type UnimplementedP2PServer struct {
}

func (UnimplementedP2PServer) RequestTailBlock(context.Context, *TailBlockRequest) (*Block, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestTailBlock not implemented")
}
func (UnimplementedP2PServer) UpdateBlockChain(context.Context, *BlockChain) (*UpdateBlockChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlockChain not implemented")
}
func (UnimplementedP2PServer) NewBlock(context.Context, *Block) (*NewBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewBlock not implemented")
}
func (UnimplementedP2PServer) NewTransaction(context.Context, *Trans) (*NewTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewTransaction not implemented")
}
func (UnimplementedP2PServer) NewPeer(context.Context, *Peer) (*PeerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewPeer not implemented")
}
func (UnimplementedP2PServer) mustEmbedUnimplementedP2PServer() {}

// UnsafeP2PServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to P2PServer will
// result in compilation errors.
type UnsafeP2PServer interface {
	mustEmbedUnimplementedP2PServer()
}

func RegisterP2PServer(s grpc.ServiceRegistrar, srv P2PServer) {
	s.RegisterService(&P2P_ServiceDesc, srv)
}

func _P2P_RequestTailBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TailBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PServer).RequestTailBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: P2P_RequestTailBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PServer).RequestTailBlock(ctx, req.(*TailBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2P_UpdateBlockChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockChain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PServer).UpdateBlockChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: P2P_UpdateBlockChain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PServer).UpdateBlockChain(ctx, req.(*BlockChain))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2P_NewBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Block)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PServer).NewBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: P2P_NewBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PServer).NewBlock(ctx, req.(*Block))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2P_NewTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Trans)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PServer).NewTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: P2P_NewTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PServer).NewTransaction(ctx, req.(*Trans))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2P_NewPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Peer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PServer).NewPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: P2P_NewPeer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PServer).NewPeer(ctx, req.(*Peer))
	}
	return interceptor(ctx, in, info, handler)
}

// P2P_ServiceDesc is the grpc.ServiceDesc for P2P service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var P2P_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "p2p.P2P",
	HandlerType: (*P2PServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestTailBlock",
			Handler:    _P2P_RequestTailBlock_Handler,
		},
		{
			MethodName: "UpdateBlockChain",
			Handler:    _P2P_UpdateBlockChain_Handler,
		},
		{
			MethodName: "NewBlock",
			Handler:    _P2P_NewBlock_Handler,
		},
		{
			MethodName: "NewTransaction",
			Handler:    _P2P_NewTransaction_Handler,
		},
		{
			MethodName: "NewPeer",
			Handler:    _P2P_NewPeer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/p2p_protocol.proto",
}
