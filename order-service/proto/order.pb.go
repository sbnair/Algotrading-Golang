// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/order.proto

package orderpb

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

type Side int32

const (
	Side_invalid_side Side = 0
	Side_buy          Side = 1
	Side_sell         Side = 2
)

var Side_name = map[int32]string{
	0: "invalid_side",
	1: "buy",
	2: "sell",
}

var Side_value = map[string]int32{
	"invalid_side": 0,
	"buy":          1,
	"sell":         2,
}

func (x Side) String() string {
	return proto.EnumName(Side_name, int32(x))
}

func (Side) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f65b0626cc3aada8, []int{0}
}

type OrderClass int32

const (
	OrderClass_invalid_order_class OrderClass = 0
	OrderClass_bracket             OrderClass = 1
	OrderClass_oto                 OrderClass = 2
	OrderClass_oco                 OrderClass = 3
	OrderClass_simple              OrderClass = 4
)

var OrderClass_name = map[int32]string{
	0: "invalid_order_class",
	1: "bracket",
	2: "oto",
	3: "oco",
	4: "simple",
}

var OrderClass_value = map[string]int32{
	"invalid_order_class": 0,
	"bracket":             1,
	"oto":                 2,
	"oco":                 3,
	"simple":              4,
}

func (x OrderClass) String() string {
	return proto.EnumName(OrderClass_name, int32(x))
}

func (OrderClass) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f65b0626cc3aada8, []int{1}
}

type TimeInForce int32

const (
	TimeInForce_invalid_time_in_force TimeInForce = 0
	TimeInForce_day                   TimeInForce = 1
	TimeInForce_gtc                   TimeInForce = 2
	TimeInForce_opg                   TimeInForce = 3
	TimeInForce_ioc                   TimeInForce = 4
	TimeInForce_fok                   TimeInForce = 5
	TimeInForce_gtx                   TimeInForce = 6
	TimeInForce_gtd                   TimeInForce = 7
	TimeInForce_cls                   TimeInForce = 8
)

var TimeInForce_name = map[int32]string{
	0: "invalid_time_in_force",
	1: "day",
	2: "gtc",
	3: "opg",
	4: "ioc",
	5: "fok",
	6: "gtx",
	7: "gtd",
	8: "cls",
}

var TimeInForce_value = map[string]int32{
	"invalid_time_in_force": 0,
	"day":                   1,
	"gtc":                   2,
	"opg":                   3,
	"ioc":                   4,
	"fok":                   5,
	"gtx":                   6,
	"gtd":                   7,
	"cls":                   8,
}

func (x TimeInForce) String() string {
	return proto.EnumName(TimeInForce_name, int32(x))
}

func (TimeInForce) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f65b0626cc3aada8, []int{2}
}

type OrderType int32

const (
	OrderType_invalid_order_type OrderType = 0
	OrderType_market             OrderType = 1
	OrderType_limit              OrderType = 2
	OrderType_stop               OrderType = 3
	OrderType_stop_limit         OrderType = 4
	OrderType_trailing_stop      OrderType = 5
)

var OrderType_name = map[int32]string{
	0: "invalid_order_type",
	1: "market",
	2: "limit",
	3: "stop",
	4: "stop_limit",
	5: "trailing_stop",
}

var OrderType_value = map[string]int32{
	"invalid_order_type": 0,
	"market":             1,
	"limit":              2,
	"stop":               3,
	"stop_limit":         4,
	"trailing_stop":      5,
}

func (x OrderType) String() string {
	return proto.EnumName(OrderType_name, int32(x))
}

func (OrderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f65b0626cc3aada8, []int{3}
}

