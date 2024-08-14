// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: transaction.proto

package transaction_service

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

type TransactionRPCResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response int32  `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TransactionRPCResponse) Reset() {
	*x = TransactionRPCResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionRPCResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionRPCResponse) ProtoMessage() {}

func (x *TransactionRPCResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionRPCResponse.ProtoReflect.Descriptor instead.
func (*TransactionRPCResponse) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionRPCResponse) GetResponse() int32 {
	if x != nil {
		return x.Response
	}
	return 0
}

func (x *TransactionRPCResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type TransactionQueryCriteria struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID              uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	VendorAddress   string `protobuf:"bytes,2,opt,name=VendorAddress,proto3" json:"VendorAddress,omitempty"`
	TransactionUUID string `protobuf:"bytes,3,opt,name=TransactionUUID,proto3" json:"TransactionUUID,omitempty"`
}

func (x *TransactionQueryCriteria) Reset() {
	*x = TransactionQueryCriteria{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionQueryCriteria) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionQueryCriteria) ProtoMessage() {}

func (x *TransactionQueryCriteria) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionQueryCriteria.ProtoReflect.Descriptor instead.
func (*TransactionQueryCriteria) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionQueryCriteria) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *TransactionQueryCriteria) GetVendorAddress() string {
	if x != nil {
		return x.VendorAddress
	}
	return ""
}

func (x *TransactionQueryCriteria) GetTransactionUUID() string {
	if x != nil {
		return x.TransactionUUID
	}
	return ""
}

type TransferMoneyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKey                string `protobuf:"bytes,1,opt,name=apiKey,proto3" json:"apiKey,omitempty"`
	SourceOwnerID         uint64 `protobuf:"varint,2,opt,name=SourceOwnerID,proto3" json:"SourceOwnerID,omitempty"`
	SourceOwnerTable      string `protobuf:"bytes,3,opt,name=SourceOwnerTable,proto3" json:"SourceOwnerTable,omitempty"`
	DestinationOwnerID    uint64 `protobuf:"varint,4,opt,name=DestinationOwnerID,proto3" json:"DestinationOwnerID,omitempty"`
	DestinationOwnerTable string `protobuf:"bytes,5,opt,name=DestinationOwnerTable,proto3" json:"DestinationOwnerTable,omitempty"`
	IDWalletUpdateAction  uint64 `protobuf:"varint,6,opt,name=IDWalletUpdateAction,proto3" json:"IDWalletUpdateAction,omitempty"`
	Ammount               int64  `protobuf:"varint,7,opt,name=Ammount,proto3" json:"Ammount,omitempty"`
	Description           string `protobuf:"bytes,8,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *TransferMoneyRequest) Reset() {
	*x = TransferMoneyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferMoneyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferMoneyRequest) ProtoMessage() {}

func (x *TransferMoneyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferMoneyRequest.ProtoReflect.Descriptor instead.
func (*TransferMoneyRequest) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *TransferMoneyRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *TransferMoneyRequest) GetSourceOwnerID() uint64 {
	if x != nil {
		return x.SourceOwnerID
	}
	return 0
}

func (x *TransferMoneyRequest) GetSourceOwnerTable() string {
	if x != nil {
		return x.SourceOwnerTable
	}
	return ""
}

func (x *TransferMoneyRequest) GetDestinationOwnerID() uint64 {
	if x != nil {
		return x.DestinationOwnerID
	}
	return 0
}

func (x *TransferMoneyRequest) GetDestinationOwnerTable() string {
	if x != nil {
		return x.DestinationOwnerTable
	}
	return ""
}

func (x *TransferMoneyRequest) GetIDWalletUpdateAction() uint64 {
	if x != nil {
		return x.IDWalletUpdateAction
	}
	return 0
}

func (x *TransferMoneyRequest) GetAmmount() int64 {
	if x != nil {
		return x.Ammount
	}
	return 0
}

