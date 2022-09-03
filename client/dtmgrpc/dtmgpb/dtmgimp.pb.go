// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: client/dtmgrpc/dtmgpb/dtmgimp.proto

package dtmgpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DtmTransOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WaitResult    bool  `protobuf:"varint,1,opt,name=WaitResult,proto3" json:"WaitResult,omitempty"`
	TimeoutToFail int64 `protobuf:"varint,2,opt,name=TimeoutToFail,proto3" json:"TimeoutToFail,omitempty"`
	RetryInterval int64 `protobuf:"varint,3,opt,name=RetryInterval,proto3" json:"RetryInterval,omitempty"`
	// repeated string PassthroughHeaders = 4; // depreceated
	BranchHeaders  map[string]string `protobuf:"bytes,5,rep,name=BranchHeaders,proto3" json:"BranchHeaders,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RequestTimeout int64             `protobuf:"varint,6,opt,name=RequestTimeout,proto3" json:"RequestTimeout,omitempty"`
	RetryLimit     int64             `protobuf:"varint,7,opt,name=RetryLimit,proto3" json:"RetryLimit,omitempty"`
}

func (x *DtmTransOptions) Reset() {
	*x = DtmTransOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DtmTransOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DtmTransOptions) ProtoMessage() {}

func (x *DtmTransOptions) ProtoReflect() protoreflect.Message {
	mi := &file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DtmTransOptions.ProtoReflect.Descriptor instead.
func (*DtmTransOptions) Descriptor() ([]byte, []int) {
	return file_client_dtmgrpc_dtmgpb_dtmgimp_proto_rawDescGZIP(), []int{0}
}

func (x *DtmTransOptions) GetWaitResult() bool {
	if x != nil {
		return x.WaitResult
	}
	return false
}

func (x *DtmTransOptions) GetTimeoutToFail() int64 {
	if x != nil {
		return x.TimeoutToFail
	}
	return 0
}

func (x *DtmTransOptions) GetRetryInterval() int64 {
	if x != nil {
		return x.RetryInterval
	}
	return 0
}

func (x *DtmTransOptions) GetBranchHeaders() map[string]string {
	if x != nil {
		return x.BranchHeaders
	}
	return nil
}

func (x *DtmTransOptions) GetRequestTimeout() int64 {
	if x != nil {
		return x.RequestTimeout
	}
	return 0
}

func (x *DtmTransOptions) GetRetryLimit() int64 {
	if x != nil {
		return x.RetryLimit
	}
	return 0
}

// DtmRequest request sent to dtm server
type DtmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gid            string            `protobuf:"bytes,1,opt,name=Gid,proto3" json:"Gid,omitempty"`
	TransType      string            `protobuf:"bytes,2,opt,name=TransType,proto3" json:"TransType,omitempty"`
	TransOptions   *DtmTransOptions  `protobuf:"bytes,3,opt,name=TransOptions,proto3" json:"TransOptions,omitempty"`
	CustomedData   string            `protobuf:"bytes,4,opt,name=CustomedData,proto3" json:"CustomedData,omitempty"`
	BinPayloads    [][]byte          `protobuf:"bytes,5,rep,name=BinPayloads,proto3" json:"BinPayloads,omitempty"`     // for Msg/Saga/Workflow branch payloads
	QueryPrepared  string            `protobuf:"bytes,6,opt,name=QueryPrepared,proto3" json:"QueryPrepared,omitempty"` // for Msg
	Steps          string            `protobuf:"bytes,7,opt,name=Steps,proto3" json:"Steps,omitempty"`
	ReqExtra       map[string]string `protobuf:"bytes,8,rep,name=ReqExtra,proto3" json:"ReqExtra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RollbackReason string            `protobuf:"bytes,9,opt,name=RollbackReason,proto3" json:"RollbackReason,omitempty"`
}

func (x *DtmRequest) Reset() {
	*x = DtmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DtmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DtmRequest) ProtoMessage() {}

