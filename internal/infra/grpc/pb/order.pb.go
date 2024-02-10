// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protofiles/order.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_da0bb636fa0e5ac5, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type CreateOrderRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Price                float32  `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Tax                  float32  `protobuf:"fixed32,3,opt,name=tax,proto3" json:"tax,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrderRequest) Reset()         { *m = CreateOrderRequest{} }
func (m *CreateOrderRequest) String() string { return proto.CompactTextString(m) }
func (*CreateOrderRequest) ProtoMessage()    {}
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da0bb636fa0e5ac5, []int{1}
}

func (m *CreateOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderRequest.Unmarshal(m, b)
}
func (m *CreateOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderRequest.Marshal(b, m, deterministic)
}
func (m *CreateOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderRequest.Merge(m, src)
}
func (m *CreateOrderRequest) XXX_Size() int {
	return xxx_messageInfo_CreateOrderRequest.Size(m)
}
func (m *CreateOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderRequest proto.InternalMessageInfo

func (m *CreateOrderRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateOrderRequest) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *CreateOrderRequest) GetTax() float32 {
	if m != nil {
		return m.Tax
	}
	return 0
}

type CreateOrderResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Price                float32  `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Tax                  float32  `protobuf:"fixed32,3,opt,name=tax,proto3" json:"tax,omitempty"`
	FinalPrice           float32  `protobuf:"fixed32,4,opt,name=final_price,json=finalPrice,proto3" json:"final_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrderResponse) Reset()         { *m = CreateOrderResponse{} }
func (m *CreateOrderResponse) String() string { return proto.CompactTextString(m) }
func (*CreateOrderResponse) ProtoMessage()    {}
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_da0bb636fa0e5ac5, []int{2}
}

func (m *CreateOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderResponse.Unmarshal(m, b)
}
func (m *CreateOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderResponse.Marshal(b, m, deterministic)
}
func (m *CreateOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderResponse.Merge(m, src)
}
func (m *CreateOrderResponse) XXX_Size() int {
	return xxx_messageInfo_CreateOrderResponse.Size(m)
}
func (m *CreateOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderResponse proto.InternalMessageInfo

func (m *CreateOrderResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateOrderResponse) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *CreateOrderResponse) GetTax() float32 {
	if m != nil {
		return m.Tax
	}
	return 0
}

func (m *CreateOrderResponse) GetFinalPrice() float32 {
	if m != nil {
		return m.FinalPrice
	}
	return 0
}

type Order struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Price                float32  `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Tax                  float32  `protobuf:"fixed32,3,opt,name=tax,proto3" json:"tax,omitempty"`
	FinalPrice           float32  `protobuf:"fixed32,4,opt,name=final_price,json=finalPrice,proto3" json:"final_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_da0bb636fa0e5ac5, []int{3}
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

func (m *Order) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Order) GetTax() float32 {
	if m != nil {
		return m.Tax
	}
	return 0
}

func (m *Order) GetFinalPrice() float32 {
	if m != nil {
		return m.FinalPrice
	}
	return 0
}

type FindAllOrdersResponse struct {
	Orders               []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindAllOrdersResponse) Reset()         { *m = FindAllOrdersResponse{} }
func (m *FindAllOrdersResponse) String() string { return proto.CompactTextString(m) }
func (*FindAllOrdersResponse) ProtoMessage()    {}
func (*FindAllOrdersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_da0bb636fa0e5ac5, []int{4}
}

func (m *FindAllOrdersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindAllOrdersResponse.Unmarshal(m, b)
}
func (m *FindAllOrdersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindAllOrdersResponse.Marshal(b, m, deterministic)
}
func (m *FindAllOrdersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindAllOrdersResponse.Merge(m, src)
}
func (m *FindAllOrdersResponse) XXX_Size() int {
	return xxx_messageInfo_FindAllOrdersResponse.Size(m)
}
func (m *FindAllOrdersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindAllOrdersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindAllOrdersResponse proto.InternalMessageInfo

func (m *FindAllOrdersResponse) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "pb.empty")
	proto.RegisterType((*CreateOrderRequest)(nil), "pb.CreateOrderRequest")
	proto.RegisterType((*CreateOrderResponse)(nil), "pb.CreateOrderResponse")
	proto.RegisterType((*Order)(nil), "pb.Order")
	proto.RegisterType((*FindAllOrdersResponse)(nil), "pb.FindAllOrdersResponse")
}

