// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/price.proto

package pricepb

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

type ListMyPositionReq struct {
	ExchangeId           string   `protobuf:"bytes,1,opt,name=exchange_id,json=exchangeId,proto3" json:"exchange_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMyPositionReq) Reset()         { *m = ListMyPositionReq{} }
func (m *ListMyPositionReq) String() string { return proto.CompactTextString(m) }
func (*ListMyPositionReq) ProtoMessage()    {}
func (*ListMyPositionReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f96ae06f70243894, []int{0}
}

func (m *ListMyPositionReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMyPositionReq.Unmarshal(m, b)
}
func (m *ListMyPositionReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMyPositionReq.Marshal(b, m, deterministic)
}
func (m *ListMyPositionReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMyPositionReq.Merge(m, src)
}
func (m *ListMyPositionReq) XXX_Size() int {
	return xxx_messageInfo_ListMyPositionReq.Size(m)
}
func (m *ListMyPositionReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMyPositionReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListMyPositionReq proto.InternalMessageInfo

func (m *ListMyPositionReq) GetExchangeId() string {
	if m != nil {
		return m.ExchangeId
	}
	return ""
}

type ListMyPositionRes struct {
	Position             *Position `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListMyPositionRes) Reset()         { *m = ListMyPositionRes{} }
func (m *ListMyPositionRes) String() string { return proto.CompactTextString(m) }
func (*ListMyPositionRes) ProtoMessage()    {}
func (*ListMyPositionRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_f96ae06f70243894, []int{1}
}

func (m *ListMyPositionRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMyPositionRes.Unmarshal(m, b)
}
func (m *ListMyPositionRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMyPositionRes.Marshal(b, m, deterministic)
}
func (m *ListMyPositionRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMyPositionRes.Merge(m, src)
}
func (m *ListMyPositionRes) XXX_Size() int {
	return xxx_messageInfo_ListMyPositionRes.Size(m)
}
func (m *ListMyPositionRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMyPositionRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListMyPositionRes proto.InternalMessageInfo

func (m *ListMyPositionRes) GetPosition() *Position {
	if m != nil {
		return m.Position
	}
	return nil
}

type ListAssetsReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAssetsReq) Reset()         { *m = ListAssetsReq{} }
func (m *ListAssetsReq) String() string { return proto.CompactTextString(m) }
func (*ListAssetsReq) ProtoMessage()    {}
func (*ListAssetsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f96ae06f70243894, []int{2}
}

