// 声明 proto 语法版本，固定值

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: goods/goods.proto

// proto 包名

package proto

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

type Types int32

const (
	Types_Default Types = 0
	Types_DianZi  Types = 20
	Types_LingShi Types = 30
)

// Enum value maps for Types.
var (
	Types_name = map[int32]string{
		0:  "Default",
		20: "DianZi",
		30: "LingShi",
	}
	Types_value = map[string]int32{
		"Default": 0,
		"DianZi":  20,
		"LingShi": 30,
	}
)

func (x Types) Enum() *Types {
	p := new(Types)
	*p = x
	return p
}

func (x Types) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Types) Descriptor() protoreflect.EnumDescriptor {
	return file_goods_goods_proto_enumTypes[0].Descriptor()
}

func (Types) Type() protoreflect.EnumType {
	return &file_goods_goods_proto_enumTypes[0]
}

func (x Types) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Types.Descriptor instead.
func (Types) EnumDescriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{0}
}

type Goods struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64   `protobuf:"varint,10,opt,name=id,proto3" json:"id,omitempty"`
	GoodsName  string  `protobuf:"bytes,20,opt,name=GoodsName,proto3" json:"GoodsName,omitempty"`
	GoodsPrice float32 `protobuf:"fixed32,30,opt,name=GoodsPrice,proto3" json:"GoodsPrice,omitempty"`
	GoodsNum   int64   `protobuf:"varint,40,opt,name=GoodsNum,proto3" json:"GoodsNum,omitempty"`
	GoodsType  Types   `protobuf:"varint,50,opt,name=GoodsType,proto3,enum=Goods.Types" json:"GoodsType,omitempty"`
}

func (x *Goods) Reset() {
	*x = Goods{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Goods) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Goods) ProtoMessage() {}

func (x *Goods) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Goods.ProtoReflect.Descriptor instead.
func (*Goods) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{0}
}

func (x *Goods) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Goods) GetGoodsName() string {
	if x != nil {
		return x.GoodsName
	}
	return ""
}

func (x *Goods) GetGoodsPrice() float32 {
	if x != nil {
		return x.GoodsPrice
	}
	return 0
}

func (x *Goods) GetGoodsNum() int64 {
	if x != nil {
		return x.GoodsNum
	}
	return 0
}

func (x *Goods) GetGoodsType() Types {
	if x != nil {
		return x.GoodsType
	}
	return Types_Default
}

// 定义商品列表展示的请求体
type GoodsListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     string `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize string `protobuf:"bytes,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *GoodsListReq) Reset() {
	*x = GoodsListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsListReq) ProtoMessage() {}

func (x *GoodsListReq) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsListReq.ProtoReflect.Descriptor instead.
func (*GoodsListReq) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{1}
}

func (x *GoodsListReq) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

func (x *GoodsListReq) GetPageSize() string {
	if x != nil {
		return x.PageSize
	}
	return ""
}

// 定义商品列表展示的响应体
type GoodsListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Goods []*Goods `protobuf:"bytes,2,rep,name=goods,proto3" json:"goods,omitempty"`
}

func (x *GoodsListResp) Reset() {
	*x = GoodsListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsListResp) ProtoMessage() {}

func (x *GoodsListResp) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsListResp.ProtoReflect.Descriptor instead.
func (*GoodsListResp) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{2}
}

func (x *GoodsListResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GoodsListResp) GetGoods() []*Goods {
	if x != nil {
		return x.Goods
	}
	return nil
}

// 定义商品修改的请求体
type GoodsUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Goods *Goods `protobuf:"bytes,1,opt,name=goods,proto3" json:"goods,omitempty"`
}

func (x *GoodsUpdateReq) Reset() {
	*x = GoodsUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsUpdateReq) ProtoMessage() {}

func (x *GoodsUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsUpdateReq.ProtoReflect.Descriptor instead.
func (*GoodsUpdateReq) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{3}
}

func (x *GoodsUpdateReq) GetGoods() *Goods {
	if x != nil {
		return x.Goods
	}
	return nil
}

// 定义商品修改的响应体
type GoodsUpdateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Goods *Goods `protobuf:"bytes,1,opt,name=goods,proto3" json:"goods,omitempty"`
}

func (x *GoodsUpdateResp) Reset() {
	*x = GoodsUpdateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsUpdateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsUpdateResp) ProtoMessage() {}

func (x *GoodsUpdateResp) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsUpdateResp.ProtoReflect.Descriptor instead.
func (*GoodsUpdateResp) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{4}
}

func (x *GoodsUpdateResp) GetGoods() *Goods {
	if x != nil {
		return x.Goods
	}
	return nil
}

// 定义商品创建的请求体
type GoodsCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Goods *Goods `protobuf:"bytes,1,opt,name=goods,proto3" json:"goods,omitempty"`
}

func (x *GoodsCreateReq) Reset() {
	*x = GoodsCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsCreateReq) ProtoMessage() {}

func (x *GoodsCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsCreateReq.ProtoReflect.Descriptor instead.
func (*GoodsCreateReq) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{5}
}

func (x *GoodsCreateReq) GetGoods() *Goods {
	if x != nil {
		return x.Goods
	}
	return nil
}

