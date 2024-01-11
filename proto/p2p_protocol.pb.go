// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.6.1
// source: proto/p2p_protocol.proto

package p2p_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PVHash    string `protobuf:"bytes,1,opt,name=PVHash,proto3" json:"PVHash,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Data      string `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	Index     int64  `protobuf:"varint,4,opt,name=Index,proto3" json:"Index,omitempty"`
	Nonce     int64  `protobuf:"varint,5,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	Hash      string `protobuf:"bytes,6,opt,name=Hash,proto3" json:"Hash,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_p2p_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p2p_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_proto_p2p_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetPVHash() string {
	if x != nil {
		return x.PVHash
	}
	return ""
}

func (x *Block) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Block) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Block) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Block) GetNonce() int64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *Block) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type Trans struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account     string `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`
	Cipher      string `protobuf:"bytes,2,opt,name=Cipher,proto3" json:"Cipher,omitempty"`
	Transaction string `protobuf:"bytes,3,opt,name=Transaction,proto3" json:"Transaction,omitempty"`
}

func (x *Trans) Reset() {
	*x = Trans{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_p2p_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trans) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trans) ProtoMessage() {}

func (x *Trans) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p2p_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trans.ProtoReflect.Descriptor instead.
func (*Trans) Descriptor() ([]byte, []int) {
	return file_proto_p2p_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *Trans) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *Trans) GetCipher() string {
	if x != nil {
		return x.Cipher
	}
	return ""
}

func (x *Trans) GetTransaction() string {
	if x != nil {
		return x.Transaction
	}
	return ""
}

type BlockChain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chain []*Block `protobuf:"bytes,1,rep,name=Chain,proto3" json:"Chain,omitempty"`
}

func (x *BlockChain) Reset() {
	*x = BlockChain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_p2p_protocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockChain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockChain) ProtoMessage() {}

func (x *BlockChain) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p2p_protocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockChain.ProtoReflect.Descriptor instead.
func (*BlockChain) Descriptor() ([]byte, []int) {
	return file_proto_p2p_protocol_proto_rawDescGZIP(), []int{2}
}

func (x *BlockChain) GetChain() []*Block {
	if x != nil {
		return x.Chain
	}
	return nil
}

type TailBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TailBlockRequest) Reset() {
	*x = TailBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_p2p_protocol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TailBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TailBlockRequest) ProtoMessage() {}

func (x *TailBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p2p_protocol_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TailBlockRequest.ProtoReflect.Descriptor instead.
func (*TailBlockRequest) Descriptor() ([]byte, []int) {
	return file_proto_p2p_protocol_proto_rawDescGZIP(), []int{3}
}

type UpdateBlockChainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateBlockChainResponse) Reset() {
	*x = UpdateBlockChainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_p2p_protocol_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBlockChainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlockChainResponse) ProtoMessage() {}

func (x *UpdateBlockChainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p2p_protocol_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlockChainResponse.ProtoReflect.Descriptor instead.
func (*UpdateBlockChainResponse) Descriptor() ([]byte, []int) {
	return file_proto_p2p_protocol_proto_rawDescGZIP(), []int{4}
}

type NewTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NewTransactionResponse) Reset() {
	*x = NewTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_p2p_protocol_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTransactionResponse) ProtoMessage() {}

func (x *NewTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p2p_protocol_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTransactionResponse.ProtoReflect.Descriptor instead.
func (*NewTransactionResponse) Descriptor() ([]byte, []int) {
	return file_proto_p2p_protocol_proto_rawDescGZIP(), []int{5}
}

type NewBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainNeedUpdate bool `protobuf:"varint,1,opt,name=ChainNeedUpdate,proto3" json:"ChainNeedUpdate,omitempty"`
}

func (x *NewBlockResponse) Reset() {
	*x = NewBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_p2p_protocol_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewBlockResponse) ProtoMessage() {}

