// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.1
// source: counter/v1/counter.proto

package counterv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CounterService_AddCounter_FullMethodName          = "/counter.v1.CounterService/AddCounter"
	CounterService_GetCounterLevels_FullMethodName    = "/counter.v1.CounterService/GetCounterLevels"
	CounterService_ChangeCounterLevels_FullMethodName = "/counter.v1.CounterService/ChangeCounterLevels"
	CounterService_ClearCounterLevels_FullMethodName  = "/counter.v1.CounterService/ClearCounterLevels"
)

// CounterServiceClient is the client API for CounterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CounterServiceClient interface {
	AddCounter(ctx context.Context, in *AddCounterReq, opts ...grpc.CallOption) (*AddCounterResp, error)
	GetCounterLevels(ctx context.Context, in *GetCounterLevelsReq, opts ...grpc.CallOption) (*GetCounterLevelsResp, error)
	ChangeCounterLevels(ctx context.Context, in *ChangeCounterLevelsReq, opts ...grpc.CallOption) (*ChangeCounterLevelsResp, error)
	ClearCounterLevels(ctx context.Context, in *ClearCounterLevelsReq, opts ...grpc.CallOption) (*ClearCounterLevelsResp, error)
}

type counterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCounterServiceClient(cc grpc.ClientConnInterface) CounterServiceClient {
	return &counterServiceClient{cc}
}

func (c *counterServiceClient) AddCounter(ctx context.Context, in *AddCounterReq, opts ...grpc.CallOption) (*AddCounterResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddCounterResp)
	err := c.cc.Invoke(ctx, CounterService_AddCounter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) GetCounterLevels(ctx context.Context, in *GetCounterLevelsReq, opts ...grpc.CallOption) (*GetCounterLevelsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCounterLevelsResp)
	err := c.cc.Invoke(ctx, CounterService_GetCounterLevels_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) ChangeCounterLevels(ctx context.Context, in *ChangeCounterLevelsReq, opts ...grpc.CallOption) (*ChangeCounterLevelsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChangeCounterLevelsResp)
	err := c.cc.Invoke(ctx, CounterService_ChangeCounterLevels_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) ClearCounterLevels(ctx context.Context, in *ClearCounterLevelsReq, opts ...grpc.CallOption) (*ClearCounterLevelsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClearCounterLevelsResp)
	err := c.cc.Invoke(ctx, CounterService_ClearCounterLevels_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CounterServiceServer is the server API for CounterService service.
// All implementations must embed UnimplementedCounterServiceServer
// for forward compatibility.
type CounterServiceServer interface {
	AddCounter(context.Context, *AddCounterReq) (*AddCounterResp, error)
	GetCounterLevels(context.Context, *GetCounterLevelsReq) (*GetCounterLevelsResp, error)
	ChangeCounterLevels(context.Context, *ChangeCounterLevelsReq) (*ChangeCounterLevelsResp, error)
	ClearCounterLevels(context.Context, *ClearCounterLevelsReq) (*ClearCounterLevelsResp, error)
	mustEmbedUnimplementedCounterServiceServer()
}

// UnimplementedCounterServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCounterServiceServer struct{}

func (UnimplementedCounterServiceServer) AddCounter(context.Context, *AddCounterReq) (*AddCounterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCounter not implemented")
}
func (UnimplementedCounterServiceServer) GetCounterLevels(context.Context, *GetCounterLevelsReq) (*GetCounterLevelsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCounterLevels not implemented")
}
func (UnimplementedCounterServiceServer) ChangeCounterLevels(context.Context, *ChangeCounterLevelsReq) (*ChangeCounterLevelsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeCounterLevels not implemented")
}
func (UnimplementedCounterServiceServer) ClearCounterLevels(context.Context, *ClearCounterLevelsReq) (*ClearCounterLevelsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearCounterLevels not implemented")
}
func (UnimplementedCounterServiceServer) mustEmbedUnimplementedCounterServiceServer() {}
func (UnimplementedCounterServiceServer) testEmbeddedByValue()                        {}

// UnsafeCounterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CounterServiceServer will
// result in compilation errors.
type UnsafeCounterServiceServer interface {
	mustEmbedUnimplementedCounterServiceServer()
}

func RegisterCounterServiceServer(s grpc.ServiceRegistrar, srv CounterServiceServer) {
	// If the following call pancis, it indicates UnimplementedCounterServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CounterService_ServiceDesc, srv)
}

func _CounterService_AddCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCounterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).AddCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CounterService_AddCounter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).AddCounter(ctx, req.(*AddCounterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_GetCounterLevels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCounterLevelsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).GetCounterLevels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CounterService_GetCounterLevels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).GetCounterLevels(ctx, req.(*GetCounterLevelsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_ChangeCounterLevels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeCounterLevelsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).ChangeCounterLevels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CounterService_ChangeCounterLevels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).ChangeCounterLevels(ctx, req.(*ChangeCounterLevelsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_ClearCounterLevels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearCounterLevelsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).ClearCounterLevels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CounterService_ClearCounterLevels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).ClearCounterLevels(ctx, req.(*ClearCounterLevelsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CounterService_ServiceDesc is the grpc.ServiceDesc for CounterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CounterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "counter.v1.CounterService",
	HandlerType: (*CounterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCounter",
			Handler:    _CounterService_AddCounter_Handler,
		},
		{
			MethodName: "GetCounterLevels",
			Handler:    _CounterService_GetCounterLevels_Handler,
		},
		{
			MethodName: "ChangeCounterLevels",
			Handler:    _CounterService_ChangeCounterLevels_Handler,
		},
		{
			MethodName: "ClearCounterLevels",
			Handler:    _CounterService_ClearCounterLevels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "counter/v1/counter.proto",
}
