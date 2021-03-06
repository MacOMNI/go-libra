// +build !js
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: admission_control.proto

package pbac

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	pbtypes "github.com/the729/go-libra/generated/pbtypes"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Additional statuses that are possible from admission control in addition
// to VM statuses.
type AdmissionControlStatusCode int32

const (
	// Validator accepted the transaction.
	AdmissionControlStatusCode_Accepted AdmissionControlStatusCode = 0
	// The sender is blacklisted.
	AdmissionControlStatusCode_Blacklisted AdmissionControlStatusCode = 1
	// The transaction is rejected, e.g. due to incorrect signature.
	AdmissionControlStatusCode_Rejected AdmissionControlStatusCode = 2
)

var AdmissionControlStatusCode_name = map[int32]string{
	0: "Accepted",
	1: "Blacklisted",
	2: "Rejected",
}

var AdmissionControlStatusCode_value = map[string]int32{
	"Accepted":    0,
	"Blacklisted": 1,
	"Rejected":    2,
}

func (x AdmissionControlStatusCode) String() string {
	return proto.EnumName(AdmissionControlStatusCode_name, int32(x))
}

func (AdmissionControlStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d13d6ff9aac892b4, []int{0}
}

// The request for submitting a transaction to an upstream validator or full
// node.
type AdmissionControlMsg struct {
	// Types that are valid to be assigned to Message:
	//	*AdmissionControlMsg_SubmitTransactionRequest
	//	*AdmissionControlMsg_SubmitTransactionResponse
	Message              isAdmissionControlMsg_Message `protobuf_oneof:"message"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *AdmissionControlMsg) Reset()         { *m = AdmissionControlMsg{} }
func (m *AdmissionControlMsg) String() string { return proto.CompactTextString(m) }
func (*AdmissionControlMsg) ProtoMessage()    {}
func (*AdmissionControlMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_d13d6ff9aac892b4, []int{0}
}

func (m *AdmissionControlMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdmissionControlMsg.Unmarshal(m, b)
}
func (m *AdmissionControlMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdmissionControlMsg.Marshal(b, m, deterministic)
}
func (m *AdmissionControlMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdmissionControlMsg.Merge(m, src)
}
func (m *AdmissionControlMsg) XXX_Size() int {
	return xxx_messageInfo_AdmissionControlMsg.Size(m)
}
func (m *AdmissionControlMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_AdmissionControlMsg.DiscardUnknown(m)
}

var xxx_messageInfo_AdmissionControlMsg proto.InternalMessageInfo

type isAdmissionControlMsg_Message interface {
	isAdmissionControlMsg_Message()
}

type AdmissionControlMsg_SubmitTransactionRequest struct {
	SubmitTransactionRequest *SubmitTransactionRequest `protobuf:"bytes,1,opt,name=submit_transaction_request,json=submitTransactionRequest,proto3,oneof"`
}

type AdmissionControlMsg_SubmitTransactionResponse struct {
	SubmitTransactionResponse *SubmitTransactionResponse `protobuf:"bytes,2,opt,name=submit_transaction_response,json=submitTransactionResponse,proto3,oneof"`
}

func (*AdmissionControlMsg_SubmitTransactionRequest) isAdmissionControlMsg_Message() {}

func (*AdmissionControlMsg_SubmitTransactionResponse) isAdmissionControlMsg_Message() {}

func (m *AdmissionControlMsg) GetMessage() isAdmissionControlMsg_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *AdmissionControlMsg) GetSubmitTransactionRequest() *SubmitTransactionRequest {
	if x, ok := m.GetMessage().(*AdmissionControlMsg_SubmitTransactionRequest); ok {
		return x.SubmitTransactionRequest
	}
	return nil
}

func (m *AdmissionControlMsg) GetSubmitTransactionResponse() *SubmitTransactionResponse {
	if x, ok := m.GetMessage().(*AdmissionControlMsg_SubmitTransactionResponse); ok {
		return x.SubmitTransactionResponse
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdmissionControlMsg) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdmissionControlMsg_SubmitTransactionRequest)(nil),
		(*AdmissionControlMsg_SubmitTransactionResponse)(nil),
	}
}

// -----------------------------------------------------------------------------
// ---------------- Submit transaction
// -----------------------------------------------------------------------------
// The request for transaction submission.
type SubmitTransactionRequest struct {
	// Transaction submitted by user.
	Transaction          *pbtypes.SignedTransaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *SubmitTransactionRequest) Reset()         { *m = SubmitTransactionRequest{} }
func (m *SubmitTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitTransactionRequest) ProtoMessage()    {}
func (*SubmitTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d13d6ff9aac892b4, []int{1}
}

func (m *SubmitTransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitTransactionRequest.Unmarshal(m, b)
}
func (m *SubmitTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitTransactionRequest.Marshal(b, m, deterministic)
}
func (m *SubmitTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitTransactionRequest.Merge(m, src)
}
func (m *SubmitTransactionRequest) XXX_Size() int {
	return xxx_messageInfo_SubmitTransactionRequest.Size(m)
}
func (m *SubmitTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitTransactionRequest proto.InternalMessageInfo

func (m *SubmitTransactionRequest) GetTransaction() *pbtypes.SignedTransaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

// AC response status containing code and optionally an error message.
type AdmissionControlStatus struct {
	Code                 AdmissionControlStatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=admission_control.AdmissionControlStatusCode" json:"code,omitempty"`
	Message              string                     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *AdmissionControlStatus) Reset()         { *m = AdmissionControlStatus{} }