// 定义商品创建的响应体
type GoodsCreateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Goods *Goods `protobuf:"bytes,1,opt,name=goods,proto3" json:"goods,omitempty"`
}

func (x *GoodsCreateResp) Reset() {
	*x = GoodsCreateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsCreateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsCreateResp) ProtoMessage() {}

func (x *GoodsCreateResp) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsCreateResp.ProtoReflect.Descriptor instead.
func (*GoodsCreateResp) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{6}
}

func (x *GoodsCreateResp) GetGoods() *Goods {
	if x != nil {
		return x.Goods
	}
	return nil
}

var File_goods_goods_proto protoreflect.FileDescriptor

var file_goods_goods_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x22, 0x9d, 0x01, 0x0a, 0x05, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x75, 0x6d, 0x18, 0x28,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x2a,
	0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52,
	0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3e, 0x0a, 0x0c, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x49, 0x0a, 0x0d, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x22, 0x0a, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x05,
	0x67, 0x6f, 0x6f, 0x64, 0x73, 0x22, 0x34, 0x0a, 0x0e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x52, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x22, 0x35, 0x0a, 0x0f, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22,
	0x0a, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x05, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x22, 0x34, 0x0a, 0x0e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x52, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x22, 0x35, 0x0a, 0x0f, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x05, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2a,
	0x2d, 0x0a, 0x05, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x69, 0x61, 0x6e, 0x5a, 0x69, 0x10,
	0x14, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x69, 0x6e, 0x67, 0x53, 0x68, 0x69, 0x10, 0x1e, 0x32, 0xc2,
	0x01, 0x0a, 0x0c, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x36, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x14, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3c, 0x0a, 0x0b, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3c, 0x0a, 0x0b, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goods_goods_proto_rawDescOnce sync.Once
	file_goods_goods_proto_rawDescData = file_goods_goods_proto_rawDesc
)

func file_goods_goods_proto_rawDescGZIP() []byte {
	file_goods_goods_proto_rawDescOnce.Do(func() {
		file_goods_goods_proto_rawDescData = protoimpl.X.CompressGZIP(file_goods_goods_proto_rawDescData)
	})
	return file_goods_goods_proto_rawDescData
}

var file_goods_goods_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_goods_goods_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_goods_goods_proto_goTypes = []interface{}{
	(Types)(0),              // 0: Goods.Types
	(*Goods)(nil),           // 1: Goods.Goods
	(*GoodsListReq)(nil),    // 2: Goods.GoodsListReq
	(*GoodsListResp)(nil),   // 3: Goods.GoodsListResp
	(*GoodsUpdateReq)(nil),  // 4: Goods.GoodsUpdateReq
	(*GoodsUpdateResp)(nil), // 5: Goods.GoodsUpdateResp
	(*GoodsCreateReq)(nil),  // 6: Goods.GoodsCreateReq
	(*GoodsCreateResp)(nil), // 7: Goods.GoodsCreateResp
}
var file_goods_goods_proto_depIdxs = []int32{
	0, // 0: Goods.Goods.GoodsType:type_name -> Goods.Types
	1, // 1: Goods.GoodsListResp.goods:type_name -> Goods.Goods
	1, // 2: Goods.GoodsUpdateReq.goods:type_name -> Goods.Goods
	1, // 3: Goods.GoodsUpdateResp.goods:type_name -> Goods.Goods
	1, // 4: Goods.GoodsCreateReq.goods:type_name -> Goods.Goods
	1, // 5: Goods.GoodsCreateResp.goods:type_name -> Goods.Goods
	2, // 6: Goods.GoodsService.GoodsList:input_type -> Goods.GoodsListReq
	4, // 7: Goods.GoodsService.GoodsUpdate:input_type -> Goods.GoodsUpdateReq
	6, // 8: Goods.GoodsService.GoodsCreate:input_type -> Goods.GoodsCreateReq
	3, // 9: Goods.GoodsService.GoodsList:output_type -> Goods.GoodsListResp
	5, // 10: Goods.GoodsService.GoodsUpdate:output_type -> Goods.GoodsUpdateResp
	7, // 11: Goods.GoodsService.GoodsCreate:output_type -> Goods.GoodsCreateResp
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_goods_goods_proto_init() }
func file_goods_goods_proto_init() {
	if File_goods_goods_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goods_goods_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Goods); i {
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
		file_goods_goods_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsListReq); i {
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
		file_goods_goods_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsListResp); i {
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
		file_goods_goods_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsUpdateReq); i {
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
		file_goods_goods_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsUpdateResp); i {
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
		file_goods_goods_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsCreateReq); i {
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
		file_goods_goods_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsCreateResp); i {
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
			RawDescriptor: file_goods_goods_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goods_goods_proto_goTypes,
		DependencyIndexes: file_goods_goods_proto_depIdxs,
		EnumInfos:         file_goods_goods_proto_enumTypes,
		MessageInfos:      file_goods_goods_proto_msgTypes,
	}.Build()
	File_goods_goods_proto = out.File
	file_goods_goods_proto_rawDesc = nil
	file_goods_goods_proto_goTypes = nil
	file_goods_goods_proto_depIdxs = nil
}