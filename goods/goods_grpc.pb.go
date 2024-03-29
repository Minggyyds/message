// 声明 proto 语法版本，固定值

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: goods/goods.proto

// proto 包名

package goods

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GoodsService_GetGoods_FullMethodName     = "/Goods.GoodsService/GetGoods"
	GoodsService_GoodsUpdate_FullMethodName  = "/Goods.GoodsService/GoodsUpdate"
	GoodsService_GoodsCreate_FullMethodName  = "/Goods.GoodsService/GoodsCreate"
	GoodsService_GetByIdGoods_FullMethodName = "/Goods.GoodsService/GetByIdGoods"
)

// GoodsServiceClient is the client API for GoodsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodsServiceClient interface {
	// 定义一个 GoodsList 商品列表展示 一元 rpc 方法，请求体和响应体必填。
	GetGoods(ctx context.Context, in *GetGoodsReq, opts ...grpc.CallOption) (*GetGoodsRes, error)
	// 定义一个 GoodsUpdate 修改商品 一元 rpc 方法，请求体和响应体必填。
	GoodsUpdate(ctx context.Context, in *GoodsUpdateReq, opts ...grpc.CallOption) (*GoodsUpdateResp, error)
	// 定义一个 GoodsCreate 创建商品 一元 rpc 方法，请求体和响应体必填。
	GoodsCreate(ctx context.Context, in *GoodsCreateReq, opts ...grpc.CallOption) (*GoodsCreateResp, error)
	GetByIdGoods(ctx context.Context, in *GetByIdGoodsReq, opts ...grpc.CallOption) (*GetByIdGoodsRes, error)
}

type goodsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodsServiceClient(cc grpc.ClientConnInterface) GoodsServiceClient {
	return &goodsServiceClient{cc}
}

func (c *goodsServiceClient) GetGoods(ctx context.Context, in *GetGoodsReq, opts ...grpc.CallOption) (*GetGoodsRes, error) {
	out := new(GetGoodsRes)
	err := c.cc.Invoke(ctx, GoodsService_GetGoods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServiceClient) GoodsUpdate(ctx context.Context, in *GoodsUpdateReq, opts ...grpc.CallOption) (*GoodsUpdateResp, error) {
	out := new(GoodsUpdateResp)
	err := c.cc.Invoke(ctx, GoodsService_GoodsUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServiceClient) GoodsCreate(ctx context.Context, in *GoodsCreateReq, opts ...grpc.CallOption) (*GoodsCreateResp, error) {
	out := new(GoodsCreateResp)
	err := c.cc.Invoke(ctx, GoodsService_GoodsCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServiceClient) GetByIdGoods(ctx context.Context, in *GetByIdGoodsReq, opts ...grpc.CallOption) (*GetByIdGoodsRes, error) {
	out := new(GetByIdGoodsRes)
	err := c.cc.Invoke(ctx, GoodsService_GetByIdGoods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsServiceServer is the server API for GoodsService service.
// All implementations must embed UnimplementedGoodsServiceServer
// for forward compatibility
type GoodsServiceServer interface {
	// 定义一个 GoodsList 商品列表展示 一元 rpc 方法，请求体和响应体必填。
	GetGoods(context.Context, *GetGoodsReq) (*GetGoodsRes, error)
	// 定义一个 GoodsUpdate 修改商品 一元 rpc 方法，请求体和响应体必填。
	GoodsUpdate(context.Context, *GoodsUpdateReq) (*GoodsUpdateResp, error)
	// 定义一个 GoodsCreate 创建商品 一元 rpc 方法，请求体和响应体必填。
	GoodsCreate(context.Context, *GoodsCreateReq) (*GoodsCreateResp, error)
	GetByIdGoods(context.Context, *GetByIdGoodsReq) (*GetByIdGoodsRes, error)
	mustEmbedUnimplementedGoodsServiceServer()
}

// UnimplementedGoodsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGoodsServiceServer struct {
}

func (UnimplementedGoodsServiceServer) GetGoods(context.Context, *GetGoodsReq) (*GetGoodsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoods not implemented")
}
func (UnimplementedGoodsServiceServer) GoodsUpdate(context.Context, *GoodsUpdateReq) (*GoodsUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GoodsUpdate not implemented")
}
func (UnimplementedGoodsServiceServer) GoodsCreate(context.Context, *GoodsCreateReq) (*GoodsCreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GoodsCreate not implemented")
}
func (UnimplementedGoodsServiceServer) GetByIdGoods(context.Context, *GetByIdGoodsReq) (*GetByIdGoodsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdGoods not implemented")
}
func (UnimplementedGoodsServiceServer) mustEmbedUnimplementedGoodsServiceServer() {}

// UnsafeGoodsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodsServiceServer will
// result in compilation errors.
type UnsafeGoodsServiceServer interface {
	mustEmbedUnimplementedGoodsServiceServer()
}

func RegisterGoodsServiceServer(s grpc.ServiceRegistrar, srv GoodsServiceServer) {
	s.RegisterService(&GoodsService_ServiceDesc, srv)
}

func _GoodsService_GetGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).GetGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoodsService_GetGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).GetGoods(ctx, req.(*GetGoodsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_GoodsUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).GoodsUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoodsService_GoodsUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).GoodsUpdate(ctx, req.(*GoodsUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_GoodsCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).GoodsCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoodsService_GoodsCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).GoodsCreate(ctx, req.(*GoodsCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_GetByIdGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdGoodsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).GetByIdGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoodsService_GetByIdGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).GetByIdGoods(ctx, req.(*GetByIdGoodsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GoodsService_ServiceDesc is the grpc.ServiceDesc for GoodsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoodsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Goods.GoodsService",
	HandlerType: (*GoodsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGoods",
			Handler:    _GoodsService_GetGoods_Handler,
		},
		{
			MethodName: "GoodsUpdate",
			Handler:    _GoodsService_GoodsUpdate_Handler,
		},
		{
			MethodName: "GoodsCreate",
			Handler:    _GoodsService_GoodsCreate_Handler,
		},
		{
			MethodName: "GetByIdGoods",
			Handler:    _GoodsService_GetByIdGoods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goods/goods.proto",
}
