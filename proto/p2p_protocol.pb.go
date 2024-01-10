// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.6.1
// source: p2p_protocol.proto

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
		mi := &file_p2p_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_protocol_proto_msgTypes[0]
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
	return file_p2p_protocol_proto_rawDescGZIP(), []int{0}
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

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction string `protobuf:"bytes,1,opt,name=Transaction,proto3" json:"Transaction,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_p2p_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *Transaction) GetTransaction() string {
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
		mi := &file_p2p_protocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockChain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockChain) ProtoMessage() {}

func (x *BlockChain) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_protocol_proto_msgTypes[2]
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
	return file_p2p_protocol_proto_rawDescGZIP(), []int{2}
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
		mi := &file_p2p_protocol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TailBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TailBlockRequest) ProtoMessage() {}

func (x *TailBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_protocol_proto_msgTypes[3]
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
	return file_p2p_protocol_proto_rawDescGZIP(), []int{3}
}

type DeliverBlockChainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeliverBlockChainResponse) Reset() {
	*x = DeliverBlockChainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_protocol_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliverBlockChainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliverBlockChainResponse) ProtoMessage() {}

func (x *DeliverBlockChainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_protocol_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliverBlockChainResponse.ProtoReflect.Descriptor instead.
func (*DeliverBlockChainResponse) Descriptor() ([]byte, []int) {
	return file_p2p_protocol_proto_rawDescGZIP(), []int{4}
}

type NewTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NewTransactionResponse) Reset() {
	*x = NewTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_protocol_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTransactionResponse) ProtoMessage() {}

func (x *NewTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_protocol_proto_msgTypes[5]
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
	return file_p2p_protocol_proto_rawDescGZIP(), []int{5}
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
		mi := &file_p2p_protocol_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewBlockResponse) ProtoMessage() {}

func (x *NewBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_protocol_proto_msgTypes[6]
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
	return file_p2p_protocol_proto_rawDescGZIP(), []int{6}
}

func (x *NewBlockResponse) GetChainNeedUpdate() bool {
	if x != nil {
		return x.ChainNeedUpdate
	}
	return false
}

var File_p2p_protocol_proto protoreflect.FileDescriptor

var file_p2p_protocol_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x32, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x70, 0x32, 0x70, 0x22, 0x91, 0x01, 0x0a, 0x05, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x56, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x56, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a,
	0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48, 0x61, 0x73, 0x68, 0x22, 0x2f, 0x0a,
	0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2e,
	0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x05,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x32,
	0x70, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x22, 0x12,
	0x0a, 0x10, 0x54, 0x61, 0x69, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x1b, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x18, 0x0a, 0x16, 0x4e, 0x65, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3c, 0x0a, 0x10, 0x4e, 0x65, 0x77,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a,
	0x0f, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x65, 0x65, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x65, 0x65,
	0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x32, 0xfa, 0x01, 0x0a, 0x03, 0x50, 0x32, 0x50, 0x12,
	0x37, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x61, 0x69, 0x6c, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x15, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x54, 0x61, 0x69, 0x6c, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x32, 0x70,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x0f, 0x2e,
	0x70, 0x32, 0x70, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x1a, 0x1e,
	0x2e, 0x70, 0x32, 0x70, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x2f, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0a, 0x2e, 0x70,
	0x32, 0x70, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x15, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x4e,
	0x65, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x41, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x10, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1b, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x4e, 0x65, 0x77, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x70, 0x32, 0x70, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_p2p_protocol_proto_rawDescOnce sync.Once
	file_p2p_protocol_proto_rawDescData = file_p2p_protocol_proto_rawDesc
)

func file_p2p_protocol_proto_rawDescGZIP() []byte {
	file_p2p_protocol_proto_rawDescOnce.Do(func() {
		file_p2p_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_p2p_protocol_proto_rawDescData)
	})
	return file_p2p_protocol_proto_rawDescData
}

var file_p2p_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_p2p_protocol_proto_goTypes = []interface{}{
	(*Block)(nil),                     // 0: p2p.Block
	(*Transaction)(nil),               // 1: p2p.Transaction
	(*BlockChain)(nil),                // 2: p2p.BlockChain
	(*TailBlockRequest)(nil),          // 3: p2p.TailBlockRequest
	(*DeliverBlockChainResponse)(nil), // 4: p2p.DeliverBlockChainResponse
	(*NewTransactionResponse)(nil),    // 5: p2p.NewTransactionResponse
	(*NewBlockResponse)(nil),          // 6: p2p.NewBlockResponse
}
var file_p2p_protocol_proto_depIdxs = []int32{
	0, // 0: p2p.BlockChain.Chain:type_name -> p2p.Block
	3, // 1: p2p.P2P.RequestTailBlock:input_type -> p2p.TailBlockRequest
	2, // 2: p2p.P2P.DeliverBlockChain:input_type -> p2p.BlockChain
	0, // 3: p2p.P2P.NewBlock:input_type -> p2p.Block
	1, // 4: p2p.P2P.NewTransaction:input_type -> p2p.Transaction
	0, // 5: p2p.P2P.RequestTailBlock:output_type -> p2p.Block
	4, // 6: p2p.P2P.DeliverBlockChain:output_type -> p2p.DeliverBlockChainResponse
	6, // 7: p2p.P2P.NewBlock:output_type -> p2p.NewBlockResponse
	5, // 8: p2p.P2P.NewTransaction:output_type -> p2p.NewTransactionResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_p2p_protocol_proto_init() }
func file_p2p_protocol_proto_init() {
	if File_p2p_protocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_p2p_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_p2p_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_p2p_protocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_p2p_protocol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_p2p_protocol_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliverBlockChainResponse); i {
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
		file_p2p_protocol_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_p2p_protocol_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_p2p_protocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_p2p_protocol_proto_goTypes,
		DependencyIndexes: file_p2p_protocol_proto_depIdxs,
		MessageInfos:      file_p2p_protocol_proto_msgTypes,
	}.Build()
	File_p2p_protocol_proto = out.File
	file_p2p_protocol_proto_rawDesc = nil
	file_p2p_protocol_proto_goTypes = nil
	file_p2p_protocol_proto_depIdxs = nil
}