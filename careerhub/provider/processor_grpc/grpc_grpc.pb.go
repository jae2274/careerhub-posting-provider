// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: careerhub/provider/processor_grpc/grpc.proto

package processor_grpc

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

// DataProcessorClient is the client API for DataProcessor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataProcessorClient interface {
	CloseJobPostings(ctx context.Context, in *JobPostings, opts ...grpc.CallOption) (*BoolResponse, error)
	RegisterJobPostingInfo(ctx context.Context, in *JobPostingInfo, opts ...grpc.CallOption) (*BoolResponse, error)
	RegisterCompany(ctx context.Context, in *Company, opts ...grpc.CallOption) (*BoolResponse, error)
}

type dataProcessorClient struct {
	cc grpc.ClientConnInterface
}

func NewDataProcessorClient(cc grpc.ClientConnInterface) DataProcessorClient {
	return &dataProcessorClient{cc}
}

func (c *dataProcessorClient) CloseJobPostings(ctx context.Context, in *JobPostings, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/careerhub.processor.processor_grpc.DataProcessor/CloseJobPostings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataProcessorClient) RegisterJobPostingInfo(ctx context.Context, in *JobPostingInfo, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/careerhub.processor.processor_grpc.DataProcessor/RegisterJobPostingInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataProcessorClient) RegisterCompany(ctx context.Context, in *Company, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/careerhub.processor.processor_grpc.DataProcessor/RegisterCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataProcessorServer is the server API for DataProcessor service.
// All implementations must embed UnimplementedDataProcessorServer
// for forward compatibility
type DataProcessorServer interface {
	CloseJobPostings(context.Context, *JobPostings) (*BoolResponse, error)
	RegisterJobPostingInfo(context.Context, *JobPostingInfo) (*BoolResponse, error)
	RegisterCompany(context.Context, *Company) (*BoolResponse, error)
	mustEmbedUnimplementedDataProcessorServer()
}

// UnimplementedDataProcessorServer must be embedded to have forward compatible implementations.
type UnimplementedDataProcessorServer struct {
}

func (UnimplementedDataProcessorServer) CloseJobPostings(context.Context, *JobPostings) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseJobPostings not implemented")
}
func (UnimplementedDataProcessorServer) RegisterJobPostingInfo(context.Context, *JobPostingInfo) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterJobPostingInfo not implemented")
}
func (UnimplementedDataProcessorServer) RegisterCompany(context.Context, *Company) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterCompany not implemented")
}
func (UnimplementedDataProcessorServer) mustEmbedUnimplementedDataProcessorServer() {}

// UnsafeDataProcessorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataProcessorServer will
// result in compilation errors.
type UnsafeDataProcessorServer interface {
	mustEmbedUnimplementedDataProcessorServer()
}

func RegisterDataProcessorServer(s grpc.ServiceRegistrar, srv DataProcessorServer) {
	s.RegisterService(&DataProcessor_ServiceDesc, srv)
}

func _DataProcessor_CloseJobPostings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobPostings)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataProcessorServer).CloseJobPostings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/careerhub.processor.processor_grpc.DataProcessor/CloseJobPostings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataProcessorServer).CloseJobPostings(ctx, req.(*JobPostings))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataProcessor_RegisterJobPostingInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobPostingInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataProcessorServer).RegisterJobPostingInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/careerhub.processor.processor_grpc.DataProcessor/RegisterJobPostingInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataProcessorServer).RegisterJobPostingInfo(ctx, req.(*JobPostingInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataProcessor_RegisterCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Company)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataProcessorServer).RegisterCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/careerhub.processor.processor_grpc.DataProcessor/RegisterCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataProcessorServer).RegisterCompany(ctx, req.(*Company))
	}
	return interceptor(ctx, in, info, handler)
}

// DataProcessor_ServiceDesc is the grpc.ServiceDesc for DataProcessor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataProcessor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "careerhub.processor.processor_grpc.DataProcessor",
	HandlerType: (*DataProcessorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CloseJobPostings",
			Handler:    _DataProcessor_CloseJobPostings_Handler,
		},
		{
			MethodName: "RegisterJobPostingInfo",
			Handler:    _DataProcessor_RegisterJobPostingInfo_Handler,
		},
		{
			MethodName: "RegisterCompany",
			Handler:    _DataProcessor_RegisterCompany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "careerhub/provider/processor_grpc/grpc.proto",
}