func (x *DtmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DtmRequest.ProtoReflect.Descriptor instead.
func (*DtmRequest) Descriptor() ([]byte, []int) {
	return file_client_dtmgrpc_dtmgpb_dtmgimp_proto_rawDescGZIP(), []int{1}
}

func (x *DtmRequest) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

func (x *DtmRequest) GetTransType() string {
	if x != nil {
		return x.TransType
	}
	return ""
}

func (x *DtmRequest) GetTransOptions() *DtmTransOptions {
	if x != nil {
		return x.TransOptions
	}
	return nil
}

func (x *DtmRequest) GetCustomedData() string {
	if x != nil {
		return x.CustomedData
	}
	return ""
}

func (x *DtmRequest) GetBinPayloads() [][]byte {
	if x != nil {
		return x.BinPayloads
	}
	return nil
}

func (x *DtmRequest) GetQueryPrepared() string {
	if x != nil {
		return x.QueryPrepared
	}
	return ""
}

func (x *DtmRequest) GetSteps() string {
	if x != nil {
		return x.Steps
	}
	return ""
}

func (x *DtmRequest) GetReqExtra() map[string]string {
	if x != nil {
		return x.ReqExtra
	}
	return nil
}

func (x *DtmRequest) GetRollbackReason() string {
	if x != nil {
		return x.RollbackReason
	}
	return ""
}

type DtmGidReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gid string `protobuf:"bytes,1,opt,name=Gid,proto3" json:"Gid,omitempty"`
}

func (x *DtmGidReply) Reset() {
	*x = DtmGidReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DtmGidReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DtmGidReply) ProtoMessage() {}

func (x *DtmGidReply) ProtoReflect() protoreflect.Message {
	mi := &file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DtmGidReply.ProtoReflect.Descriptor instead.
func (*DtmGidReply) Descriptor() ([]byte, []int) {
	return file_client_dtmgrpc_dtmgpb_dtmgimp_proto_rawDescGZIP(), []int{2}
}

func (x *DtmGidReply) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

type DtmBranchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gid         string            `protobuf:"bytes,1,opt,name=Gid,proto3" json:"Gid,omitempty"`
	TransType   string            `protobuf:"bytes,2,opt,name=TransType,proto3" json:"TransType,omitempty"`
	BranchID    string            `protobuf:"bytes,3,opt,name=BranchID,proto3" json:"BranchID,omitempty"`
	Op          string            `protobuf:"bytes,4,opt,name=Op,proto3" json:"Op,omitempty"`
	Data        map[string]string `protobuf:"bytes,5,rep,name=Data,proto3" json:"Data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	BusiPayload []byte            `protobuf:"bytes,6,opt,name=BusiPayload,proto3" json:"BusiPayload,omitempty"`
}

func (x *DtmBranchRequest) Reset() {
	*x = DtmBranchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DtmBranchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DtmBranchRequest) ProtoMessage() {}

func (x *DtmBranchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DtmBranchRequest.ProtoReflect.Descriptor instead.
func (*DtmBranchRequest) Descriptor() ([]byte, []int) {
	return file_client_dtmgrpc_dtmgpb_dtmgimp_proto_rawDescGZIP(), []int{3}
}

func (x *DtmBranchRequest) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

func (x *DtmBranchRequest) GetTransType() string {
	if x != nil {
		return x.TransType
	}
	return ""
}

func (x *DtmBranchRequest) GetBranchID() string {
	if x != nil {
		return x.BranchID
	}
	return ""
}

func (x *DtmBranchRequest) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *DtmBranchRequest) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DtmBranchRequest) GetBusiPayload() []byte {
	if x != nil {
		return x.BusiPayload
	}
	return nil
}

type DtmProgressesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *DtmTransaction `protobuf:"bytes,1,opt,name=Transaction,proto3" json:"Transaction,omitempty"`
	Progresses  []*DtmProgress  `protobuf:"bytes,2,rep,name=Progresses,proto3" json:"Progresses,omitempty"`
}

func (x *DtmProgressesReply) Reset() {
	*x = DtmProgressesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DtmProgressesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DtmProgressesReply) ProtoMessage() {}

func (x *DtmProgressesReply) ProtoReflect() protoreflect.Message {
	mi := &file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DtmProgressesReply.ProtoReflect.Descriptor instead.
func (*DtmProgressesReply) Descriptor() ([]byte, []int) {
	return file_client_dtmgrpc_dtmgpb_dtmgimp_proto_rawDescGZIP(), []int{4}
}

func (x *DtmProgressesReply) GetTransaction() *DtmTransaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *DtmProgressesReply) GetProgresses() []*DtmProgress {
	if x != nil {
		return x.Progresses
	}
	return nil
}

type DtmTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gid            string `protobuf:"bytes,1,opt,name=Gid,proto3" json:"Gid,omitempty"`
	Status         string `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	RollbackReason string `protobuf:"bytes,3,opt,name=RollbackReason,proto3" json:"RollbackReason,omitempty"`
	Result         string `protobuf:"bytes,4,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *DtmTransaction) Reset() {
	*x = DtmTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DtmTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DtmTransaction) ProtoMessage() {}

func (x *DtmTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DtmTransaction.ProtoReflect.Descriptor instead.
func (*DtmTransaction) Descriptor() ([]byte, []int) {
	return file_client_dtmgrpc_dtmgpb_dtmgimp_proto_rawDescGZIP(), []int{5}
}

func (x *DtmTransaction) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

func (x *DtmTransaction) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DtmTransaction) GetRollbackReason() string {
	if x != nil {
		return x.RollbackReason
	}
	return ""
}

func (x *DtmTransaction) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type DtmProgress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   string `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	BinData  []byte `protobuf:"bytes,2,opt,name=BinData,proto3" json:"BinData,omitempty"`
	BranchID string `protobuf:"bytes,3,opt,name=BranchID,proto3" json:"BranchID,omitempty"`
	Op       string `protobuf:"bytes,4,opt,name=Op,proto3" json:"Op,omitempty"`
}

