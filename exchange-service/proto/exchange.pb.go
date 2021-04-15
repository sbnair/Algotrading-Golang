// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/exchange.proto

package exchangepb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Exchange struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Exchange             string   `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	ExchangeName         string   `protobuf:"bytes,3,opt,name=exchange_name,json=exchangeName,proto3" json:"exchange_name,omitempty"`
	ExchangeType         string   `protobuf:"bytes,4,opt,name=exchange_type,json=exchangeType,proto3" json:"exchange_type,omitempty"`
	UserId               string   `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ApiKey               string   `protobuf:"bytes,6,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	ApiSecret            string   `protobuf:"bytes,7,opt,name=api_secret,json=apiSecret,proto3" json:"api_secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Exchange) Reset()         { *m = Exchange{} }
func (m *Exchange) String() string { return proto.CompactTextString(m) }
func (*Exchange) ProtoMessage()    {}
func (*Exchange) Descriptor() ([]byte, []int) {
	return fileDescriptor_95588e2623cfabeb, []int{0}
}

func (m *Exchange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Exchange.Unmarshal(m, b)
}
func (m *Exchange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Exchange.Marshal(b, m, deterministic)
}
func (m *Exchange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Exchange.Merge(m, src)
}
func (m *Exchange) XXX_Size() int {
	return xxx_messageInfo_Exchange.Size(m)
}
func (m *Exchange) XXX_DiscardUnknown() {
	xxx_messageInfo_Exchange.DiscardUnknown(m)
}

var xxx_messageInfo_Exchange proto.InternalMessageInfo

func (m *Exchange) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Exchange) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *Exchange) GetExchangeName() string {
	if m != nil {
		return m.ExchangeName
	}
	return ""
}

func (m *Exchange) GetExchangeType() string {
	if m != nil {
		return m.ExchangeType
	}
	return ""
}

func (m *Exchange) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Exchange) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func (m *Exchange) GetApiSecret() string {
	if m != nil {
		return m.ApiSecret
	}
	return ""
}

type CreateExchangeReq struct {
	Exchange             *Exchange `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateExchangeReq) Reset()         { *m = CreateExchangeReq{} }
func (m *CreateExchangeReq) String() string { return proto.CompactTextString(m) }
func (*CreateExchangeReq) ProtoMessage()    {}
func (*CreateExchangeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_95588e2623cfabeb, []int{1}
}

func (m *CreateExchangeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateExchangeReq.Unmarshal(m, b)
}
func (m *CreateExchangeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateExchangeReq.Marshal(b, m, deterministic)
}
func (m *CreateExchangeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateExchangeReq.Merge(m, src)
}
func (m *CreateExchangeReq) XXX_Size() int {
	return xxx_messageInfo_CreateExchangeReq.Size(m)
}
func (m *CreateExchangeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateExchangeReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateExchangeReq proto.InternalMessageInfo

func (m *CreateExchangeReq) GetExchange() *Exchange {
	if m != nil {
		return m.Exchange
	}
	return nil
}

type CreateExchangeRes struct {
	Exchange             *Exchange `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateExchangeRes) Reset()         { *m = CreateExchangeRes{} }
func (m *CreateExchangeRes) String() string { return proto.CompactTextString(m) }
func (*CreateExchangeRes) ProtoMessage()    {}
func (*CreateExchangeRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_95588e2623cfabeb, []int{2}
}