func init() { proto.RegisterFile("protofiles/order.proto", fileDescriptor_da0bb636fa0e5ac5) }

var fileDescriptor_da0bb636fa0e5ac5 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x69, 0xd7, 0xae, 0xec, 0x74, 0x11, 0x89, 0x5a, 0xeb, 0x5e, 0xac, 0x3d, 0xf5, 0xd4,
	0xc2, 0xea, 0xc9, 0x83, 0xa0, 0x82, 0x27, 0x41, 0xa9, 0x37, 0x2f, 0x9a, 0x6e, 0xa7, 0x12, 0x88,
	0x69, 0x4c, 0xa2, 0xe8, 0xdd, 0x1f, 0x2e, 0x99, 0x2e, 0xb2, 0xb2, 0x9e, 0x04, 0x6f, 0x93, 0x37,
	0x93, 0xf9, 0x66, 0x5e, 0x02, 0x89, 0x36, 0xbd, 0xeb, 0x3b, 0x21, 0xd1, 0x56, 0xbd, 0x69, 0xd1,
	0x94, 0x24, 0xb0, 0x50, 0x37, 0xf9, 0x26, 0x44, 0xf8, 0xac, 0xdd, 0x47, 0x7e, 0x0d, 0xec, 0xd2,
	0x20, 0x77, 0x78, 0xe3, 0x2b, 0x6a, 0x7c, 0x79, 0x45, 0xeb, 0xd8, 0x16, 0x84, 0xa2, 0x4d, 0x83,
	0x2c, 0x28, 0x26, 0x75, 0x28, 0x5a, 0xb6, 0x0b, 0x91, 0x36, 0x62, 0x81, 0x69, 0x98, 0x05, 0x45,
	0x58, 0x0f, 0x07, 0xb6, 0x0d, 0x23, 0xc7, 0xdf, 0xd3, 0x11, 0x69, 0x3e, 0xcc, 0x15, 0xec, 0xfc,
	0xe8, 0x66, 0x75, 0xaf, 0x2c, 0xfe, 0xb5, 0x1d, 0x3b, 0x84, 0xb8, 0x13, 0x8a, 0xcb, 0x87, 0xa1,
	0x7a, 0x83, 0x32, 0x40, 0xd2, 0xad, 0x57, 0xf2, 0x47, 0x88, 0x88, 0xf4, 0x7f, 0x84, 0x53, 0xd8,
	0xbb, 0x12, 0xaa, 0x3d, 0x97, 0x92, 0x40, 0xf6, 0x7b, 0xa7, 0x23, 0x18, 0x93, 0xa9, 0x36, 0x0d,
	0xb2, 0x51, 0x11, 0xcf, 0x27, 0xa5, 0x6e, 0xca, 0x61, 0xed, 0x65, 0x62, 0xfe, 0x19, 0xc0, 0x94,
	0x94, 0x3b, 0x34, 0x6f, 0x9e, 0x7f, 0x06, 0xf1, 0x8a, 0x3d, 0x2c, 0xf1, 0x57, 0xd6, 0xdd, 0x9f,
	0xed, 0xaf, 0xe9, 0x4b, 0xe6, 0x09, 0x4c, 0x57, 0x87, 0x61, 0xc4, 0xa4, 0x77, 0x9c, 0x1d, 0xf8,
	0xf0, 0xd7, 0x49, 0x2f, 0xd2, 0xfb, 0x44, 0x28, 0x87, 0x46, 0x71, 0x59, 0x09, 0xd5, 0x19, 0x5e,
	0x3d, 0x19, 0xbd, 0xa8, 0x74, 0xd3, 0x8c, 0xe9, 0x43, 0x1c, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x84, 0x9a, 0xd0, 0xed, 0x2a, 0x02, 0x00, 0x00,
}