func (x *DtmProgress) Reset() {
	*x = DtmProgress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DtmProgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DtmProgress) ProtoMessage() {}

func (x *DtmProgress) ProtoReflect() protoreflect.Message {
	mi := &file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DtmProgress.ProtoReflect.Descriptor instead.
func (*DtmProgress) Descriptor() ([]byte, []int) {
	return file_client_dtmgrpc_dtmgpb_dtmgimp_proto_rawDescGZIP(), []int{6}
}

func (x *DtmProgress) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DtmProgress) GetBinData() []byte {
	if x != nil {
		return x.BinData
	}
	return nil
}

func (x *DtmProgress) GetBranchID() string {
	if x != nil {
		return x.BranchID
	}
	return ""
}

func (x *DtmProgress) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

type DtmTopicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic  string `protobuf:"bytes,1,opt,name=Topic,proto3" json:"Topic,omitempty"`
	URL    string `protobuf:"bytes,2,opt,name=URL,proto3" json:"URL,omitempty"`
	Remark string `protobuf:"bytes,3,opt,name=Remark,proto3" json:"Remark,omitempty"`
}

func (x *DtmTopicRequest) Reset() {
	*x = DtmTopicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DtmTopicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DtmTopicRequest) ProtoMessage() {}

func (x *DtmTopicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DtmTopicRequest.ProtoReflect.Descriptor instead.
func (*DtmTopicRequest) Descriptor() ([]byte, []int) {
	return file_client_dtmgrpc_dtmgpb_dtmgimp_proto_rawDescGZIP(), []int{7}
}

func (x *DtmTopicRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *DtmTopicRequest) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *DtmTopicRequest) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

var File_client_dtmgrpc_dtmgpb_dtmgimp_proto protoreflect.FileDescriptor