func (m *AdmissionControlStatus) String() string { return proto.CompactTextString(m) }
func (*AdmissionControlStatus) ProtoMessage()    {}
func (*AdmissionControlStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_d13d6ff9aac892b4, []int{2}
}

func (m *AdmissionControlStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdmissionControlStatus.Unmarshal(m, b)
}
func (m *AdmissionControlStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdmissionControlStatus.Marshal(b, m, deterministic)
}
func (m *AdmissionControlStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdmissionControlStatus.Merge(m, src)
}
func (m *AdmissionControlStatus) XXX_Size() int {
	return xxx_messageInfo_AdmissionControlStatus.Size(m)
}
func (m *AdmissionControlStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_AdmissionControlStatus.DiscardUnknown(m)
}

var xxx_messageInfo_AdmissionControlStatus proto.InternalMessageInfo

func (m *AdmissionControlStatus) GetCode() AdmissionControlStatusCode {
	if m != nil {
		return m.Code
	}
	return AdmissionControlStatusCode_Accepted
}

func (m *AdmissionControlStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// The response for transaction submission.
//
// How does a client know if their transaction was included?
// A response from the transaction submission only means that the transaction
// was successfully added to mempool, but not that it is guaranteed to be
// included in the chain.  Each transaction should include an expiration time in
// the signed transaction.  Let's call this T0.  As a client, I submit my
// transaction to a validator. I now need to poll for the transaction I
// submitted.  I can use the query that takes my account and sequence number. If
// I receive back that the transaction is completed, I will verify the proofs to
// ensure that this is the transaction I expected.  If I receive a response that
// my transaction is not yet completed, I must check the latest timestamp in the
// ledgerInfo that I receive back from the query.  If this time is greater than
// T0, I can be certain that my transaction will never be included.  If this
// time is less than T0, I need to continue polling.
type SubmitTransactionResponse struct {
	// The status of a transaction submission can either be a VM status, or
	// some other admission control/mempool specific status e.g. Blacklisted.
	//
	// Types that are valid to be assigned to Status:
	//	*SubmitTransactionResponse_VmStatus
	//	*SubmitTransactionResponse_AcStatus
	//	*SubmitTransactionResponse_MempoolStatus
	Status isSubmitTransactionResponse_Status `protobuf_oneof:"status"`
	// Public key(id) of the validator that processed this transaction
	ValidatorId          []byte   `protobuf:"bytes,4,opt,name=validator_id,json=validatorId,proto3" json:"validator_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitTransactionResponse) Reset()         { *m = SubmitTransactionResponse{} }
func (m *SubmitTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*SubmitTransactionResponse) ProtoMessage()    {}
func (*SubmitTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d13d6ff9aac892b4, []int{3}
}

func (m *SubmitTransactionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitTransactionResponse.Unmarshal(m, b)
}
func (m *SubmitTransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitTransactionResponse.Marshal(b, m, deterministic)
}
func (m *SubmitTransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitTransactionResponse.Merge(m, src)
}
func (m *SubmitTransactionResponse) XXX_Size() int {
	return xxx_messageInfo_SubmitTransactionResponse.Size(m)
}
func (m *SubmitTransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitTransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitTransactionResponse proto.InternalMessageInfo

type isSubmitTransactionResponse_Status interface {
	isSubmitTransactionResponse_Status()
}

type SubmitTransactionResponse_VmStatus struct {
	VmStatus *pbtypes.VMStatus `protobuf:"bytes,1,opt,name=vm_status,json=vmStatus,proto3,oneof"`
}

type SubmitTransactionResponse_AcStatus struct {
	AcStatus *AdmissionControlStatus `protobuf:"bytes,2,opt,name=ac_status,json=acStatus,proto3,oneof"`
}

type SubmitTransactionResponse_MempoolStatus struct {
	MempoolStatus *pbtypes.MempoolStatus `protobuf:"bytes,3,opt,name=mempool_status,json=mempoolStatus,proto3,oneof"`
}

func (*SubmitTransactionResponse_VmStatus) isSubmitTransactionResponse_Status() {}

func (*SubmitTransactionResponse_AcStatus) isSubmitTransactionResponse_Status() {}

func (*SubmitTransactionResponse_MempoolStatus) isSubmitTransactionResponse_Status() {}

func (m *SubmitTransactionResponse) GetStatus() isSubmitTransactionResponse_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *SubmitTransactionResponse) GetVmStatus() *pbtypes.VMStatus {
	if x, ok := m.GetStatus().(*SubmitTransactionResponse_VmStatus); ok {
		return x.VmStatus
	}
	return nil
}

func (m *SubmitTransactionResponse) GetAcStatus() *AdmissionControlStatus {
	if x, ok := m.GetStatus().(*SubmitTransactionResponse_AcStatus); ok {
		return x.AcStatus
	}
	return nil
}

func (m *SubmitTransactionResponse) GetMempoolStatus() *pbtypes.MempoolStatus {
	if x, ok := m.GetStatus().(*SubmitTransactionResponse_MempoolStatus); ok {
		return x.MempoolStatus
	}
	return nil
}

func (m *SubmitTransactionResponse) GetValidatorId() []byte {
	if m != nil {
		return m.ValidatorId
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SubmitTransactionResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SubmitTransactionResponse_VmStatus)(nil),
		(*SubmitTransactionResponse_AcStatus)(nil),
		(*SubmitTransactionResponse_MempoolStatus)(nil),
	}
}

func init() {
	proto.RegisterEnum("admission_control.AdmissionControlStatusCode", AdmissionControlStatusCode_name, AdmissionControlStatusCode_value)
	proto.RegisterType((*AdmissionControlMsg)(nil), "admission_control.AdmissionControlMsg")
	proto.RegisterType((*SubmitTransactionRequest)(nil), "admission_control.SubmitTransactionRequest")
	proto.RegisterType((*AdmissionControlStatus)(nil), "admission_control.AdmissionControlStatus")
	proto.RegisterType((*SubmitTransactionResponse)(nil), "admission_control.SubmitTransactionResponse")
}

func init() {
	proto.RegisterFile("admission_control.proto", fileDescriptor_d13d6ff9aac892b4)
}

var fileDescriptor_d13d6ff9aac892b4 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4f, 0x6f, 0x12, 0x41,
	0x14, 0x67, 0xb1, 0xa9, 0xf0, 0xc0, 0x52, 0x46, 0xa2, 0x5b, 0xbc, 0x54, 0xbc, 0xb4, 0xd6, 0x2e,
	0x09, 0x1e, 0x8c, 0x26, 0x1e, 0xa0, 0x17, 0x9a, 0x94, 0xcb, 0x52, 0x7b, 0xf0, 0xb2, 0x19, 0x66,
	0x9e, 0xcb, 0xd8, 0x9d, 0x9d, 0x75, 0x66, 0xc0, 0x78, 0xf4, 0xd3, 0xfa, 0x29, 0x4c, 0x0c, 0xbb,
	0x03, 0x62, 0x01, 0xd3, 0x1e, 0x79, 0xef, 0xfd, 0xfe, 0xbc, 0xdf, 0x3c, 0x16, 0x9e, 0x53, 0x2e,
	0x85, 0x31, 0x42, 0xa5, 0x11, 0x53, 0xa9, 0xd5, 0x2a, 0x09, 0x32, 0xad, 0xac, 0x22, 0xcd, 0x8d,
	0x46, 0xbb, 0x15, 0xa3, 0x8d, 0xbe, 0x0b, 0x3b, 0x8d, 0x32, 0xad, 0xd4, 0x97, 0x62, 0xb0, 0xdd,
	0x92, 0x28, 0x33, 0xa5, 0x92, 0xc8, 0x58, 0x6a, 0x67, 0xc6, 0x55, 0x9b, 0x56, 0xd3, 0xd4, 0x50,
	0x66, 0x85, 0x4a, 0x5d, 0xa9, 0x31, 0x97, 0x11, 0x6a, 0xad, 0xb4, 0x9b, 0xe9, 0xfc, 0xf6, 0xe0,
	0x69, 0x7f, 0xa9, 0x72, 0x51, 0x88, 0x8c, 0x4c, 0x4c, 0x6e, 0xa1, 0x6d, 0x66, 0x13, 0x29, 0x6c,
	0xb4, 0x46, 0x12, 0x69, 0xfc, 0x36, 0x43, 0x63, 0x7d, 0xef, 0xd8, 0x3b, 0xa9, 0xf5, 0xce, 0x82,
	0x4d, 0xe3, 0xe3, 0x1c, 0x74, 0xfd, 0x17, 0x13, 0x16, 0x90, 0x61, 0x29, 0xf4, 0xcd, 0x8e, 0x1e,
	0x49, 0xe1, 0xc5, 0x56, 0x31, 0x93, 0xa9, 0xd4, 0xa0, 0x5f, 0xce, 0xd5, 0xde, 0xdc, 0x4f, 0xad,
	0xc0, 0x0c, 0x4b, 0xe1, 0x91, 0xd9, 0xd5, 0x1c, 0x54, 0xe1, 0xb1, 0x44, 0x63, 0x68, 0x8c, 0x9d,
	0x1b, 0xf0, 0x77, 0x59, 0x26, 0x1f, 0xa0, 0xb6, 0xe6, 0xc7, 0x2d, 0xed, 0x07, 0xf6, 0x47, 0x86,
	0x26, 0x18, 0x8b, 0x38, 0x45, 0xbe, 0x8e, 0x5a, 0x1f, 0xee, 0xcc, 0xe0, 0xd9, 0xdd, 0x58, 0xc7,
	0xf9, 0xdb, 0x90, 0x3e, 0xec, 0x31, 0xc5, 0x31, 0xa7, 0x3b, 0xe8, 0x9d, 0x6f, 0xd9, 0x6a, 0x3b,
	0xf0, 0x42, 0x71, 0x0c, 0x73, 0x28, 0xf1, 0x57, 0xfe, 0xf3, 0x6c, 0xaa, 0xe1, 0x6a, 0x9d, 0x9f,
	0x65, 0x38, 0xda, 0x19, 0x0a, 0x09, 0xa0, 0x3a, 0x97, 0xee, 0x46, 0xdc, 0x3a, 0x0d, 0xb7, 0xce,
	0xcd, 0xa8, 0x50, 0x19, 0x96, 0xc2, 0xca, 0x5c, 0x3a, 0xab, 0x43, 0xa8, 0x52, 0xb6, 0x9c, 0x2f,
	0x5e, 0xe1, 0xf4, 0xde, 0x7e, 0x17, 0x4c, 0x94, 0x39, 0xa6, 0x8f, 0x70, 0xf0, 0xef, 0x89, 0xfa,
	0x8f, 0x72, 0xba, 0x96, 0x93, 0x1f, 0x15, 0xcd, 0x15, 0xf2, 0x89, 0x5c, 0x2f, 0x90, 0x97, 0x50,
	0x9f, 0xd3, 0x44, 0x70, 0x6a, 0x95, 0x8e, 0x04, 0xf7, 0xf7, 0x8e, 0xbd, 0x93, 0x7a, 0x58, 0x5b,
	0xd5, 0x2e, 0xf9, 0xa0, 0x02, 0xfb, 0x05, 0xf3, 0xeb, 0x4b, 0x68, 0xef, 0x4e, 0x90, 0xd4, 0xa1,
	0xd2, 0x67, 0x0c, 0x33, 0x8b, 0xfc, 0xb0, 0x44, 0x1a, 0x50, 0x1b, 0x24, 0x94, 0xdd, 0x26, 0xc2,
	0x2c, 0x0a, 0xde, 0xa2, 0x1d, 0xe2, 0x57, 0x64, 0x8b, 0x5f, 0xe5, 0xde, 0x2f, 0x0f, 0x0e, 0xef,
	0x72, 0x91, 0x0c, 0x9a, 0x1b, 0x11, 0x93, 0x87, 0xfc, 0x17, 0xda, 0x0f, 0x3a, 0xe5, 0x4e, 0x89,
	0x50, 0x68, 0x7d, 0xca, 0x38, 0xb5, 0x78, 0xad, 0xae, 0xa8, 0x45, 0x63, 0xaf, 0x90, 0xc7, 0xa8,
	0x49, 0xc7, 0xa5, 0xb7, 0xad, 0xb9, 0xd4, 0x7a, 0xf5, 0xdf, 0x99, 0xa5, 0xc4, 0xe0, 0xec, 0xf3,
	0x69, 0x2c, 0xec, 0x74, 0x36, 0x09, 0x98, 0x92, 0x5d, 0x3b, 0xc5, 0x77, 0xbd, 0xf7, 0xdd, 0x58,
	0x9d, 0x27, 0x62, 0xa2, 0x69, 0x37, 0xc6, 0x14, 0x35, 0xb5, 0xc8, 0xbb, 0xd9, 0x84, 0xb2, 0xc9,
	0x7e, 0xfe, 0xed, 0x78, 0xfb, 0x27, 0x00, 0x00, 0xff, 0xff, 0xff, 0x73, 0x23, 0x84, 0xb9, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdmissionControlClient is the client API for AdmissionControl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdmissionControlClient interface {
	// Public API to submit transaction to a validator.
	SubmitTransaction(ctx context.Context, in *SubmitTransactionRequest, opts ...grpc.CallOption) (*SubmitTransactionResponse, error)
	// This API is used to update the client to the latest ledger version and
	// optionally also request 1..n other pieces of data.  This allows for batch
	// queries.  All queries return proofs that a client should check to validate
	// the data. Note that if a client only wishes to update to the latest
	// LedgerInfo and receive the proof of this latest version, they can simply
	// omit the requested_items (or pass an empty list)
	UpdateToLatestLedger(ctx context.Context, in *pbtypes.UpdateToLatestLedgerRequest, opts ...grpc.CallOption) (*pbtypes.UpdateToLatestLedgerResponse, error)
}

type admissionControlClient struct {
	cc grpc.ClientConnInterface
}

func NewAdmissionControlClient(cc grpc.ClientConnInterface) AdmissionControlClient {
	return &admissionControlClient{cc}
}

func (c *admissionControlClient) SubmitTransaction(ctx context.Context, in *SubmitTransactionRequest, opts ...grpc.CallOption) (*SubmitTransactionResponse, error) {
	out := new(SubmitTransactionResponse)
	err := c.cc.Invoke(ctx, "/admission_control.AdmissionControl/SubmitTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *admissionControlClient) UpdateToLatestLedger(ctx context.Context, in *pbtypes.UpdateToLatestLedgerRequest, opts ...grpc.CallOption) (*pbtypes.UpdateToLatestLedgerResponse, error) {
	out := new(pbtypes.UpdateToLatestLedgerResponse)
	err := c.cc.Invoke(ctx, "/admission_control.AdmissionControl/UpdateToLatestLedger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdmissionControlServer is the server API for AdmissionControl service.
type AdmissionControlServer interface {
	// Public API to submit transaction to a validator.
	SubmitTransaction(context.Context, *SubmitTransactionRequest) (*SubmitTransactionResponse, error)
	// This API is used to update the client to the latest ledger version and
	// optionally also request 1..n other pieces of data.  This allows for batch
	// queries.  All queries return proofs that a client should check to validate
	// the data. Note that if a client only wishes to update to the latest
	// LedgerInfo and receive the proof of this latest version, they can simply
	// omit the requested_items (or pass an empty list)
	UpdateToLatestLedger(context.Context, *pbtypes.UpdateToLatestLedgerRequest) (*pbtypes.UpdateToLatestLedgerResponse, error)
}

// UnimplementedAdmissionControlServer can be embedded to have forward compatible implementations.
type UnimplementedAdmissionControlServer struct {
}

func (*UnimplementedAdmissionControlServer) SubmitTransaction(ctx context.Context, req *SubmitTransactionRequest) (*SubmitTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTransaction not implemented")
}
func (*UnimplementedAdmissionControlServer) UpdateToLatestLedger(ctx context.Context, req *pbtypes.UpdateToLatestLedgerRequest) (*pbtypes.UpdateToLatestLedgerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateToLatestLedger not implemented")
}

func RegisterAdmissionControlServer(s *grpc.Server, srv AdmissionControlServer) {
	s.RegisterService(&_AdmissionControl_serviceDesc, srv)
}

func _AdmissionControl_SubmitTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdmissionControlServer).SubmitTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admission_control.AdmissionControl/SubmitTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdmissionControlServer).SubmitTransaction(ctx, req.(*SubmitTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdmissionControl_UpdateToLatestLedger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbtypes.UpdateToLatestLedgerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdmissionControlServer).UpdateToLatestLedger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admission_control.AdmissionControl/UpdateToLatestLedger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdmissionControlServer).UpdateToLatestLedger(ctx, req.(*pbtypes.UpdateToLatestLedgerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdmissionControl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "admission_control.AdmissionControl",
	HandlerType: (*AdmissionControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitTransaction",
			Handler:    _AdmissionControl_SubmitTransaction_Handler,
		},
		{
			MethodName: "UpdateToLatestLedger",
			Handler:    _AdmissionControl_UpdateToLatestLedger_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admission_control.proto",
}