type PlaceOrderRes struct {
	Order                *Order   `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlaceOrderRes) Reset()         { *m = PlaceOrderRes{} }
func (m *PlaceOrderRes) String() string { return proto.CompactTextString(m) }
func (*PlaceOrderRes) ProtoMessage()    {}
func (*PlaceOrderRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_f65b0626cc3aada8, []int{0}
}

func (m *PlaceOrderRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlaceOrderRes.Unmarshal(m, b)
}
func (m *PlaceOrderRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlaceOrderRes.Marshal(b, m, deterministic)
}
func (m *PlaceOrderRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlaceOrderRes.Merge(m, src)
}
func (m *PlaceOrderRes) XXX_Size() int {
	return xxx_messageInfo_PlaceOrderRes.Size(m)
}
func (m *PlaceOrderRes) XXX_DiscardUnknown() {
	xxx_messageInfo_PlaceOrderRes.DiscardUnknown(m)
}

var xxx_messageInfo_PlaceOrderRes proto.InternalMessageInfo

func (m *PlaceOrderRes) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type PlaceOrderReq struct {
	ExchangeId           string      `protobuf:"bytes,1,opt,name=exchange_id,json=exchangeId,proto3" json:"exchange_id,omitempty"`
	Symbol               string      `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Qty                  float64     `protobuf:"fixed64,3,opt,name=qty,proto3" json:"qty,omitempty"`
	Notional             float64     `protobuf:"fixed64,4,opt,name=notional,proto3" json:"notional,omitempty"`
	Side                 Side        `protobuf:"varint,5,opt,name=side,proto3,enum=order.Side" json:"side,omitempty"`
	OrderType            OrderType   `protobuf:"varint,6,opt,name=order_type,json=orderType,proto3,enum=order.OrderType" json:"order_type,omitempty"`
	TimeInForce          TimeInForce `protobuf:"varint,7,opt,name=time_in_force,json=timeInForce,proto3,enum=order.TimeInForce" json:"time_in_force,omitempty"`
	LimitPrice           float64     `protobuf:"fixed64,8,opt,name=limit_price,json=limitPrice,proto3" json:"limit_price,omitempty"`
	ExtendedHours        bool        `protobuf:"varint,9,opt,name=extended_hours,json=extendedHours,proto3" json:"extended_hours,omitempty"`
	StopPrice            float64     `protobuf:"fixed64,10,opt,name=stop_price,json=stopPrice,proto3" json:"stop_price,omitempty"`
	ClientOrderId        string      `protobuf:"bytes,11,opt,name=client_order_id,json=clientOrderId,proto3" json:"client_order_id,omitempty"`
	OrderClass           OrderClass  `protobuf:"varint,12,opt,name=order_class,json=orderClass,proto3,enum=order.OrderClass" json:"order_class,omitempty"`
	TakeProfit           *TakeProfit `protobuf:"bytes,13,opt,name=take_profit,json=takeProfit,proto3" json:"take_profit,omitempty"`
	StopLoss             *StopLoss   `protobuf:"bytes,14,opt,name=stop_loss,json=stopLoss,proto3" json:"stop_loss,omitempty"`
	TrailPrice           float64     `protobuf:"fixed64,15,opt,name=trail_price,json=trailPrice,proto3" json:"trail_price,omitempty"`
	TrailPercent         float64     `protobuf:"fixed64,16,opt,name=trail_percent,json=trailPercent,proto3" json:"trail_percent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PlaceOrderReq) Reset()         { *m = PlaceOrderReq{} }
func (m *PlaceOrderReq) String() string { return proto.CompactTextString(m) }
func (*PlaceOrderReq) ProtoMessage()    {}
func (*PlaceOrderReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f65b0626cc3aada8, []int{1}
}

func (m *PlaceOrderReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlaceOrderReq.Unmarshal(m, b)
}
func (m *PlaceOrderReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlaceOrderReq.Marshal(b, m, deterministic)
}
func (m *PlaceOrderReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlaceOrderReq.Merge(m, src)
}
func (m *PlaceOrderReq) XXX_Size() int {
	return xxx_messageInfo_PlaceOrderReq.Size(m)
}
func (m *PlaceOrderReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PlaceOrderReq.DiscardUnknown(m)
}

var xxx_messageInfo_PlaceOrderReq proto.InternalMessageInfo

func (m *PlaceOrderReq) GetExchangeId() string {
	if m != nil {
		return m.ExchangeId
	}
	return ""
}

func (m *PlaceOrderReq) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *PlaceOrderReq) GetQty() float64 {
	if m != nil {
		return m.Qty
	}
	return 0
}

func (m *PlaceOrderReq) GetNotional() float64 {
	if m != nil {
		return m.Notional
	}
	return 0
}

func (m *PlaceOrderReq) GetSide() Side {
	if m != nil {
		return m.Side
	}
	return Side_invalid_side
}

func (m *PlaceOrderReq) GetOrderType() OrderType {
	if m != nil {
		return m.OrderType
	}
	return OrderType_invalid_order_type
}

func (m *PlaceOrderReq) GetTimeInForce() TimeInForce {
	if m != nil {
		return m.TimeInForce
	}
	return TimeInForce_invalid_time_in_force
}

func (m *PlaceOrderReq) GetLimitPrice() float64 {
	if m != nil {
		return m.LimitPrice
	}
	return 0
}

func (m *PlaceOrderReq) GetExtendedHours() bool {
	if m != nil {
		return m.ExtendedHours
	}
	return false
}

func (m *PlaceOrderReq) GetStopPrice() float64 {
	if m != nil {
		return m.StopPrice
	}
	return 0
}

func (m *PlaceOrderReq) GetClientOrderId() string {
	if m != nil {
		return m.ClientOrderId
	}
	return ""
}

func (m *PlaceOrderReq) GetOrderClass() OrderClass {
	if m != nil {
		return m.OrderClass
	}
	return OrderClass_invalid_order_class
}

func (m *PlaceOrderReq) GetTakeProfit() *TakeProfit {
	if m != nil {
		return m.TakeProfit
	}
	return nil
}

func (m *PlaceOrderReq) GetStopLoss() *StopLoss {
	if m != nil {
		return m.StopLoss
	}
	return nil
}

func (m *PlaceOrderReq) GetTrailPrice() float64 {
	if m != nil {
		return m.TrailPrice
	}
	return 0
}

func (m *PlaceOrderReq) GetTrailPercent() float64 {
	if m != nil {
		return m.TrailPercent
	}
	return 0
}

type TakeProfit struct {
	LimitPrice           float64  `protobuf:"fixed64,1,opt,name=limit_price,json=limitPrice,proto3" json:"limit_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TakeProfit) Reset()         { *m = TakeProfit{} }
func (m *TakeProfit) String() string { return proto.CompactTextString(m) }
func (*TakeProfit) ProtoMessage()    {}
func (*TakeProfit) Descriptor() ([]byte, []int) {
	return fileDescriptor_f65b0626cc3aada8, []int{2}
}

func (m *TakeProfit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TakeProfit.Unmarshal(m, b)
}
func (m *TakeProfit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TakeProfit.Marshal(b, m, deterministic)
}
func (m *TakeProfit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TakeProfit.Merge(m, src)
}
func (m *TakeProfit) XXX_Size() int {
	return xxx_messageInfo_TakeProfit.Size(m)
}
func (m *TakeProfit) XXX_DiscardUnknown() {
	xxx_messageInfo_TakeProfit.DiscardUnknown(m)
}

var xxx_messageInfo_TakeProfit proto.InternalMessageInfo

func (m *TakeProfit) GetLimitPrice() float64 {
	if m != nil {
		return m.LimitPrice
	}
	return 0
}

type StopLoss struct {
	LimitPrice           float64  `protobuf:"fixed64,1,opt,name=limit_price,json=limitPrice,proto3" json:"limit_price,omitempty"`
	StopPrice            float64  `protobuf:"fixed64,2,opt,name=stop_price,json=stopPrice,proto3" json:"stop_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopLoss) Reset()         { *m = StopLoss{} }
func (m *StopLoss) String() string { return proto.CompactTextString(m) }
func (*StopLoss) ProtoMessage()    {}
func (*StopLoss) Descriptor() ([]byte, []int) {
	return fileDescriptor_f65b0626cc3aada8, []int{3}
}

func (m *StopLoss) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopLoss.Unmarshal(m, b)
}
func (m *StopLoss) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopLoss.Marshal(b, m, deterministic)
}
func (m *StopLoss) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopLoss.Merge(m, src)
}
func (m *StopLoss) XXX_Size() int {
	return xxx_messageInfo_StopLoss.Size(m)
}
func (m *StopLoss) XXX_DiscardUnknown() {
	xxx_messageInfo_StopLoss.DiscardUnknown(m)
}

var xxx_messageInfo_StopLoss proto.InternalMessageInfo

func (m *StopLoss) GetLimitPrice() float64 {
	if m != nil {
		return m.LimitPrice
	}
	return 0
}

func (m *StopLoss) GetStopPrice() float64 {
	if m != nil {
		return m.StopPrice
	}
	return 0
}

type ListOrdersReq struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOrdersReq) Reset()         { *m = ListOrdersReq{} }
func (m *ListOrdersReq) String() string { return proto.CompactTextString(m) }
func (*ListOrdersReq) ProtoMessage()    {}
func (*ListOrdersReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f65b0626cc3aada8, []int{4}
}

func (m *ListOrdersReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOrdersReq.Unmarshal(m, b)
}
func (m *ListOrdersReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOrdersReq.Marshal(b, m, deterministic)
}
func (m *ListOrdersReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOrdersReq.Merge(m, src)
}
func (m *ListOrdersReq) XXX_Size() int {
	return xxx_messageInfo_ListOrdersReq.Size(m)
}
func (m *ListOrdersReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOrdersReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListOrdersReq proto.InternalMessageInfo

func (m *ListOrdersReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type ListOrdersRes struct {
	Order                *Order   `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOrdersRes) Reset()         { *m = ListOrdersRes{} }
func (m *ListOrdersRes) String() string { return proto.CompactTextString(m) }
func (*ListOrdersRes) ProtoMessage()    {}
func (*ListOrdersRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_f65b0626cc3aada8, []int{5}
}

func (m *ListOrdersRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOrdersRes.Unmarshal(m, b)
}
func (m *ListOrdersRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOrdersRes.Marshal(b, m, deterministic)
}
func (m *ListOrdersRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOrdersRes.Merge(m, src)
}
func (m *ListOrdersRes) XXX_Size() int {
	return xxx_messageInfo_ListOrdersRes.Size(m)
}
func (m *ListOrdersRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOrdersRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListOrdersRes proto.InternalMessageInfo

func (m *ListOrdersRes) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type Order struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientOrderId        string   `protobuf:"bytes,2,opt,name=client_order_id,json=clientOrderId,proto3" json:"client_order_id,omitempty"`
	CreatedAt            string   `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	SubmittedAt          string   `protobuf:"bytes,5,opt,name=submitted_at,json=submittedAt,proto3" json:"submitted_at,omitempty"`
	FilledAt             string   `protobuf:"bytes,6,opt,name=filled_at,json=filledAt,proto3" json:"filled_at,omitempty"`
	ExpiredAt            string   `protobuf:"bytes,7,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	CanceledAt           string   `protobuf:"bytes,8,opt,name=canceled_at,json=canceledAt,proto3" json:"canceled_at,omitempty"`
	FailedAt             string   `protobuf:"bytes,9,opt,name=failed_at,json=failedAt,proto3" json:"failed_at,omitempty"`
	ReplacedAt           string   `protobuf:"bytes,10,opt,name=replaced_at,json=replacedAt,proto3" json:"replaced_at,omitempty"`
	Replaces             string   `protobuf:"bytes,11,opt,name=replaces,proto3" json:"replaces,omitempty"`
	ReplacedBy           string   `protobuf:"bytes,12,opt,name=replaced_by,json=replacedBy,proto3" json:"replaced_by,omitempty"`
	AssetId              string   `protobuf:"bytes,13,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	Symbol               string   `protobuf:"bytes,14,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Exchange             string   `protobuf:"bytes,15,opt,name=exchange,proto3" json:"exchange,omitempty"`
	AssetClass           string   `protobuf:"bytes,16,opt,name=asset_class,json=assetClass,proto3" json:"asset_class,omitempty"`
	Qty                  float64  `protobuf:"fixed64,17,opt,name=qty,proto3" json:"qty,omitempty"`
	Notional             float64  `protobuf:"fixed64,18,opt,name=notional,proto3" json:"notional,omitempty"`
	FilledQty            float64  `protobuf:"fixed64,19,opt,name=filled_qty,json=filledQty,proto3" json:"filled_qty,omitempty"`
	OrderType            string   `protobuf:"bytes,20,opt,name=order_type,json=orderType,proto3" json:"order_type,omitempty"`
	Side                 string   `protobuf:"bytes,21,opt,name=side,proto3" json:"side,omitempty"`
	TimeInForce          string   `protobuf:"bytes,22,opt,name=time_in_force,json=timeInForce,proto3" json:"time_in_force,omitempty"`
	LimitPrice           float64  `protobuf:"fixed64,23,opt,name=limit_price,json=limitPrice,proto3" json:"limit_price,omitempty"`
	FilledAvgPrice       float64  `protobuf:"fixed64,24,opt,name=filled_avg_price,json=filledAvgPrice,proto3" json:"filled_avg_price,omitempty"`
	StopPrice            float64  `protobuf:"fixed64,25,opt,name=stop_price,json=stopPrice,proto3" json:"stop_price,omitempty"`
	TrailPrice           float64  `protobuf:"fixed64,26,opt,name=trail_price,json=trailPrice,proto3" json:"trail_price,omitempty"`
	TrailPercent         float64  `protobuf:"fixed64,27,opt,name=trail_percent,json=trailPercent,proto3" json:"trail_percent,omitempty"`
	Hwm                  float64  `protobuf:"fixed64,28,opt,name=hwm,proto3" json:"hwm,omitempty"`
	Status               string   `protobuf:"bytes,29,opt,name=status,proto3" json:"status,omitempty"`
	ExtendedHours        bool     `protobuf:"varint,30,opt,name=extended_hours,json=extendedHours,proto3" json:"extended_hours,omitempty"`
	Legs                 []*Order `protobuf:"bytes,31,rep,name=legs,proto3" json:"legs,omitempty"`
	UserId               string   `protobuf:"bytes,32,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_f65b0626cc3aada8, []int{6}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetClientOrderId() string {
	if m != nil {
		return m.ClientOrderId
	}
	return ""
}

func (m *Order) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Order) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Order) GetSubmittedAt() string {
	if m != nil {
		return m.SubmittedAt
	}
	return ""
}

func (m *Order) GetFilledAt() string {
	if m != nil {
		return m.FilledAt
	}
	return ""
}

func (m *Order) GetExpiredAt() string {
	if m != nil {
		return m.ExpiredAt
	}
	return ""
}

func (m *Order) GetCanceledAt() string {
	if m != nil {
		return m.CanceledAt
	}
	return ""
}

func (m *Order) GetFailedAt() string {
	if m != nil {
		return m.FailedAt
	}
	return ""
}

func (m *Order) GetReplacedAt() string {
	if m != nil {
		return m.ReplacedAt
	}
	return ""
}

func (m *Order) GetReplaces() string {
	if m != nil {
		return m.Replaces
	}
	return ""
}

func (m *Order) GetReplacedBy() string {
	if m != nil {
		return m.ReplacedBy
	}
	return ""
}

func (m *Order) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

func (m *Order) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Order) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *Order) GetAssetClass() string {
	if m != nil {
		return m.AssetClass
	}
	return ""
}

func (m *Order) GetQty() float64 {
	if m != nil {
		return m.Qty
	}
	return 0
}

func (m *Order) GetNotional() float64 {
	if m != nil {
		return m.Notional
	}
	return 0
}

func (m *Order) GetFilledQty() float64 {
	if m != nil {
		return m.FilledQty
	}
	return 0
}

func (m *Order) GetOrderType() string {
	if m != nil {
		return m.OrderType
	}
	return ""
}

func (m *Order) GetSide() string {
	if m != nil {
		return m.Side
	}
	return ""
}

func (m *Order) GetTimeInForce() string {
	if m != nil {
		return m.TimeInForce
	}
	return ""
}

func (m *Order) GetLimitPrice() float64 {
	if m != nil {
		return m.LimitPrice
	}
	return 0
}

func (m *Order) GetFilledAvgPrice() float64 {
	if m != nil {
		return m.FilledAvgPrice
	}
	return 0
}

func (m *Order) GetStopPrice() float64 {
	if m != nil {
		return m.StopPrice
	}
	return 0
}

func (m *Order) GetTrailPrice() float64 {
	if m != nil {
		return m.TrailPrice
	}
	return 0
}

func (m *Order) GetTrailPercent() float64 {
	if m != nil {
		return m.TrailPercent
	}
	return 0
}

func (m *Order) GetHwm() float64 {
	if m != nil {
		return m.Hwm
	}
	return 0
}

func (m *Order) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Order) GetExtendedHours() bool {
	if m != nil {
		return m.ExtendedHours
	}
	return false
}

func (m *Order) GetLegs() []*Order {
	if m != nil {
		return m.Legs
	}
	return nil
}

func (m *Order) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterEnum("order.Side", Side_name, Side_value)
	proto.RegisterEnum("order.OrderClass", OrderClass_name, OrderClass_value)
	proto.RegisterEnum("order.TimeInForce", TimeInForce_name, TimeInForce_value)
	proto.RegisterEnum("order.OrderType", OrderType_name, OrderType_value)
	proto.RegisterType((*PlaceOrderRes)(nil), "order.PlaceOrderRes")
	proto.RegisterType((*PlaceOrderReq)(nil), "order.PlaceOrderReq")
	proto.RegisterType((*TakeProfit)(nil), "order.TakeProfit")
	proto.RegisterType((*StopLoss)(nil), "order.StopLoss")
	proto.RegisterType((*ListOrdersReq)(nil), "order.ListOrdersReq")
	proto.RegisterType((*ListOrdersRes)(nil), "order.ListOrdersRes")
	proto.RegisterType((*Order)(nil), "order.Order")
}

func init() {
	proto.RegisterFile("proto/order.proto", fileDescriptor_f65b0626cc3aada8)
}

var fileDescriptor_f65b0626cc3aada8 = []byte{
	// 1024 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0x5f, 0x6f, 0xe3, 0x44,
	0x10, 0x3f, 0xe7, 0xaf, 0x3d, 0xf9, 0xd3, 0xed, 0xde, 0x5d, 0xeb, 0xf6, 0x28, 0x0d, 0x41, 0xa0,
	0xa8, 0xc0, 0x1d, 0xea, 0x49, 0x08, 0xf1, 0xd6, 0x43, 0x42, 0x14, 0x9d, 0x44, 0x71, 0xfb, 0xc4,
	0x8b, 0xe5, 0xd8, 0xdb, 0x74, 0x55, 0xc7, 0xf6, 0x79, 0x37, 0xbd, 0xe6, 0x9d, 0x0f, 0xc7, 0x77,
	0xe1, 0x4b, 0xa0, 0x9d, 0x59, 0x27, 0x4e, 0x1a, 0x1d, 0xbc, 0xcd, 0xfc, 0x7e, 0x33, 0x93, 0xdd,
	0x99, 0x9d, 0x5f, 0x0c, 0xfb, 0x45, 0x99, 0xeb, 0xfc, 0x4d, 0x5e, 0x26, 0xa2, 0x7c, 0x8d, 0x36,
	0x6f, 0xa3, 0x33, 0x7e, 0x0b, 0x83, 0xab, 0x34, 0x8a, 0xc5, 0xef, 0xc6, 0x0b, 0x84, 0xe2, 0x63,
	0x20, 0xc6, 0x77, 0x46, 0xce, 0xa4, 0x77, 0xde, 0x7f, 0x4d, 0x49, 0xc4, 0xdb, 0xa4, 0x7f, 0x5a,
	0x9b, 0x59, 0x1f, 0xf8, 0x29, 0xf4, 0xc4, 0x63, 0x7c, 0x17, 0x65, 0x33, 0x11, 0xca, 0x04, 0x73,
	0xbd, 0x00, 0x2a, 0xe8, 0x32, 0xe1, 0x07, 0xd0, 0x51, 0xcb, 0xf9, 0x34, 0x4f, 0xfd, 0x06, 0x72,
	0xd6, 0xe3, 0x0c, 0x9a, 0x1f, 0xf4, 0xd2, 0x6f, 0x8e, 0x9c, 0x89, 0x13, 0x18, 0x93, 0x1f, 0x83,
	0x9b, 0xe5, 0x5a, 0xe6, 0x59, 0x94, 0xfa, 0x2d, 0x84, 0x57, 0x3e, 0x3f, 0x85, 0x96, 0x92, 0x89,
	0xf0, 0xdb, 0x23, 0x67, 0x32, 0x3c, 0xef, 0xd9, 0xb3, 0x5d, 0xcb, 0x44, 0x04, 0x48, 0xf0, 0x37,
	0x00, 0x88, 0x85, 0x7a, 0x59, 0x08, 0xbf, 0x83, 0x61, 0xac, 0x7e, 0x85, 0x9b, 0x65, 0x21, 0x02,
	0x2f, 0xaf, 0x4c, 0xfe, 0x03, 0x0c, 0xb4, 0x9c, 0x8b, 0x50, 0x66, 0xe1, 0x6d, 0x5e, 0xc6, 0xc2,
	0xef, 0x62, 0x0e, 0xb7, 0x39, 0x37, 0x72, 0x2e, 0x2e, 0xb3, 0x5f, 0x0c, 0x13, 0xf4, 0xf4, 0xda,
	0x31, 0x17, 0x4e, 0xe5, 0x5c, 0xea, 0xb0, 0x28, 0x65, 0x2c, 0x7c, 0x17, 0x0f, 0x0a, 0x08, 0x5d,
	0x19, 0x84, 0x7f, 0x05, 0x43, 0xf1, 0xa8, 0x45, 0x96, 0x88, 0x24, 0xbc, 0xcb, 0x17, 0xa5, 0xf2,
	0xbd, 0x91, 0x33, 0x71, 0x83, 0x41, 0x85, 0xfe, 0x6a, 0x40, 0x7e, 0x02, 0xa0, 0x74, 0x5e, 0xd8,
	0x32, 0x80, 0x65, 0x3c, 0x83, 0x50, 0x95, 0xaf, 0x61, 0x2f, 0x4e, 0xa5, 0xc8, 0x74, 0x48, 0xd7,
	0x92, 0x89, 0xdf, 0xc3, 0xfe, 0x0d, 0x08, 0xc6, 0x3b, 0x5d, 0x26, 0xfc, 0x1c, 0x7a, 0x14, 0x10,
	0xa7, 0x91, 0x52, 0x7e, 0x1f, 0x2f, 0xb1, 0x5f, 0xbf, 0xf8, 0xcf, 0x86, 0x08, 0xa8, 0x3b, 0x68,
	0x9b, 0x1c, 0x1d, 0xdd, 0x8b, 0xb0, 0x28, 0xf3, 0x5b, 0xa9, 0xfd, 0x01, 0xce, 0xbb, 0xca, 0xb9,
	0x89, 0xee, 0xc5, 0x15, 0x12, 0x01, 0xe8, 0x95, 0xcd, 0xbf, 0x05, 0x3c, 0x5c, 0x98, 0xe6, 0x4a,
	0xf9, 0x43, 0xcc, 0xd8, 0xab, 0xa6, 0xa0, 0xf3, 0xe2, 0x7d, 0xae, 0x54, 0xe0, 0x2a, 0x6b, 0x99,
	0x26, 0xe9, 0x32, 0x92, 0xa9, 0xbd, 0xdd, 0x1e, 0x35, 0x09, 0x21, 0xba, 0xde, 0x97, 0x30, 0xb0,
	0x01, 0xa2, 0x8c, 0x45, 0xa6, 0x7d, 0x86, 0x21, 0x7d, 0x0a, 0x21, 0x6c, 0xfc, 0x1d, 0xc0, 0xfa,
	0x34, 0xdb, 0x8d, 0x77, 0xb6, 0x1b, 0x3f, 0xfe, 0x0d, 0xdc, 0xeb, 0xda, 0x01, 0x3e, 0x19, 0xbc,
	0xd5, 0xfe, 0xc6, 0x56, 0xfb, 0xc7, 0x13, 0x18, 0xbc, 0x97, 0x8a, 0xba, 0xac, 0xcc, 0x3b, 0x3f,
	0x84, 0xee, 0x42, 0xd1, 0x1c, 0xe8, 0x8d, 0x77, 0x8c, 0x7b, 0x99, 0x98, 0x3d, 0xaa, 0x47, 0xfe,
	0xbf, 0x3d, 0xfa, 0xbb, 0x0b, 0x6d, 0x04, 0xf8, 0x10, 0x1a, 0xab, 0x92, 0x0d, 0x99, 0xec, 0x9a,
	0x7b, 0x63, 0xd7, 0xdc, 0x4f, 0x00, 0xe2, 0x52, 0x44, 0x5a, 0x24, 0x61, 0xa4, 0x71, 0x8b, 0xbc,
	0xc0, 0xb3, 0xc8, 0x85, 0x36, 0xf4, 0xa2, 0x48, 0x2a, 0xba, 0x45, 0xb4, 0x45, 0x2e, 0x34, 0xff,
	0x02, 0xfa, 0x6a, 0x31, 0x9d, 0x4b, 0x6d, 0x03, 0xda, 0x18, 0xd0, 0x5b, 0x61, 0x17, 0x9a, 0xbf,
	0x02, 0xef, 0x56, 0xa6, 0x29, 0xf1, 0x1d, 0xe4, 0x5d, 0x02, 0xa8, 0xbc, 0x78, 0x2c, 0x64, 0x49,
	0x6c, 0x97, 0xca, 0x5b, 0xe4, 0x02, 0x47, 0x15, 0x47, 0x59, 0x2c, 0x6c, 0xb6, 0x4b, 0xa2, 0x50,
	0x41, 0xb6, 0x78, 0x24, 0x2d, 0xed, 0xd9, 0xe2, 0x08, 0x50, 0x76, 0x29, 0x0a, 0xa3, 0x32, 0x48,
	0x03, 0x65, 0x57, 0xd0, 0x85, 0x36, 0x42, 0x61, 0x3d, 0x65, 0x97, 0x62, 0xe5, 0x6f, 0x24, 0x4f,
	0x97, 0xb8, 0x0f, 0xb5, 0xe4, 0x77, 0x4b, 0x7e, 0x04, 0x6e, 0xa4, 0x94, 0xd0, 0xa6, 0xb3, 0x03,
	0x64, 0xbb, 0xe8, 0x6f, 0x48, 0xd5, 0x70, 0x43, 0xaa, 0x8e, 0xc1, 0xad, 0x04, 0x0d, 0x9f, 0xb2,
	0x17, 0xac, 0x7c, 0xf3, 0x7b, 0x54, 0x8e, 0xf6, 0x8f, 0xd1, 0xef, 0x21, 0x44, 0xcb, 0x66, 0x75,
	0x6e, 0x7f, 0xb7, 0xce, 0xf1, 0x2d, 0x9d, 0x3b, 0x01, 0xb0, 0x5d, 0x37, 0x49, 0xcf, 0xe9, 0x59,
	0x12, 0xf2, 0x87, 0x5e, 0x1a, 0xba, 0xa6, 0x72, 0x2f, 0xa8, 0xef, 0x6b, 0x4d, 0xe3, 0x56, 0x25,
	0x5f, 0x22, 0x41, 0xc2, 0x38, 0xde, 0xd6, 0xb9, 0x03, 0x9a, 0xf5, 0x27, 0x34, 0xed, 0xf0, 0xc9,
	0xb6, 0x4c, 0x80, 0x55, 0x8f, 0xe1, 0x61, 0x66, 0xa3, 0x7c, 0x8c, 0x1a, 0xda, 0x37, 0xf1, 0x30,
	0xdb, 0xb5, 0x57, 0x47, 0xdb, 0xb2, 0xb6, 0x25, 0x0c, 0xc7, 0xff, 0x2d, 0x0c, 0xaf, 0x9e, 0x0a,
	0x83, 0xe9, 0xe9, 0xdd, 0xc7, 0xb9, 0xff, 0x19, 0xf5, 0xf4, 0xee, 0xe3, 0x1c, 0x47, 0xa7, 0x23,
	0xbd, 0x50, 0xfe, 0x89, 0x1d, 0x1d, 0x7a, 0x3b, 0xc4, 0xf8, 0xf3, 0x5d, 0x62, 0x3c, 0x82, 0x56,
	0x2a, 0x66, 0xca, 0x3f, 0x1d, 0x35, 0x9f, 0xac, 0x2c, 0x32, 0xf5, 0xfd, 0x1f, 0xd5, 0xf7, 0xff,
	0xec, 0x1b, 0x68, 0x99, 0xbf, 0x21, 0xce, 0xa0, 0x2f, 0xb3, 0x87, 0x28, 0x95, 0x49, 0x68, 0xfa,
	0xce, 0x9e, 0xf1, 0x2e, 0x34, 0xa7, 0x8b, 0x25, 0x73, 0xb8, 0x0b, 0x2d, 0x25, 0xd2, 0x94, 0x35,
	0xce, 0xae, 0x00, 0xd6, 0x9a, 0xcc, 0x0f, 0xe1, 0x79, 0x95, 0x52, 0xd3, 0x70, 0xf6, 0x8c, 0xf7,
	0xa0, 0x3b, 0x2d, 0xa3, 0xf8, 0x5e, 0x68, 0xe6, 0x98, 0x32, 0xb9, 0xce, 0x59, 0x03, 0x8d, 0x38,
	0x67, 0x4d, 0x0e, 0xd0, 0x51, 0x72, 0x5e, 0xa4, 0x82, 0xb5, 0xce, 0x0a, 0xe8, 0xd5, 0xfe, 0xaa,
	0xf8, 0x11, 0xbc, 0xac, 0x4a, 0x6e, 0x4c, 0x9d, 0x8e, 0x93, 0x44, 0x4b, 0x2a, 0x38, 0xd3, 0xb1,
	0x2d, 0x58, 0xcc, 0x58, 0xd3, 0x18, 0x32, 0x8f, 0x59, 0xcb, 0x18, 0xb7, 0xf9, 0x3d, 0x6b, 0x53,
	0xcc, 0x23, 0xeb, 0x90, 0x91, 0xb0, 0xae, 0x31, 0xe2, 0x54, 0x31, 0xf7, 0x6c, 0x06, 0xde, 0xea,
	0x0f, 0x95, 0x1f, 0x00, 0xdf, 0xbc, 0x82, 0x79, 0x98, 0xec, 0x99, 0x39, 0xe2, 0x3c, 0x2a, 0xe9,
	0x02, 0x1e, 0xb4, 0xf1, 0x29, 0xb1, 0x06, 0x76, 0x42, 0xe7, 0x05, 0x6b, 0xf2, 0xa1, 0x7d, 0x27,
	0xc4, 0xb4, 0xf8, 0xbe, 0x9d, 0xbb, 0xcc, 0x66, 0x21, 0x86, 0xb4, 0xcf, 0xff, 0x72, 0xa0, 0x8f,
	0xbf, 0x74, 0x2d, 0xca, 0x07, 0xf3, 0x36, 0x7e, 0x02, 0x58, 0x4b, 0x2d, 0x7f, 0x61, 0xa7, 0xb4,
	0xa1, 0xd3, 0xc7, 0xbb, 0x50, 0xf5, 0xbd, 0xc3, 0x7f, 0x04, 0x58, 0x7f, 0xb8, 0xac, 0x72, 0x37,
	0xbe, 0x65, 0x8e, 0x77, 0xa1, 0xea, 0x9d, 0xf7, 0x67, 0x17, 0xe1, 0x62, 0x3a, 0xed, 0xe0, 0x17,
	0xd4, 0xdb, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x93, 0x96, 0x24, 0x37, 0x56, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderServiceClient interface {
	ListOrders(ctx context.Context, in *ListOrdersReq, opts ...grpc.CallOption) (OrderService_ListOrdersClient, error)
	PlaceOrder(ctx context.Context, in *PlaceOrderReq, opts ...grpc.CallOption) (*PlaceOrderRes, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) ListOrders(ctx context.Context, in *ListOrdersReq, opts ...grpc.CallOption) (OrderService_ListOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OrderService_serviceDesc.Streams[0], "/order.OrderService/ListOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderServiceListOrdersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderService_ListOrdersClient interface {
	Recv() (*ListOrdersRes, error)
	grpc.ClientStream
}

type orderServiceListOrdersClient struct {
	grpc.ClientStream
}

func (x *orderServiceListOrdersClient) Recv() (*ListOrdersRes, error) {
	m := new(ListOrdersRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderServiceClient) PlaceOrder(ctx context.Context, in *PlaceOrderReq, opts ...grpc.CallOption) (*PlaceOrderRes, error) {
	out := new(PlaceOrderRes)
	err := c.cc.Invoke(ctx, "/order.OrderService/PlaceOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
type OrderServiceServer interface {
	ListOrders(*ListOrdersReq, OrderService_ListOrdersServer) error
	PlaceOrder(context.Context, *PlaceOrderReq) (*PlaceOrderRes, error)
}

// UnimplementedOrderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (*UnimplementedOrderServiceServer) ListOrders(req *ListOrdersReq, srv OrderService_ListOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method ListOrders not implemented")
}
func (*UnimplementedOrderServiceServer) PlaceOrder(ctx context.Context, req *PlaceOrderReq) (*PlaceOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&_OrderService_serviceDesc, srv)
}

func _OrderService_ListOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListOrdersReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderServiceServer).ListOrders(m, &orderServiceListOrdersServer{stream})
}

type OrderService_ListOrdersServer interface {
	Send(*ListOrdersRes) error
	grpc.ServerStream
}

type orderServiceListOrdersServer struct {
	grpc.ServerStream
}

func (x *orderServiceListOrdersServer) Send(m *ListOrdersRes) error {
	return x.ServerStream.SendMsg(m)
}

func _OrderService_PlaceOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).PlaceOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/PlaceOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).PlaceOrder(ctx, req.(*PlaceOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceOrder",
			Handler:    _OrderService_PlaceOrder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListOrders",
			Handler:       _OrderService_ListOrders_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/order.proto",
}