var file_client_dtmgrpc_dtmgpb_dtmgimp_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x64, 0x74, 0x6d, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x64, 0x74, 0x6d, 0x67, 0x70, 0x62, 0x2f, 0x64, 0x74, 0x6d, 0x67, 0x69, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x74, 0x6d, 0x67, 0x69, 0x6d, 0x70, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x02, 0x0a, 0x0f,
	0x44, 0x74, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x57, 0x61, 0x69, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x57, 0x61, 0x69, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x24, 0x0a, 0x0d, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x54, 0x6f, 0x46, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x54,
	0x6f, 0x46, 0x61, 0x69, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x52, 0x65, 0x74, 0x72, 0x79, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x52, 0x65,
	0x74, 0x72, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x51, 0x0a, 0x0d, 0x42,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x64, 0x74, 0x6d, 0x67, 0x69, 0x6d, 0x70, 0x2e, 0x44, 0x74, 0x6d,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0d, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x26,
	0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x74, 0x72, 0x79, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x52, 0x65, 0x74, 0x72,
	0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x1a, 0x40, 0x0a, 0x12, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa0, 0x03, 0x0a, 0x0a, 0x44, 0x74, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x47, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x47, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x64, 0x74, 0x6d, 0x67, 0x69, 0x6d, 0x70, 0x2e, 0x44, 0x74, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x69, 0x6e,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0b,
	0x42, 0x69, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x65, 0x70, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x53, 0x74, 0x65, 0x70, 0x73, 0x12, 0x3d, 0x0a, 0x08, 0x52, 0x65, 0x71, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x74, 0x6d, 0x67,
	0x69, 0x6d, 0x70, 0x2e, 0x44, 0x74, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52,
	0x65, 0x71, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x52, 0x65,
	0x71, 0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x1a, 0x3b,
	0x0a, 0x0d, 0x52, 0x65, 0x71, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1f, 0x0a, 0x0b, 0x44,
	0x74, 0x6d, 0x47, 0x69, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x47, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x47, 0x69, 0x64, 0x22, 0x82, 0x02, 0x0a,
	0x10, 0x44, 0x74, 0x6d, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x47, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x47, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x44, 0x12, 0x0e, 0x0a,
	0x02, 0x4f, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x4f, 0x70, 0x12, 0x37, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x64, 0x74,
	0x6d, 0x67, 0x69, 0x6d, 0x70, 0x2e, 0x44, 0x74, 0x6d, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x75, 0x73, 0x69, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x42, 0x75, 0x73,
	0x69, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x85, 0x01, 0x0a, 0x12, 0x44, 0x74, 0x6d, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x39, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x64, 0x74, 0x6d, 0x67, 0x69, 0x6d, 0x70, 0x2e, 0x44, 0x74, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x74, 0x6d, 0x67, 0x69, 0x6d,
	0x70, 0x2e, 0x44, 0x74, 0x6d, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0a, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x7a, 0x0a, 0x0e, 0x44, 0x74, 0x6d,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x47,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x47, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x52,
	0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x6b, 0x0a, 0x0b, 0x44, 0x74, 0x6d, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x42, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x42,
	0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x4f, 0x70, 0x22, 0x51, 0x0a, 0x0f, 0x44, 0x74, 0x6d, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x55,
	0x52, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x32, 0xbf, 0x04, 0x0a, 0x03, 0x44, 0x74, 0x6d, 0x12, 0x38, 0x0a,
	0x06, 0x4e, 0x65, 0x77, 0x47, 0x69, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x14, 0x2e, 0x64, 0x74, 0x6d, 0x67, 0x69, 0x6d, 0x70, 0x2e, 0x44, 0x74, 0x6d, 0x47, 0x69, 0x64,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x12, 0x13, 0x2e, 0x64, 0x74, 0x6d, 0x67, 0x69, 0x6d, 0x70, 0x2e, 0x44, 0x74, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x38, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x12, 0x13, 0x2e, 0x64, 0x74,
	0x6d, 0x67, 0x69, 0x6d, 0x70, 0x2e, 0x44, 0x74, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x05, 0x41, 0x62,
	0x6f, 0x72, 0x74, 0x12, 0x13, 0x2e, 0x64, 0x74, 0x6d, 0x67, 0x69, 0x6d, 0x70, 0x2e, 0x44, 0x74,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x45, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x12, 0x19, 0x2e, 0x64, 0x74, 0x6d, 0x67, 0x69, 0x6d, 0x70, 0x2e, 0x44,
	0x74, 0x6d, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0f, 0x50, 0x72, 0x65,
	0x70, 0x61, 0x72, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x13, 0x2e, 0x64,
	0x74, 0x6d, 0x67, 0x69, 0x6d, 0x70, 0x2e, 0x44, 0x74, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x74, 0x6d, 0x67, 0x69, 0x6d, 0x70, 0x2e, 0x44, 0x74, 0x6d, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x3f, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x18, 0x2e,
	0x64, 0x74, 0x6d, 0x67, 0x69, 0x6d, 0x70, 0x2e, 0x44, 0x74, 0x6d, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x41, 0x0a, 0x0b, 0x55, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x12, 0x18, 0x2e, 0x64, 0x74, 0x6d, 0x67, 0x69, 0x6d, 0x70, 0x2e, 0x44, 0x74, 0x6d, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x12, 0x18, 0x2e, 0x64, 0x74, 0x6d, 0x67, 0x69, 0x6d, 0x70, 0x2e, 0x44, 0x74,
	0x6d, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x64, 0x74, 0x6d,
	0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_dtmgrpc_dtmgpb_dtmgimp_proto_rawDescOnce sync.Once
	file_client_dtmgrpc_dtmgpb_dtmgimp_proto_rawDescData = file_client_dtmgrpc_dtmgpb_dtmgimp_proto_rawDesc
)