func (x *TransferMoneyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type TopUpWalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKey               string `protobuf:"bytes,1,opt,name=apiKey,proto3" json:"apiKey,omitempty"`
	OwnerID              uint64 `protobuf:"varint,2,opt,name=OwnerID,proto3" json:"OwnerID,omitempty"`
	OwnerTable           string `protobuf:"bytes,3,opt,name=OwnerTable,proto3" json:"OwnerTable,omitempty"`
	IDWalletUpdateAction uint64 `protobuf:"varint,4,opt,name=IDWalletUpdateAction,proto3" json:"IDWalletUpdateAction,omitempty"`
	Ammount              int64  `protobuf:"varint,5,opt,name=Ammount,proto3" json:"Ammount,omitempty"`
	Description          string `protobuf:"bytes,6,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *TopUpWalletRequest) Reset() {
	*x = TopUpWalletRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopUpWalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopUpWalletRequest) ProtoMessage() {}

func (x *TopUpWalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopUpWalletRequest.ProtoReflect.Descriptor instead.
func (*TopUpWalletRequest) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{3}
}

func (x *TopUpWalletRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *TopUpWalletRequest) GetOwnerID() uint64 {
	if x != nil {
		return x.OwnerID
	}
	return 0
}

func (x *TopUpWalletRequest) GetOwnerTable() string {
	if x != nil {
		return x.OwnerTable
	}
	return ""
}

func (x *TopUpWalletRequest) GetIDWalletUpdateAction() uint64 {
	if x != nil {
		return x.IDWalletUpdateAction
	}
	return 0
}

func (x *TopUpWalletRequest) GetAmmount() int64 {
	if x != nil {
		return x.Ammount
	}
	return 0
}

func (x *TopUpWalletRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// Transactions Structure
type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionUUID      string                `protobuf:"bytes,1,opt,name=TransactionUUID,proto3" json:"TransactionUUID,omitempty"`
	TransferDetails      *TransferMoneyRequest `protobuf:"bytes,2,opt,name=TransferDetails,proto3,oneof" json:"TransferDetails,omitempty"`
	TopUpDetails         *TopUpWalletRequest   `protobuf:"bytes,3,opt,name=TopUpDetails,proto3,oneof" json:"TopUpDetails,omitempty"`
	IDWalletUpdateAction uint64                `protobuf:"varint,4,opt,name=IDWalletUpdateAction,proto3" json:"IDWalletUpdateAction,omitempty"`
	WalletUpdateAction   string                `protobuf:"bytes,5,opt,name=WalletUpdateAction,proto3" json:"WalletUpdateAction,omitempty"`
	IDWalletLogStatus    uint64                `protobuf:"varint,6,opt,name=IDWalletLogStatus,proto3" json:"IDWalletLogStatus,omitempty"`
	WalletLogStatus      string                `protobuf:"bytes,7,opt,name=WalletLogStatus,proto3" json:"WalletLogStatus,omitempty"`
	Ammount              int64                 `protobuf:"zigzag64,8,opt,name=Ammount,proto3" json:"Ammount,omitempty"`
	Description          string                `protobuf:"bytes,9,opt,name=Description,proto3" json:"Description,omitempty"`
	CreatedAt            string                `protobuf:"bytes,10,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[4]
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
	return file_transaction_proto_rawDescGZIP(), []int{4}
}

func (x *Transaction) GetTransactionUUID() string {
	if x != nil {
		return x.TransactionUUID
	}
	return ""
}

func (x *Transaction) GetTransferDetails() *TransferMoneyRequest {
	if x != nil {
		return x.TransferDetails
	}
	return nil
}

func (x *Transaction) GetTopUpDetails() *TopUpWalletRequest {
	if x != nil {
		return x.TopUpDetails
	}
	return nil
}

func (x *Transaction) GetIDWalletUpdateAction() uint64 {
	if x != nil {
		return x.IDWalletUpdateAction
	}
	return 0
}

func (x *Transaction) GetWalletUpdateAction() string {
	if x != nil {
		return x.WalletUpdateAction
	}
	return ""
}

func (x *Transaction) GetIDWalletLogStatus() uint64 {
	if x != nil {
		return x.IDWalletLogStatus
	}
	return 0
}

func (x *Transaction) GetWalletLogStatus() string {
	if x != nil {
		return x.WalletLogStatus
	}
	return ""
}

func (x *Transaction) GetAmmount() int64 {
	if x != nil {
		return x.Ammount
	}
	return 0
}

func (x *Transaction) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Transaction) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type GetTransactionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKey        string                    `protobuf:"bytes,1,opt,name=apiKey,proto3" json:"apiKey,omitempty"`
	QueryCriteria *TransactionQueryCriteria `protobuf:"bytes,2,opt,name=queryCriteria,proto3,oneof" json:"queryCriteria,omitempty"`
}

func (x *GetTransactionsRequest) Reset() {
	*x = GetTransactionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionsRequest) ProtoMessage() {}

func (x *GetTransactionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionsRequest.ProtoReflect.Descriptor instead.
func (*GetTransactionsRequest) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{5}
}

func (x *GetTransactionsRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *GetTransactionsRequest) GetQueryCriteria() *TransactionQueryCriteria {
	if x != nil {
		return x.QueryCriteria
	}
	return nil
}

type GetTransactionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response     *TransactionRPCResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Transactions []*Transaction          `protobuf:"bytes,3,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *GetTransactionsResponse) Reset() {
	*x = GetTransactionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionsResponse) ProtoMessage() {}

func (x *GetTransactionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionsResponse.ProtoReflect.Descriptor instead.
func (*GetTransactionsResponse) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{6}
}

func (x *GetTransactionsResponse) GetResponse() *TransactionRPCResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *GetTransactionsResponse) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

var File_transaction_proto protoreflect.FileDescriptor

var file_transaction_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x4e, 0x0a, 0x16, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x7a, 0x0a, 0x18, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d,
	0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x55, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x55, 0x49, 0x44, 0x22, 0xd6, 0x02, 0x0a,
	0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x24, 0x0a,
	0x0d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x10, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x2e, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x34, 0x0a, 0x15, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15,
	0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x49, 0x44, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x14, 0x49, 0x44, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x6d, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x41, 0x6d, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd6, 0x01, 0x0a, 0x12, 0x54, 0x6f, 0x70, 0x55, 0x70, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70,
	0x69, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e,
	0x0a, 0x0a, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x32,
	0x0a, 0x14, 0x49, 0x44, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x49, 0x44,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x6d, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x41, 0x6d, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8e,
	0x04, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28,
	0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x55, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x55, 0x55, 0x49, 0x44, 0x12, 0x50, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x88, 0x01, 0x01, 0x12, 0x48, 0x0a, 0x0c, 0x54, 0x6f,
	0x70, 0x55, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54,
	0x6f, 0x70, 0x55, 0x70, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x01, 0x52, 0x0c, 0x54, 0x6f, 0x70, 0x55, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x14, 0x49, 0x44, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x14, 0x49, 0x44, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x49, 0x44, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x11, 0x49, 0x44, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x4c, 0x6f, 0x67,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x4c, 0x6f, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x41, 0x6d, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x12, 0x52, 0x07, 0x41, 0x6d, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x0f,
	0x0a, 0x0d, 0x5f, 0x54, 0x6f, 0x70, 0x55, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22,
	0x94, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70,
	0x69, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b,
	0x65, 0x79, 0x12, 0x50, 0x0a, 0x0d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x43, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x48, 0x00, 0x52, 0x0d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x43, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x22, 0x98, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x50,
	0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x32, 0xa0, 0x02, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x21, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x53, 0x0a, 0x0b, 0x54, 0x6f, 0x70, 0x55, 0x70, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x12, 0x1f, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54,
	0x6f, 0x70, 0x55, 0x70, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x50, 0x43, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transaction_proto_rawDescOnce sync.Once
	file_transaction_proto_rawDescData = file_transaction_proto_rawDesc
)

func file_transaction_proto_rawDescGZIP() []byte {
	file_transaction_proto_rawDescOnce.Do(func() {
		file_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_proto_rawDescData)
	})
	return file_transaction_proto_rawDescData
}

var file_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_transaction_proto_goTypes = []interface{}{
	(*TransactionRPCResponse)(nil),   // 0: transaction.TransactionRPCResponse
	(*TransactionQueryCriteria)(nil), // 1: transaction.TransactionQueryCriteria
	(*TransferMoneyRequest)(nil),     // 2: transaction.TransferMoneyRequest
	(*TopUpWalletRequest)(nil),       // 3: transaction.TopUpWalletRequest
	(*Transaction)(nil),              // 4: transaction.Transaction
	(*GetTransactionsRequest)(nil),   // 5: transaction.GetTransactionsRequest
	(*GetTransactionsResponse)(nil),  // 6: transaction.GetTransactionsResponse
}
var file_transaction_proto_depIdxs = []int32{
	2, // 0: transaction.Transaction.TransferDetails:type_name -> transaction.TransferMoneyRequest
	3, // 1: transaction.Transaction.TopUpDetails:type_name -> transaction.TopUpWalletRequest
	1, // 2: transaction.GetTransactionsRequest.queryCriteria:type_name -> transaction.TransactionQueryCriteria
	0, // 3: transaction.GetTransactionsResponse.response:type_name -> transaction.TransactionRPCResponse
	4, // 4: transaction.GetTransactionsResponse.transactions:type_name -> transaction.Transaction
	2, // 5: transaction.TransactionService.TransferMoney:input_type -> transaction.TransferMoneyRequest
	3, // 6: transaction.TransactionService.TopUpWallet:input_type -> transaction.TopUpWalletRequest
	5, // 7: transaction.TransactionService.GetTransactions:input_type -> transaction.GetTransactionsRequest
	0, // 8: transaction.TransactionService.TransferMoney:output_type -> transaction.TransactionRPCResponse
	0, // 9: transaction.TransactionService.TopUpWallet:output_type -> transaction.TransactionRPCResponse
	6, // 10: transaction.TransactionService.GetTransactions:output_type -> transaction.GetTransactionsResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_transaction_proto_init() }
func file_transaction_proto_init() {
	if File_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionRPCResponse); i {
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
		file_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionQueryCriteria); i {
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
		file_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferMoneyRequest); i {
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
		file_transaction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopUpWalletRequest); i {
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
		file_transaction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_transaction_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionsRequest); i {
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
		file_transaction_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionsResponse); i {
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
	file_transaction_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_transaction_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transaction_proto_goTypes,
		DependencyIndexes: file_transaction_proto_depIdxs,
		MessageInfos:      file_transaction_proto_msgTypes,
	}.Build()
	File_transaction_proto = out.File
	file_transaction_proto_rawDesc = nil
	file_transaction_proto_goTypes = nil
	file_transaction_proto_depIdxs = nil
}