func (m *ListAssetsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAssetsReq.Unmarshal(m, b)
}
func (m *ListAssetsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAssetsReq.Marshal(b, m, deterministic)
}
func (m *ListAssetsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAssetsReq.Merge(m, src)
}
func (m *ListAssetsReq) XXX_Size() int {
	return xxx_messageInfo_ListAssetsReq.Size(m)
}
func (m *ListAssetsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAssetsReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListAssetsReq proto.InternalMessageInfo

type ListAssetsRes struct {
	Asset                *Asset   `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAssetsRes) Reset()         { *m = ListAssetsRes{} }
func (m *ListAssetsRes) String() string { return proto.CompactTextString(m) }
func (*ListAssetsRes) ProtoMessage()    {}
func (*ListAssetsRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_f96ae06f70243894, []int{3}
}

func (m *ListAssetsRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAssetsRes.Unmarshal(m, b)
}
func (m *ListAssetsRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAssetsRes.Marshal(b, m, deterministic)
}
func (m *ListAssetsRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAssetsRes.Merge(m, src)
}
func (m *ListAssetsRes) XXX_Size() int {
	return xxx_messageInfo_ListAssetsRes.Size(m)
}
func (m *ListAssetsRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAssetsRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListAssetsRes proto.InternalMessageInfo

func (m *ListAssetsRes) GetAsset() *Asset {
	if m != nil {
		return m.Asset
	}
	return nil
}

type ListAssetBySymbolReq struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAssetBySymbolReq) Reset()         { *m = ListAssetBySymbolReq{} }
func (m *ListAssetBySymbolReq) String() string { return proto.CompactTextString(m) }
func (*ListAssetBySymbolReq) ProtoMessage()    {}
func (*ListAssetBySymbolReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f96ae06f70243894, []int{4}
}

func (m *ListAssetBySymbolReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAssetBySymbolReq.Unmarshal(m, b)
}
func (m *ListAssetBySymbolReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAssetBySymbolReq.Marshal(b, m, deterministic)
}
func (m *ListAssetBySymbolReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAssetBySymbolReq.Merge(m, src)
}
func (m *ListAssetBySymbolReq) XXX_Size() int {
	return xxx_messageInfo_ListAssetBySymbolReq.Size(m)
}
func (m *ListAssetBySymbolReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAssetBySymbolReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListAssetBySymbolReq proto.InternalMessageInfo

func (m *ListAssetBySymbolReq) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

type ListAssetBySymbolRes struct {
	Asset                *Asset   `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAssetBySymbolRes) Reset()         { *m = ListAssetBySymbolRes{} }
func (m *ListAssetBySymbolRes) String() string { return proto.CompactTextString(m) }
func (*ListAssetBySymbolRes) ProtoMessage()    {}
func (*ListAssetBySymbolRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_f96ae06f70243894, []int{5}
}

func (m *ListAssetBySymbolRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAssetBySymbolRes.Unmarshal(m, b)
}
func (m *ListAssetBySymbolRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAssetBySymbolRes.Marshal(b, m, deterministic)
}
func (m *ListAssetBySymbolRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAssetBySymbolRes.Merge(m, src)
}
func (m *ListAssetBySymbolRes) XXX_Size() int {
	return xxx_messageInfo_ListAssetBySymbolRes.Size(m)
}
func (m *ListAssetBySymbolRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAssetBySymbolRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListAssetBySymbolRes proto.InternalMessageInfo

func (m *ListAssetBySymbolRes) GetAsset() *Asset {
	if m != nil {
		return m.Asset
	}
	return nil
}

type ListAssetByNameReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAssetByNameReq) Reset()         { *m = ListAssetByNameReq{} }
func (m *ListAssetByNameReq) String() string { return proto.CompactTextString(m) }
func (*ListAssetByNameReq) ProtoMessage()    {}
func (*ListAssetByNameReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f96ae06f70243894, []int{6}
}

func (m *ListAssetByNameReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAssetByNameReq.Unmarshal(m, b)
}
func (m *ListAssetByNameReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAssetByNameReq.Marshal(b, m, deterministic)
}
func (m *ListAssetByNameReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAssetByNameReq.Merge(m, src)
}
func (m *ListAssetByNameReq) XXX_Size() int {
	return xxx_messageInfo_ListAssetByNameReq.Size(m)
}
func (m *ListAssetByNameReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAssetByNameReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListAssetByNameReq proto.InternalMessageInfo

func (m *ListAssetByNameReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListAssetByNameRes struct {
	Asset                *Asset   `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAssetByNameRes) Reset()         { *m = ListAssetByNameRes{} }
func (m *ListAssetByNameRes) String() string { return proto.CompactTextString(m) }
func (*ListAssetByNameRes) ProtoMessage()    {}
func (*ListAssetByNameRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_f96ae06f70243894, []int{7}
}

func (m *ListAssetByNameRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAssetByNameRes.Unmarshal(m, b)
}
func (m *ListAssetByNameRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAssetByNameRes.Marshal(b, m, deterministic)
}
func (m *ListAssetByNameRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAssetByNameRes.Merge(m, src)
}
func (m *ListAssetByNameRes) XXX_Size() int {
	return xxx_messageInfo_ListAssetByNameRes.Size(m)
}
func (m *ListAssetByNameRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAssetByNameRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListAssetByNameRes proto.InternalMessageInfo

func (m *ListAssetByNameRes) GetAsset() *Asset {
	if m != nil {
		return m.Asset
	}
	return nil
}

type Position struct {
	AssetId              string   `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	Symbol               string   `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Exchange             string   `protobuf:"bytes,3,opt,name=exchange,proto3" json:"exchange,omitempty"`
	AssetClass           string   `protobuf:"bytes,4,opt,name=asset_class,json=assetClass,proto3" json:"asset_class,omitempty"`
	AccountId            string   `protobuf:"bytes,5,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	AvgEntryPrice        float64  `protobuf:"fixed64,6,opt,name=avg_entry_price,json=avgEntryPrice,proto3" json:"avg_entry_price,omitempty"`
	Qty                  float64  `protobuf:"fixed64,7,opt,name=qty,proto3" json:"qty,omitempty"`
	Side                 string   `protobuf:"bytes,8,opt,name=side,proto3" json:"side,omitempty"`
	MarketValue          float64  `protobuf:"fixed64,9,opt,name=market_value,json=marketValue,proto3" json:"market_value,omitempty"`
	CostBasis            float64  `protobuf:"fixed64,10,opt,name=cost_basis,json=costBasis,proto3" json:"cost_basis,omitempty"`
	UnrealizedPl         float64  `protobuf:"fixed64,11,opt,name=unrealized_pl,json=unrealizedPl,proto3" json:"unrealized_pl,omitempty"`
	UnrealizedPlpc       float64  `protobuf:"fixed64,12,opt,name=unrealized_plpc,json=unrealizedPlpc,proto3" json:"unrealized_plpc,omitempty"`
	CurrentPrice         float64  `protobuf:"fixed64,13,opt,name=current_price,json=currentPrice,proto3" json:"current_price,omitempty"`
	LastdayPrice         float64  `protobuf:"fixed64,14,opt,name=lastday_price,json=lastdayPrice,proto3" json:"lastday_price,omitempty"`
	ChangeToday          float64  `protobuf:"fixed64,15,opt,name=change_today,json=changeToday,proto3" json:"change_today,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_f96ae06f70243894, []int{8}
}

func (m *Position) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Position.Unmarshal(m, b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Position.Marshal(b, m, deterministic)
}
func (m *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(m, src)
}
func (m *Position) XXX_Size() int {
	return xxx_messageInfo_Position.Size(m)
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

func (m *Position) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Position) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *Position) GetAssetClass() string {
	if m != nil {
		return m.AssetClass
	}
	return ""
}

func (m *Position) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Position) GetAvgEntryPrice() float64 {
	if m != nil {
		return m.AvgEntryPrice
	}
	return 0
}

func (m *Position) GetQty() float64 {
	if m != nil {
		return m.Qty
	}
	return 0
}

func (m *Position) GetSide() string {
	if m != nil {
		return m.Side
	}
	return ""
}

func (m *Position) GetMarketValue() float64 {
	if m != nil {
		return m.MarketValue
	}
	return 0
}

func (m *Position) GetCostBasis() float64 {
	if m != nil {
		return m.CostBasis
	}
	return 0
}

func (m *Position) GetUnrealizedPl() float64 {
	if m != nil {
		return m.UnrealizedPl
	}
	return 0
}

func (m *Position) GetUnrealizedPlpc() float64 {
	if m != nil {
		return m.UnrealizedPlpc
	}
	return 0
}

func (m *Position) GetCurrentPrice() float64 {
	if m != nil {
		return m.CurrentPrice
	}
	return 0
}

func (m *Position) GetLastdayPrice() float64 {
	if m != nil {
		return m.LastdayPrice
	}
	return 0
}

func (m *Position) GetChangeToday() float64 {
	if m != nil {
		return m.ChangeToday
	}
	return 0
}

type Asset struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Exchange             string   `protobuf:"bytes,3,opt,name=exchange,proto3" json:"exchange,omitempty"`
	AssetClass           string   `protobuf:"bytes,4,opt,name=asset_class,json=assetClass,proto3" json:"asset_class,omitempty"`
	Symbol               string   `protobuf:"bytes,5,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Status               string   `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Tradable             bool     `protobuf:"varint,7,opt,name=tradable,proto3" json:"tradable,omitempty"`
	Marginable           bool     `protobuf:"varint,8,opt,name=marginable,proto3" json:"marginable,omitempty"`
	Shortable            bool     `protobuf:"varint,9,opt,name=shortable,proto3" json:"shortable,omitempty"`
	EasyToBorrow         bool     `protobuf:"varint,10,opt,name=easy_to_borrow,json=easyToBorrow,proto3" json:"easy_to_borrow,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_f96ae06f70243894, []int{9}
}

func (m *Asset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset.Unmarshal(m, b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return xxx_messageInfo_Asset.Size(m)
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Asset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Asset) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *Asset) GetAssetClass() string {
	if m != nil {
		return m.AssetClass
	}
	return ""
}

func (m *Asset) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Asset) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Asset) GetTradable() bool {
	if m != nil {
		return m.Tradable
	}
	return false
}

func (m *Asset) GetMarginable() bool {
	if m != nil {
		return m.Marginable
	}
	return false
}

func (m *Asset) GetShortable() bool {
	if m != nil {
		return m.Shortable
	}
	return false
}

func (m *Asset) GetEasyToBorrow() bool {
	if m != nil {
		return m.EasyToBorrow
	}
	return false
}

func init() {
	proto.RegisterType((*ListMyPositionReq)(nil), "price.ListMyPositionReq")
	proto.RegisterType((*ListMyPositionRes)(nil), "price.ListMyPositionRes")
	proto.RegisterType((*ListAssetsReq)(nil), "price.ListAssetsReq")
	proto.RegisterType((*ListAssetsRes)(nil), "price.ListAssetsRes")
	proto.RegisterType((*ListAssetBySymbolReq)(nil), "price.ListAssetBySymbolReq")
	proto.RegisterType((*ListAssetBySymbolRes)(nil), "price.ListAssetBySymbolRes")
	proto.RegisterType((*ListAssetByNameReq)(nil), "price.ListAssetByNameReq")
	proto.RegisterType((*ListAssetByNameRes)(nil), "price.ListAssetByNameRes")
	proto.RegisterType((*Position)(nil), "price.Position")
	proto.RegisterType((*Asset)(nil), "price.Asset")
}

func init() {
	proto.RegisterFile("proto/price.proto", fileDescriptor_f96ae06f70243894)
}

var fileDescriptor_f96ae06f70243894 = []byte{
	// 651 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x95, 0xb4, 0x69, 0xed, 0x89, 0x93, 0xd0, 0x55, 0x85, 0xdc, 0xf0, 0xd5, 0xba, 0x08,
	0x2a, 0x21, 0x95, 0xaa, 0xe5, 0x80, 0x7a, 0x82, 0x20, 0x84, 0x22, 0xf1, 0x51, 0xb9, 0x15, 0x07,
	0x2e, 0xd6, 0xc6, 0x5e, 0xa5, 0x16, 0x8e, 0xd7, 0xdd, 0xdd, 0x04, 0xcc, 0x93, 0xf0, 0x16, 0xbc,
	0x14, 0x0f, 0x82, 0x76, 0x76, 0xe3, 0x38, 0x34, 0x45, 0x95, 0xb8, 0xed, 0xfc, 0xe6, 0xbf, 0x33,
	0x9e, 0x99, 0xf5, 0xc0, 0x56, 0x21, 0xb8, 0xe2, 0xcf, 0x0b, 0x91, 0xc6, 0xec, 0x10, 0xcf, 0xa4,
	0x85, 0x46, 0xf0, 0x02, 0xb6, 0xde, 0xa7, 0x52, 0x7d, 0x28, 0xcf, 0xb8, 0x4c, 0x55, 0xca, 0xf3,
	0x90, 0x5d, 0x91, 0x47, 0xd0, 0x66, 0xdf, 0xe3, 0x4b, 0x9a, 0x8f, 0x59, 0x94, 0x26, 0x7e, 0x63,
	0xb7, 0x71, 0xe0, 0x86, 0x30, 0x47, 0xc3, 0x24, 0x78, 0x75, 0xfd, 0x96, 0x24, 0xcf, 0xc0, 0x29,
	0xac, 0x89, 0x57, 0xda, 0xc7, 0xbd, 0x43, 0x93, 0xb1, 0x52, 0x55, 0x82, 0xa0, 0x07, 0x1d, 0x1d,
	0xe1, 0xb5, 0x94, 0x4c, 0xc9, 0x90, 0x5d, 0x05, 0x27, 0xcb, 0x40, 0x92, 0x00, 0x5a, 0x54, 0x1b,
	0x36, 0x96, 0x67, 0x63, 0xa1, 0x20, 0x34, 0xae, 0xe0, 0x10, 0xb6, 0xab, 0x4b, 0x83, 0xf2, 0xbc,
	0x9c, 0x8c, 0x78, 0xa6, 0x0b, 0xb8, 0x0b, 0x1b, 0x12, 0x0d, 0xfb, 0xed, 0xd6, 0x0a, 0x4e, 0x57,
	0xea, 0x6f, 0x97, 0xeb, 0x00, 0x48, 0xed, 0xee, 0x47, 0x3a, 0x61, 0x3a, 0x13, 0x81, 0xf5, 0x9c,
	0x4e, 0x98, 0xcd, 0x83, 0xe7, 0xe0, 0xe5, 0x0a, 0xe5, 0xed, 0x72, 0xfc, 0x5e, 0x03, 0x67, 0xde,
	0x2c, 0xb2, 0x03, 0x0e, 0xd2, 0xc5, 0x08, 0x36, 0xd1, 0x1e, 0x26, 0xb5, 0xfa, 0x9a, 0xf5, 0xfa,
	0x48, 0x1f, 0x9c, 0xf9, 0x94, 0xfc, 0x35, 0xf4, 0x54, 0xb6, 0x1e, 0xaa, 0x09, 0x17, 0x67, 0x54,
	0x4a, 0x7f, 0xdd, 0x0c, 0x15, 0xd1, 0x1b, 0x4d, 0xc8, 0x03, 0x00, 0x1a, 0xc7, 0x7c, 0x9a, 0x63,
	0xc6, 0x16, 0xfa, 0x5d, 0x4b, 0x86, 0x09, 0x79, 0x02, 0x3d, 0x3a, 0x1b, 0x47, 0x2c, 0x57, 0xa2,
	0x8c, 0xf0, 0xdb, 0xfd, 0x8d, 0xdd, 0xc6, 0x41, 0x23, 0xec, 0xd0, 0xd9, 0xf8, 0xad, 0xa6, 0x67,
	0x1a, 0x92, 0x3b, 0xb0, 0x76, 0xa5, 0x4a, 0x7f, 0x13, 0x7d, 0xfa, 0xa8, 0x7b, 0x24, 0xd3, 0x84,
	0xf9, 0x8e, 0xe9, 0x91, 0x3e, 0x93, 0x3d, 0xf0, 0x26, 0x54, 0x7c, 0x65, 0x2a, 0x9a, 0xd1, 0x6c,
	0xca, 0x7c, 0x17, 0xe5, 0x6d, 0xc3, 0x3e, 0x6b, 0xa4, 0xbf, 0x27, 0xe6, 0x52, 0x45, 0x23, 0x2a,
	0x53, 0xe9, 0x03, 0x0a, 0x5c, 0x4d, 0x06, 0x1a, 0x90, 0x7d, 0xe8, 0x4c, 0x73, 0xc1, 0x68, 0x96,
	0xfe, 0x60, 0x49, 0x54, 0x64, 0x7e, 0x1b, 0x15, 0xde, 0x02, 0x9e, 0x65, 0xe4, 0x29, 0xf4, 0x96,
	0x44, 0x45, 0xec, 0x7b, 0x28, 0xeb, 0xd6, 0x65, 0x45, 0xac, 0xa3, 0xc5, 0x53, 0x21, 0x58, 0xae,
	0x6c, 0x6d, 0x1d, 0x13, 0xcd, 0x42, 0x53, 0xda, 0x3e, 0x74, 0x32, 0x2a, 0x55, 0x42, 0xe7, 0x0d,
	0xe8, 0x1a, 0x91, 0x85, 0x46, 0xb4, 0x07, 0x9e, 0xfd, 0x75, 0x14, 0x4f, 0x68, 0xe9, 0xf7, 0x4c,
	0x65, 0x86, 0x5d, 0x68, 0x14, 0xfc, 0x6c, 0x42, 0x0b, 0xe7, 0x4e, 0xba, 0xd0, 0xac, 0xa6, 0xdb,
	0x4c, 0x93, 0xea, 0x39, 0x35, 0x17, 0xcf, 0xe9, 0xff, 0x86, 0xba, 0x78, 0x29, 0xad, 0xa5, 0x97,
	0xa2, 0xb9, 0xa2, 0x6a, 0x2a, 0x71, 0x88, 0x9a, 0xa3, 0xa5, 0x93, 0x29, 0x41, 0x13, 0x3a, 0xca,
	0x18, 0x8e, 0xd0, 0x09, 0x2b, 0x9b, 0x3c, 0x04, 0x98, 0x50, 0x31, 0x4e, 0x73, 0xf4, 0x3a, 0xe8,
	0xad, 0x11, 0x72, 0x1f, 0x5c, 0x79, 0xc9, 0x85, 0x42, 0xb7, 0x8b, 0xee, 0x05, 0x20, 0x8f, 0xa1,
	0xcb, 0xa8, 0x2c, 0x23, 0xc5, 0xa3, 0x11, 0x17, 0x82, 0x7f, 0xc3, 0x91, 0x3a, 0xa1, 0xa7, 0xe9,
	0x05, 0x1f, 0x20, 0x3b, 0xfe, 0xd5, 0x04, 0x0f, 0xfb, 0x78, 0xce, 0xc4, 0x4c, 0xb7, 0xf3, 0x1d,
	0xf4, 0x96, 0x57, 0x8d, 0x24, 0xbe, 0xfd, 0x75, 0xae, 0x2d, 0xae, 0xfe, 0x4d, 0x1e, 0x79, 0xd4,
	0x20, 0xa7, 0x00, 0x8b, 0x05, 0x43, 0xb6, 0x6b, 0xca, 0x6a, 0x09, 0xf5, 0x57, 0x51, 0x7d, 0xf7,
	0x93, 0xd9, 0x77, 0x4b, 0x7b, 0x83, 0xdc, 0xfb, 0x5b, 0x5c, 0xdb, 0x40, 0xfd, 0x7f, 0x38, 0x75,
	0xc0, 0xa1, 0xa9, 0xaa, 0xb6, 0x22, 0xc8, 0xce, 0xf5, 0x1b, 0x76, 0xc9, 0xf4, 0x6f, 0x74, 0xc9,
	0xa3, 0xc6, 0xc0, 0xfd, 0xb2, 0x89, 0xde, 0x62, 0x34, 0xda, 0xc0, 0xd5, 0x7e, 0xf2, 0x27, 0x00,
	0x00, 0xff, 0xff, 0xca, 0xdc, 0xee, 0x6f, 0xef, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PriceServiceClient is the client API for PriceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PriceServiceClient interface {
	ListMyPositions(ctx context.Context, in *ListMyPositionReq, opts ...grpc.CallOption) (PriceService_ListMyPositionsClient, error)
	ListAssets(ctx context.Context, in *ListAssetsReq, opts ...grpc.CallOption) (PriceService_ListAssetsClient, error)
	ListAssetBySymbol(ctx context.Context, in *ListAssetBySymbolReq, opts ...grpc.CallOption) (PriceService_ListAssetBySymbolClient, error)
	ListAssetByName(ctx context.Context, in *ListAssetByNameReq, opts ...grpc.CallOption) (PriceService_ListAssetByNameClient, error)
}

type priceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPriceServiceClient(cc grpc.ClientConnInterface) PriceServiceClient {
	return &priceServiceClient{cc}
}

func (c *priceServiceClient) ListMyPositions(ctx context.Context, in *ListMyPositionReq, opts ...grpc.CallOption) (PriceService_ListMyPositionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PriceService_serviceDesc.Streams[0], "/price.PriceService/ListMyPositions", opts...)
	if err != nil {
		return nil, err
	}
	x := &priceServiceListMyPositionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PriceService_ListMyPositionsClient interface {
	Recv() (*ListMyPositionRes, error)
	grpc.ClientStream
}

type priceServiceListMyPositionsClient struct {
	grpc.ClientStream
}

func (x *priceServiceListMyPositionsClient) Recv() (*ListMyPositionRes, error) {
	m := new(ListMyPositionRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *priceServiceClient) ListAssets(ctx context.Context, in *ListAssetsReq, opts ...grpc.CallOption) (PriceService_ListAssetsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PriceService_serviceDesc.Streams[1], "/price.PriceService/ListAssets", opts...)
	if err != nil {
		return nil, err
	}
	x := &priceServiceListAssetsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PriceService_ListAssetsClient interface {
	Recv() (*ListAssetsRes, error)
	grpc.ClientStream
}

type priceServiceListAssetsClient struct {
	grpc.ClientStream
}

func (x *priceServiceListAssetsClient) Recv() (*ListAssetsRes, error) {
	m := new(ListAssetsRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *priceServiceClient) ListAssetBySymbol(ctx context.Context, in *ListAssetBySymbolReq, opts ...grpc.CallOption) (PriceService_ListAssetBySymbolClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PriceService_serviceDesc.Streams[2], "/price.PriceService/ListAssetBySymbol", opts...)
	if err != nil {
		return nil, err
	}
	x := &priceServiceListAssetBySymbolClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PriceService_ListAssetBySymbolClient interface {
	Recv() (*ListAssetBySymbolRes, error)
	grpc.ClientStream
}

type priceServiceListAssetBySymbolClient struct {
	grpc.ClientStream
}

func (x *priceServiceListAssetBySymbolClient) Recv() (*ListAssetBySymbolRes, error) {
	m := new(ListAssetBySymbolRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *priceServiceClient) ListAssetByName(ctx context.Context, in *ListAssetByNameReq, opts ...grpc.CallOption) (PriceService_ListAssetByNameClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PriceService_serviceDesc.Streams[3], "/price.PriceService/ListAssetByName", opts...)
	if err != nil {
		return nil, err
	}
	x := &priceServiceListAssetByNameClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PriceService_ListAssetByNameClient interface {
	Recv() (*ListAssetByNameRes, error)
	grpc.ClientStream
}

type priceServiceListAssetByNameClient struct {
	grpc.ClientStream
}

func (x *priceServiceListAssetByNameClient) Recv() (*ListAssetByNameRes, error) {
	m := new(ListAssetByNameRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PriceServiceServer is the server API for PriceService service.
type PriceServiceServer interface {
	ListMyPositions(*ListMyPositionReq, PriceService_ListMyPositionsServer) error
	ListAssets(*ListAssetsReq, PriceService_ListAssetsServer) error
	ListAssetBySymbol(*ListAssetBySymbolReq, PriceService_ListAssetBySymbolServer) error
	ListAssetByName(*ListAssetByNameReq, PriceService_ListAssetByNameServer) error
}

// UnimplementedPriceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPriceServiceServer struct {
}

func (*UnimplementedPriceServiceServer) ListMyPositions(req *ListMyPositionReq, srv PriceService_ListMyPositionsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListMyPositions not implemented")
}
func (*UnimplementedPriceServiceServer) ListAssets(req *ListAssetsReq, srv PriceService_ListAssetsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAssets not implemented")
}
func (*UnimplementedPriceServiceServer) ListAssetBySymbol(req *ListAssetBySymbolReq, srv PriceService_ListAssetBySymbolServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAssetBySymbol not implemented")
}
func (*UnimplementedPriceServiceServer) ListAssetByName(req *ListAssetByNameReq, srv PriceService_ListAssetByNameServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAssetByName not implemented")
}

func RegisterPriceServiceServer(s *grpc.Server, srv PriceServiceServer) {
	s.RegisterService(&_PriceService_serviceDesc, srv)
}

func _PriceService_ListMyPositions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListMyPositionReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PriceServiceServer).ListMyPositions(m, &priceServiceListMyPositionsServer{stream})
}

type PriceService_ListMyPositionsServer interface {
	Send(*ListMyPositionRes) error
	grpc.ServerStream
}

type priceServiceListMyPositionsServer struct {
	grpc.ServerStream
}

func (x *priceServiceListMyPositionsServer) Send(m *ListMyPositionRes) error {
	return x.ServerStream.SendMsg(m)
}

func _PriceService_ListAssets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListAssetsReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PriceServiceServer).ListAssets(m, &priceServiceListAssetsServer{stream})
}

type PriceService_ListAssetsServer interface {
	Send(*ListAssetsRes) error
	grpc.ServerStream
}

type priceServiceListAssetsServer struct {
	grpc.ServerStream
}

func (x *priceServiceListAssetsServer) Send(m *ListAssetsRes) error {
	return x.ServerStream.SendMsg(m)
}

func _PriceService_ListAssetBySymbol_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListAssetBySymbolReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PriceServiceServer).ListAssetBySymbol(m, &priceServiceListAssetBySymbolServer{stream})
}

type PriceService_ListAssetBySymbolServer interface {
	Send(*ListAssetBySymbolRes) error
	grpc.ServerStream
}

type priceServiceListAssetBySymbolServer struct {
	grpc.ServerStream
}

func (x *priceServiceListAssetBySymbolServer) Send(m *ListAssetBySymbolRes) error {
	return x.ServerStream.SendMsg(m)
}

func _PriceService_ListAssetByName_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListAssetByNameReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PriceServiceServer).ListAssetByName(m, &priceServiceListAssetByNameServer{stream})
}

type PriceService_ListAssetByNameServer interface {
	Send(*ListAssetByNameRes) error
	grpc.ServerStream
}

type priceServiceListAssetByNameServer struct {
	grpc.ServerStream
}

func (x *priceServiceListAssetByNameServer) Send(m *ListAssetByNameRes) error {
	return x.ServerStream.SendMsg(m)
}

var _PriceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "price.PriceService",
	HandlerType: (*PriceServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListMyPositions",
			Handler:       _PriceService_ListMyPositions_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListAssets",
			Handler:       _PriceService_ListAssets_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListAssetBySymbol",
			Handler:       _PriceService_ListAssetBySymbol_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListAssetByName",
			Handler:       _PriceService_ListAssetByName_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/price.proto",
}