func file_client_dtmgrpc_dtmgpb_dtmgimp_proto_rawDescGZIP() []byte {
	file_client_dtmgrpc_dtmgpb_dtmgimp_proto_rawDescOnce.Do(func() {
		file_client_dtmgrpc_dtmgpb_dtmgimp_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_dtmgrpc_dtmgpb_dtmgimp_proto_rawDescData)
	})
	return file_client_dtmgrpc_dtmgpb_dtmgimp_proto_rawDescData
}

var file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_client_dtmgrpc_dtmgpb_dtmgimp_proto_goTypes = []interface{}{
	(*DtmTransOptions)(nil),    // 0: dtmgimp.DtmTransOptions
	(*DtmRequest)(nil),         // 1: dtmgimp.DtmRequest
	(*DtmGidReply)(nil),        // 2: dtmgimp.DtmGidReply
	(*DtmBranchRequest)(nil),   // 3: dtmgimp.DtmBranchRequest
	(*DtmProgressesReply)(nil), // 4: dtmgimp.DtmProgressesReply
	(*DtmTransaction)(nil),     // 5: dtmgimp.DtmTransaction
	(*DtmProgress)(nil),        // 6: dtmgimp.DtmProgress
	(*DtmTopicRequest)(nil),    // 7: dtmgimp.DtmTopicRequest
	nil,                        // 8: dtmgimp.DtmTransOptions.BranchHeadersEntry
	nil,                        // 9: dtmgimp.DtmRequest.ReqExtraEntry
	nil,                        // 10: dtmgimp.DtmBranchRequest.DataEntry
	(*emptypb.Empty)(nil),      // 11: google.protobuf.Empty
}
var file_client_dtmgrpc_dtmgpb_dtmgimp_proto_depIdxs = []int32{
	8,  // 0: dtmgimp.DtmTransOptions.BranchHeaders:type_name -> dtmgimp.DtmTransOptions.BranchHeadersEntry
	0,  // 1: dtmgimp.DtmRequest.TransOptions:type_name -> dtmgimp.DtmTransOptions
	9,  // 2: dtmgimp.DtmRequest.ReqExtra:type_name -> dtmgimp.DtmRequest.ReqExtraEntry
	10, // 3: dtmgimp.DtmBranchRequest.Data:type_name -> dtmgimp.DtmBranchRequest.DataEntry
	5,  // 4: dtmgimp.DtmProgressesReply.Transaction:type_name -> dtmgimp.DtmTransaction
	6,  // 5: dtmgimp.DtmProgressesReply.Progresses:type_name -> dtmgimp.DtmProgress
	11, // 6: dtmgimp.Dtm.NewGid:input_type -> google.protobuf.Empty
	1,  // 7: dtmgimp.Dtm.Submit:input_type -> dtmgimp.DtmRequest
	1,  // 8: dtmgimp.Dtm.Prepare:input_type -> dtmgimp.DtmRequest
	1,  // 9: dtmgimp.Dtm.Abort:input_type -> dtmgimp.DtmRequest
	3,  // 10: dtmgimp.Dtm.RegisterBranch:input_type -> dtmgimp.DtmBranchRequest
	1,  // 11: dtmgimp.Dtm.PrepareWorkflow:input_type -> dtmgimp.DtmRequest
	7,  // 12: dtmgimp.Dtm.Subscribe:input_type -> dtmgimp.DtmTopicRequest
	7,  // 13: dtmgimp.Dtm.UnSubscribe:input_type -> dtmgimp.DtmTopicRequest
	7,  // 14: dtmgimp.Dtm.DeleteTopic:input_type -> dtmgimp.DtmTopicRequest
	2,  // 15: dtmgimp.Dtm.NewGid:output_type -> dtmgimp.DtmGidReply
	11, // 16: dtmgimp.Dtm.Submit:output_type -> google.protobuf.Empty
	11, // 17: dtmgimp.Dtm.Prepare:output_type -> google.protobuf.Empty
	11, // 18: dtmgimp.Dtm.Abort:output_type -> google.protobuf.Empty
	11, // 19: dtmgimp.Dtm.RegisterBranch:output_type -> google.protobuf.Empty
	4,  // 20: dtmgimp.Dtm.PrepareWorkflow:output_type -> dtmgimp.DtmProgressesReply
	11, // 21: dtmgimp.Dtm.Subscribe:output_type -> google.protobuf.Empty
	11, // 22: dtmgimp.Dtm.UnSubscribe:output_type -> google.protobuf.Empty
	11, // 23: dtmgimp.Dtm.DeleteTopic:output_type -> google.protobuf.Empty
	15, // [15:24] is the sub-list for method output_type
	6,  // [6:15] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_client_dtmgrpc_dtmgpb_dtmgimp_proto_init() }
func file_client_dtmgrpc_dtmgpb_dtmgimp_proto_init() {
	if File_client_dtmgrpc_dtmgpb_dtmgimp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DtmTransOptions); i {
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
		file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DtmRequest); i {
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
		file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DtmGidReply); i {
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
		file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DtmBranchRequest); i {
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
		file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DtmProgressesReply); i {
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
		file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DtmTransaction); i {
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
		file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DtmProgress); i {
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
		file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DtmTopicRequest); i {
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
			RawDescriptor: file_client_dtmgrpc_dtmgpb_dtmgimp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_dtmgrpc_dtmgpb_dtmgimp_proto_goTypes,
		DependencyIndexes: file_client_dtmgrpc_dtmgpb_dtmgimp_proto_depIdxs,
		MessageInfos:      file_client_dtmgrpc_dtmgpb_dtmgimp_proto_msgTypes,
	}.Build()
	File_client_dtmgrpc_dtmgpb_dtmgimp_proto = out.File
	file_client_dtmgrpc_dtmgpb_dtmgimp_proto_rawDesc = nil
	file_client_dtmgrpc_dtmgpb_dtmgimp_proto_goTypes = nil
	file_client_dtmgrpc_dtmgpb_dtmgimp_proto_depIdxs = nil
}