func (m *CreateExchangeRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateExchangeRes.Unmarshal(m, b)
}
func (m *CreateExchangeRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateExchangeRes.Marshal(b, m, deterministic)
}
func (m *CreateExchangeRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateExchangeRes.Merge(m, src)
}
func (m *CreateExchangeRes) XXX_Size() int {
	return xxx_messageInfo_CreateExchangeRes.Size(m)
}
func (m *CreateExchangeRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateExchangeRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateExchangeRes proto.InternalMessageInfo

func (m *CreateExchangeRes) GetExchange() *Exchange {
	if m != nil {
		return m.Exchange
	}
	return nil
}

type UpdateExchangeReq struct {
	Exchange             *Exchange `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateExchangeReq) Reset()         { *m = UpdateExchangeReq{} }
func (m *UpdateExchangeReq) String() string { return proto.CompactTextString(m) }
func (*UpdateExchangeReq) ProtoMessage()    {}
func (*UpdateExchangeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_95588e2623cfabeb, []int{3}
}

func (m *UpdateExchangeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateExchangeReq.Unmarshal(m, b)
}
func (m *UpdateExchangeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateExchangeReq.Marshal(b, m, deterministic)
}
func (m *UpdateExchangeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateExchangeReq.Merge(m, src)
}
func (m *UpdateExchangeReq) XXX_Size() int {
	return xxx_messageInfo_UpdateExchangeReq.Size(m)
}
func (m *UpdateExchangeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateExchangeReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateExchangeReq proto.InternalMessageInfo

func (m *UpdateExchangeReq) GetExchange() *Exchange {
	if m != nil {
		return m.Exchange
	}
	return nil
}

type UpdateExchangeRes struct {
	Exchange             *Exchange `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateExchangeRes) Reset()         { *m = UpdateExchangeRes{} }
func (m *UpdateExchangeRes) String() string { return proto.CompactTextString(m) }
func (*UpdateExchangeRes) ProtoMessage()    {}
func (*UpdateExchangeRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_95588e2623cfabeb, []int{4}
}

func (m *UpdateExchangeRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateExchangeRes.Unmarshal(m, b)
}
func (m *UpdateExchangeRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateExchangeRes.Marshal(b, m, deterministic)
}
func (m *UpdateExchangeRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateExchangeRes.Merge(m, src)
}
func (m *UpdateExchangeRes) XXX_Size() int {
	return xxx_messageInfo_UpdateExchangeRes.Size(m)
}
func (m *UpdateExchangeRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateExchangeRes.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateExchangeRes proto.InternalMessageInfo

func (m *UpdateExchangeRes) GetExchange() *Exchange {
	if m != nil {
		return m.Exchange
	}
	return nil
}

type ReadExchangeReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadExchangeReq) Reset()         { *m = ReadExchangeReq{} }
func (m *ReadExchangeReq) String() string { return proto.CompactTextString(m) }
func (*ReadExchangeReq) ProtoMessage()    {}
func (*ReadExchangeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_95588e2623cfabeb, []int{5}
}

func (m *ReadExchangeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadExchangeReq.Unmarshal(m, b)
}
func (m *ReadExchangeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadExchangeReq.Marshal(b, m, deterministic)
}
func (m *ReadExchangeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadExchangeReq.Merge(m, src)
}
func (m *ReadExchangeReq) XXX_Size() int {
	return xxx_messageInfo_ReadExchangeReq.Size(m)
}
func (m *ReadExchangeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadExchangeReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReadExchangeReq proto.InternalMessageInfo

func (m *ReadExchangeReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadExchangeRes struct {
	Exchange             *Exchange `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ReadExchangeRes) Reset()         { *m = ReadExchangeRes{} }
func (m *ReadExchangeRes) String() string { return proto.CompactTextString(m) }
func (*ReadExchangeRes) ProtoMessage()    {}
func (*ReadExchangeRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_95588e2623cfabeb, []int{6}
}

func (m *ReadExchangeRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadExchangeRes.Unmarshal(m, b)
}
func (m *ReadExchangeRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadExchangeRes.Marshal(b, m, deterministic)
}
func (m *ReadExchangeRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadExchangeRes.Merge(m, src)
}
func (m *ReadExchangeRes) XXX_Size() int {
	return xxx_messageInfo_ReadExchangeRes.Size(m)
}
func (m *ReadExchangeRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadExchangeRes.DiscardUnknown(m)
}

var xxx_messageInfo_ReadExchangeRes proto.InternalMessageInfo

func (m *ReadExchangeRes) GetExchange() *Exchange {
	if m != nil {
		return m.Exchange
	}
	return nil
}

type DeleteExchangeReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteExchangeReq) Reset()         { *m = DeleteExchangeReq{} }
func (m *DeleteExchangeReq) String() string { return proto.CompactTextString(m) }
func (*DeleteExchangeReq) ProtoMessage()    {}
func (*DeleteExchangeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_95588e2623cfabeb, []int{7}
}

func (m *DeleteExchangeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteExchangeReq.Unmarshal(m, b)
}
func (m *DeleteExchangeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteExchangeReq.Marshal(b, m, deterministic)
}
func (m *DeleteExchangeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteExchangeReq.Merge(m, src)
}
func (m *DeleteExchangeReq) XXX_Size() int {
	return xxx_messageInfo_DeleteExchangeReq.Size(m)
}
func (m *DeleteExchangeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteExchangeReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteExchangeReq proto.InternalMessageInfo

func (m *DeleteExchangeReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteExchangeRes struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteExchangeRes) Reset()         { *m = DeleteExchangeRes{} }
func (m *DeleteExchangeRes) String() string { return proto.CompactTextString(m) }
func (*DeleteExchangeRes) ProtoMessage()    {}
func (*DeleteExchangeRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_95588e2623cfabeb, []int{8}
}

func (m *DeleteExchangeRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteExchangeRes.Unmarshal(m, b)
}
func (m *DeleteExchangeRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteExchangeRes.Marshal(b, m, deterministic)
}
func (m *DeleteExchangeRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteExchangeRes.Merge(m, src)
}
func (m *DeleteExchangeRes) XXX_Size() int {
	return xxx_messageInfo_DeleteExchangeRes.Size(m)
}
func (m *DeleteExchangeRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteExchangeRes.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteExchangeRes proto.InternalMessageInfo

func (m *DeleteExchangeRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ListExchangeReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListExchangeReq) Reset()         { *m = ListExchangeReq{} }
func (m *ListExchangeReq) String() string { return proto.CompactTextString(m) }
func (*ListExchangeReq) ProtoMessage()    {}
func (*ListExchangeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_95588e2623cfabeb, []int{9}
}

func (m *ListExchangeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListExchangeReq.Unmarshal(m, b)
}
func (m *ListExchangeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListExchangeReq.Marshal(b, m, deterministic)
}
func (m *ListExchangeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListExchangeReq.Merge(m, src)
}
func (m *ListExchangeReq) XXX_Size() int {
	return xxx_messageInfo_ListExchangeReq.Size(m)
}
func (m *ListExchangeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListExchangeReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListExchangeReq proto.InternalMessageInfo

type ListExchangeRes struct {
	Exchange             *Exchange `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListExchangeRes) Reset()         { *m = ListExchangeRes{} }
func (m *ListExchangeRes) String() string { return proto.CompactTextString(m) }
func (*ListExchangeRes) ProtoMessage()    {}
func (*ListExchangeRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_95588e2623cfabeb, []int{10}
}

func (m *ListExchangeRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListExchangeRes.Unmarshal(m, b)
}
func (m *ListExchangeRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListExchangeRes.Marshal(b, m, deterministic)
}
func (m *ListExchangeRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListExchangeRes.Merge(m, src)
}
func (m *ListExchangeRes) XXX_Size() int {
	return xxx_messageInfo_ListExchangeRes.Size(m)
}
func (m *ListExchangeRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListExchangeRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListExchangeRes proto.InternalMessageInfo

func (m *ListExchangeRes) GetExchange() *Exchange {
	if m != nil {
		return m.Exchange
	}
	return nil
}

func init() {
	proto.RegisterType((*Exchange)(nil), "exchange.Exchange")
	proto.RegisterType((*CreateExchangeReq)(nil), "exchange.CreateExchangeReq")
	proto.RegisterType((*CreateExchangeRes)(nil), "exchange.CreateExchangeRes")
	proto.RegisterType((*UpdateExchangeReq)(nil), "exchange.UpdateExchangeReq")
	proto.RegisterType((*UpdateExchangeRes)(nil), "exchange.UpdateExchangeRes")
	proto.RegisterType((*ReadExchangeReq)(nil), "exchange.ReadExchangeReq")
	proto.RegisterType((*ReadExchangeRes)(nil), "exchange.ReadExchangeRes")
	proto.RegisterType((*DeleteExchangeReq)(nil), "exchange.DeleteExchangeReq")
	proto.RegisterType((*DeleteExchangeRes)(nil), "exchange.DeleteExchangeRes")
	proto.RegisterType((*ListExchangeReq)(nil), "exchange.ListExchangeReq")
	proto.RegisterType((*ListExchangeRes)(nil), "exchange.ListExchangeRes")
}

func init() {
	proto.RegisterFile("proto/exchange.proto", fileDescriptor_95588e2623cfabeb)
}

var fileDescriptor_95588e2623cfabeb = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0x4f, 0xc2, 0x40,
	0x14, 0x4c, 0x41, 0xf9, 0x78, 0xf2, 0x11, 0x36, 0x26, 0xae, 0x35, 0x26, 0x5a, 0x2e, 0x5e, 0x44,
	0x83, 0xbf, 0x40, 0xc1, 0x18, 0x3f, 0xe2, 0xa1, 0xe8, 0xc5, 0x0b, 0x59, 0xda, 0x17, 0xdd, 0x28,
	0xb0, 0x76, 0x8b, 0xb1, 0x77, 0xff, 0xa0, 0xff, 0xc8, 0x74, 0xb5, 0x65, 0xb7, 0x15, 0x12, 0xf4,
	0xc6, 0xbc, 0x99, 0x4c, 0xe7, 0xbd, 0xd9, 0x00, 0x9b, 0x22, 0x98, 0x86, 0xd3, 0x23, 0x7c, 0xf7,
	0x9e, 0xd8, 0xe4, 0x11, 0x3b, 0x0a, 0x92, 0x4a, 0x82, 0x9d, 0x4f, 0x0b, 0x2a, 0xe7, 0x3f, 0x80,
	0x34, 0xa0, 0xc0, 0x7d, 0x6a, 0xed, 0x59, 0x07, 0x55, 0xb7, 0xc0, 0x7d, 0x62, 0x43, 0x2a, 0xa4,
	0x05, 0x35, 0x4d, 0x31, 0x69, 0x43, 0x3d, 0xf9, 0x3d, 0x9c, 0xb0, 0x31, 0xd2, 0xa2, 0x12, 0xd4,
	0x92, 0xe1, 0x2d, 0x1b, 0x9b, 0xa2, 0x30, 0x12, 0x48, 0xd7, 0x4c, 0xd1, 0x5d, 0x24, 0x90, 0x6c,
	0x41, 0x79, 0x26, 0x31, 0x18, 0x72, 0x9f, 0xae, 0x2b, 0xba, 0x14, 0xc3, 0x4b, 0x3f, 0x26, 0x98,
	0xe0, 0xc3, 0x67, 0x8c, 0x68, 0xe9, 0x9b, 0x60, 0x82, 0x5f, 0x63, 0x44, 0x76, 0x01, 0x62, 0x42,
	0xa2, 0x17, 0x60, 0x48, 0xcb, 0x8a, 0xab, 0x32, 0xc1, 0x07, 0x6a, 0xe0, 0xf4, 0xa0, 0xd5, 0x0b,
	0x90, 0x85, 0x98, 0x2c, 0xe6, 0xe2, 0x2b, 0xe9, 0x68, 0xbb, 0xc4, 0x1b, 0x6e, 0x74, 0x49, 0x27,
	0xbd, 0x4a, 0x2a, 0x9c, 0x1f, 0xe6, 0x17, 0x13, 0xf9, 0x17, 0x93, 0x7b, 0xe1, 0xff, 0x3f, 0x49,
	0xd6, 0x64, 0xf5, 0x24, 0xfb, 0xd0, 0x74, 0x91, 0xf9, 0x7a, 0x8e, 0x4c, 0xdb, 0xce, 0x69, 0x56,
	0xb2, 0xfa, 0x57, 0xda, 0xd0, 0xea, 0xe3, 0x0b, 0x9a, 0xfb, 0x66, 0xbf, 0x73, 0x98, 0x17, 0x49,
	0x42, 0xa1, 0x2c, 0x67, 0x9e, 0x87, 0x52, 0x2a, 0x65, 0xc5, 0x4d, 0xa0, 0xd3, 0x82, 0xe6, 0x0d,
	0x97, 0xa1, 0xe6, 0x18, 0x27, 0x35, 0x47, 0x2b, 0x27, 0xed, 0x7e, 0x14, 0xa1, 0x99, 0x8c, 0x07,
	0x18, 0xbc, 0x71, 0x0f, 0xc9, 0x15, 0x34, 0xcc, 0xca, 0xc9, 0xce, 0xdc, 0x23, 0xf7, 0xa2, 0xec,
	0x25, 0xa4, 0x24, 0x7d, 0xa8, 0xe9, 0xc7, 0x24, 0xdb, 0x73, 0x71, 0xa6, 0x07, 0x7b, 0x21, 0x25,
	0xe3, 0x44, 0x66, 0xf5, 0x7a, 0xa2, 0xdc, 0xcb, 0xb2, 0x97, 0x90, 0xca, 0xcb, 0x3c, 0xbb, 0xee,
	0x95, 0x6b, 0xcd, 0x5e, 0x42, 0x4a, 0x72, 0x01, 0x75, 0xbd, 0x00, 0xa9, 0xaf, 0x97, 0x29, 0xcb,
	0x5e, 0x48, 0xc9, 0x63, 0xeb, 0xac, 0xf6, 0x00, 0x09, 0x2b, 0x46, 0xa3, 0x92, 0xfa, 0x77, 0x3a,
	0xf9, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x39, 0x8b, 0x52, 0x99, 0xb5, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ExchangeServiceClient is the client API for ExchangeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExchangeServiceClient interface {
	CreateExchange(ctx context.Context, in *CreateExchangeReq, opts ...grpc.CallOption) (*CreateExchangeRes, error)
	ReadExchange(ctx context.Context, in *ReadExchangeReq, opts ...grpc.CallOption) (*ReadExchangeRes, error)
	UpdateExchange(ctx context.Context, in *UpdateExchangeReq, opts ...grpc.CallOption) (*UpdateExchangeRes, error)
	DeleteExchange(ctx context.Context, in *DeleteExchangeReq, opts ...grpc.CallOption) (*DeleteExchangeRes, error)
	ListExchanges(ctx context.Context, in *ListExchangeReq, opts ...grpc.CallOption) (ExchangeService_ListExchangesClient, error)
}

type exchangeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExchangeServiceClient(cc grpc.ClientConnInterface) ExchangeServiceClient {
	return &exchangeServiceClient{cc}
}

func (c *exchangeServiceClient) CreateExchange(ctx context.Context, in *CreateExchangeReq, opts ...grpc.CallOption) (*CreateExchangeRes, error) {
	out := new(CreateExchangeRes)
	err := c.cc.Invoke(ctx, "/exchange.ExchangeService/CreateExchange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) ReadExchange(ctx context.Context, in *ReadExchangeReq, opts ...grpc.CallOption) (*ReadExchangeRes, error) {
	out := new(ReadExchangeRes)
	err := c.cc.Invoke(ctx, "/exchange.ExchangeService/ReadExchange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) UpdateExchange(ctx context.Context, in *UpdateExchangeReq, opts ...grpc.CallOption) (*UpdateExchangeRes, error) {
	out := new(UpdateExchangeRes)
	err := c.cc.Invoke(ctx, "/exchange.ExchangeService/UpdateExchange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) DeleteExchange(ctx context.Context, in *DeleteExchangeReq, opts ...grpc.CallOption) (*DeleteExchangeRes, error) {
	out := new(DeleteExchangeRes)
	err := c.cc.Invoke(ctx, "/exchange.ExchangeService/DeleteExchange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) ListExchanges(ctx context.Context, in *ListExchangeReq, opts ...grpc.CallOption) (ExchangeService_ListExchangesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ExchangeService_serviceDesc.Streams[0], "/exchange.ExchangeService/ListExchanges", opts...)
	if err != nil {
		return nil, err
	}
	x := &exchangeServiceListExchangesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExchangeService_ListExchangesClient interface {
	Recv() (*ListExchangeRes, error)
	grpc.ClientStream
}

type exchangeServiceListExchangesClient struct {
	grpc.ClientStream
}

func (x *exchangeServiceListExchangesClient) Recv() (*ListExchangeRes, error) {
	m := new(ListExchangeRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExchangeServiceServer is the server API for ExchangeService service.
type ExchangeServiceServer interface {
	CreateExchange(context.Context, *CreateExchangeReq) (*CreateExchangeRes, error)
	ReadExchange(context.Context, *ReadExchangeReq) (*ReadExchangeRes, error)
	UpdateExchange(context.Context, *UpdateExchangeReq) (*UpdateExchangeRes, error)
	DeleteExchange(context.Context, *DeleteExchangeReq) (*DeleteExchangeRes, error)
	ListExchanges(*ListExchangeReq, ExchangeService_ListExchangesServer) error
}

// UnimplementedExchangeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedExchangeServiceServer struct {
}

func (*UnimplementedExchangeServiceServer) CreateExchange(ctx context.Context, req *CreateExchangeReq) (*CreateExchangeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExchange not implemented")
}
func (*UnimplementedExchangeServiceServer) ReadExchange(ctx context.Context, req *ReadExchangeReq) (*ReadExchangeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadExchange not implemented")
}
func (*UnimplementedExchangeServiceServer) UpdateExchange(ctx context.Context, req *UpdateExchangeReq) (*UpdateExchangeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExchange not implemented")
}
func (*UnimplementedExchangeServiceServer) DeleteExchange(ctx context.Context, req *DeleteExchangeReq) (*DeleteExchangeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExchange not implemented")
}
func (*UnimplementedExchangeServiceServer) ListExchanges(req *ListExchangeReq, srv ExchangeService_ListExchangesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListExchanges not implemented")
}

func RegisterExchangeServiceServer(s *grpc.Server, srv ExchangeServiceServer) {
	s.RegisterService(&_ExchangeService_serviceDesc, srv)
}

func _ExchangeService_CreateExchange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExchangeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).CreateExchange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exchange.ExchangeService/CreateExchange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).CreateExchange(ctx, req.(*CreateExchangeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_ReadExchange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadExchangeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).ReadExchange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exchange.ExchangeService/ReadExchange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).ReadExchange(ctx, req.(*ReadExchangeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_UpdateExchange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateExchangeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).UpdateExchange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exchange.ExchangeService/UpdateExchange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).UpdateExchange(ctx, req.(*UpdateExchangeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_DeleteExchange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteExchangeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).DeleteExchange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exchange.ExchangeService/DeleteExchange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).DeleteExchange(ctx, req.(*DeleteExchangeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_ListExchanges_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListExchangeReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExchangeServiceServer).ListExchanges(m, &exchangeServiceListExchangesServer{stream})
}

type ExchangeService_ListExchangesServer interface {
	Send(*ListExchangeRes) error
	grpc.ServerStream
}

type exchangeServiceListExchangesServer struct {
	grpc.ServerStream
}

func (x *exchangeServiceListExchangesServer) Send(m *ListExchangeRes) error {
	return x.ServerStream.SendMsg(m)
}

var _ExchangeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "exchange.ExchangeService",
	HandlerType: (*ExchangeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExchange",
			Handler:    _ExchangeService_CreateExchange_Handler,
		},
		{
			MethodName: "ReadExchange",
			Handler:    _ExchangeService_ReadExchange_Handler,
		},
		{
			MethodName: "UpdateExchange",
			Handler:    _ExchangeService_UpdateExchange_Handler,
		},
		{
			MethodName: "DeleteExchange",
			Handler:    _ExchangeService_DeleteExchange_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListExchanges",
			Handler:       _ExchangeService_ListExchanges_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/exchange.proto",
}