func (x *NewBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p2p_protocol_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewBlockResponse.ProtoReflect.Descriptor instead.
func (*NewBlockResponse) Descriptor() ([]byte, []int) {
	return file_proto_p2p_protocol_proto_rawDescGZIP(), []int{6}
}

func (x *NewBlockResponse) GetChainNeedUpdate() bool {
	if x != nil {
		return x.ChainNeedUpdate
	}
	return false
}

type Peer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IP   string `protobuf:"bytes,1,opt,name=IP,proto3" json:"IP,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (x *Peer) Reset() {
	*x = Peer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_p2p_protocol_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peer) ProtoMessage() {}

func (x *Peer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p2p_protocol_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peer.ProtoReflect.Descriptor instead.
func (*Peer) Descriptor() ([]byte, []int) {
	return file_proto_p2p_protocol_proto_rawDescGZIP(), []int{7}
}

func (x *Peer) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *Peer) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type PeerList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peers []*Peer `protobuf:"bytes,1,rep,name=Peers,proto3" json:"Peers,omitempty"`
}

func (x *PeerList) Reset() {
	*x = PeerList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_p2p_protocol_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerList) ProtoMessage() {}

func (x *PeerList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p2p_protocol_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerList.ProtoReflect.Descriptor instead.
func (*PeerList) Descriptor() ([]byte, []int) {
	return file_proto_p2p_protocol_proto_rawDescGZIP(), []int{8}
}

func (x *PeerList) GetPeers() []*Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

type TransPool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransPool []*Trans `protobuf:"bytes,1,rep,name=TransPool,proto3" json:"TransPool,omitempty"`
}

func (x *TransPool) Reset() {
	*x = TransPool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_p2p_protocol_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransPool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransPool) ProtoMessage() {}

func (x *TransPool) ProtoReflect() protoreflect.Message {
	mi := &file_proto_p2p_protocol_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransPool.ProtoReflect.Descriptor instead.
func (*TransPool) Descriptor() ([]byte, []int) {
	return file_proto_p2p_protocol_proto_rawDescGZIP(), []int{9}
}

func (x *TransPool) GetTransPool() []*Trans {
	if x != nil {
		return x.TransPool
	}
	return nil
}

var File_proto_p2p_protocol_proto protoreflect.FileDescriptor

var file_proto_p2p_protocol_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x32, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x70, 0x32, 0x70, 0x22,
	0x91, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x56, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x56, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x4e, 0x6f, 0x6e,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48,
	0x61, 0x73, 0x68, 0x22, 0x5b, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x12, 0x20,
	0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x2e, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x20,
	0x0a, 0x05, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x32, 0x70, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x22, 0x12, 0x0a, 0x10, 0x54, 0x61, 0x69, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x1a, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x18, 0x0a, 0x16, 0x4e, 0x65, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3c, 0x0a, 0x10, 0x4e, 0x65,
	0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28,
	0x0a, 0x0f, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x65, 0x65, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x65,
	0x65, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x2a, 0x0a, 0x04, 0x50, 0x65, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x50,
	0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x50, 0x6f, 0x72, 0x74, 0x22, 0x2b, 0x0a, 0x08, 0x50, 0x65, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x05, 0x50, 0x65, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x05, 0x50, 0x65, 0x65, 0x72,
	0x73, 0x22, 0x35, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x28,
	0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x50, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x52, 0x09, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x50, 0x6f, 0x6f, 0x6c, 0x32, 0x99, 0x02, 0x0a, 0x03, 0x50, 0x32, 0x50,
	0x12, 0x37, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x61, 0x69, 0x6c, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x15, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x54, 0x61, 0x69, 0x6c, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x32,
	0x70, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x10, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x0f, 0x2e,
	0x70, 0x32, 0x70, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x1a, 0x1d,
	0x2e, 0x70, 0x32, 0x70, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x2f, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0a, 0x2e, 0x70, 0x32,
	0x70, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x15, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x4e, 0x65,
	0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3b, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0a, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x1a, 0x1b,
	0x2e, 0x70, 0x32, 0x70, 0x2e, 0x4e, 0x65, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x25, 0x0a,
	0x07, 0x4e, 0x65, 0x77, 0x50, 0x65, 0x65, 0x72, 0x12, 0x09, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x50,
	0x65, 0x65, 0x72, 0x1a, 0x0d, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x70, 0x32, 0x70, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_p2p_protocol_proto_rawDescOnce sync.Once
	file_proto_p2p_protocol_proto_rawDescData = file_proto_p2p_protocol_proto_rawDesc
)

func file_proto_p2p_protocol_proto_rawDescGZIP() []byte {
	file_proto_p2p_protocol_proto_rawDescOnce.Do(func() {
		file_proto_p2p_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_p2p_protocol_proto_rawDescData)
	})
	return file_proto_p2p_protocol_proto_rawDescData
}

var file_proto_p2p_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_p2p_protocol_proto_goTypes = []interface{}{
	(*Block)(nil),                    // 0: p2p.Block
	(*Trans)(nil),                    // 1: p2p.Trans
	(*BlockChain)(nil),               // 2: p2p.BlockChain
	(*TailBlockRequest)(nil),         // 3: p2p.TailBlockRequest
	(*UpdateBlockChainResponse)(nil), // 4: p2p.UpdateBlockChainResponse
	(*NewTransactionResponse)(nil),   // 5: p2p.NewTransactionResponse
	(*NewBlockResponse)(nil),         // 6: p2p.NewBlockResponse
	(*Peer)(nil),                     // 7: p2p.Peer
	(*PeerList)(nil),                 // 8: p2p.PeerList
	(*TransPool)(nil),                // 9: p2p.TransPool
}
var file_proto_p2p_protocol_proto_depIdxs = []int32{
	0, // 0: p2p.BlockChain.Chain:type_name -> p2p.Block
	7, // 1: p2p.PeerList.Peers:type_name -> p2p.Peer
	1, // 2: p2p.TransPool.TransPool:type_name -> p2p.Trans
	3, // 3: p2p.P2P.RequestTailBlock:input_type -> p2p.TailBlockRequest
	2, // 4: p2p.P2P.UpdateBlockChain:input_type -> p2p.BlockChain
	0, // 5: p2p.P2P.NewBlock:input_type -> p2p.Block
	1, // 6: p2p.P2P.NewTransaction:input_type -> p2p.Trans
	7, // 7: p2p.P2P.NewPeer:input_type -> p2p.Peer
	0, // 8: p2p.P2P.RequestTailBlock:output_type -> p2p.Block
	4, // 9: p2p.P2P.UpdateBlockChain:output_type -> p2p.UpdateBlockChainResponse
	6, // 10: p2p.P2P.NewBlock:output_type -> p2p.NewBlockResponse
	5, // 11: p2p.P2P.NewTransaction:output_type -> p2p.NewTransactionResponse
	8, // 12: p2p.P2P.NewPeer:output_type -> p2p.PeerList
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_p2p_protocol_proto_init() }
func file_proto_p2p_protocol_proto_init() {
	if File_proto_p2p_protocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_p2p_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_p2p_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trans); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_p2p_protocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockChain); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_p2p_protocol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TailBlockRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_p2p_protocol_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBlockChainResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_p2p_protocol_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewTransactionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_p2p_protocol_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewBlockResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_p2p_protocol_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Peer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_p2p_protocol_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_p2p_protocol_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransPool); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_p2p_protocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_p2p_protocol_proto_goTypes,
		DependencyIndexes: file_proto_p2p_protocol_proto_depIdxs,
		MessageInfos:      file_proto_p2p_protocol_proto_msgTypes,
	}.Build()
	File_proto_p2p_protocol_proto = out.File
	file_proto_p2p_protocol_proto_rawDesc = nil
	file_proto_p2p_protocol_proto_goTypes = nil
	file_proto_p2p_protocol_proto_depIdxs = nil
}
