// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: application/v1/application.proto

package v1

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
	ApplicationSvc_CreateApplication_FullMethodName = "/api.application.v1.ApplicationSvc/CreateApplication"
	ApplicationSvc_ListApplications_FullMethodName  = "/api.application.v1.ApplicationSvc/ListApplications"
	ApplicationSvc_GetApplication_FullMethodName    = "/api.application.v1.ApplicationSvc/GetApplication"
	ApplicationSvc_DeleteApplication_FullMethodName = "/api.application.v1.ApplicationSvc/DeleteApplication"
)

// ApplicationSvcClient is the client API for ApplicationSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApplicationSvcClient interface {
	CreateApplication(ctx context.Context, in *CreateAppRequest, opts ...grpc.CallOption) (*Application, error)
	ListApplications(ctx context.Context, in *ListAppRequest, opts ...grpc.CallOption) (*ListAppResponse, error)
	GetApplication(ctx context.Context, in *GetAppRequest, opts ...grpc.CallOption) (*Application, error)
	DeleteApplication(ctx context.Context, in *DeleteAppRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type applicationSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewApplicationSvcClient(cc grpc.ClientConnInterface) ApplicationSvcClient {
	return &applicationSvcClient{cc}
}

func (c *applicationSvcClient) CreateApplication(ctx context.Context, in *CreateAppRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, ApplicationSvc_CreateApplication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationSvcClient) ListApplications(ctx context.Context, in *ListAppRequest, opts ...grpc.CallOption) (*ListAppResponse, error) {
	out := new(ListAppResponse)
	err := c.cc.Invoke(ctx, ApplicationSvc_ListApplications_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationSvcClient) GetApplication(ctx context.Context, in *GetAppRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, ApplicationSvc_GetApplication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationSvcClient) DeleteApplication(ctx context.Context, in *DeleteAppRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, ApplicationSvc_DeleteApplication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationSvcServer is the server API for ApplicationSvc service.
// All implementations must embed UnimplementedApplicationSvcServer
// for forward compatibility
type ApplicationSvcServer interface {
	CreateApplication(context.Context, *CreateAppRequest) (*Application, error)
	ListApplications(context.Context, *ListAppRequest) (*ListAppResponse, error)
	GetApplication(context.Context, *GetAppRequest) (*Application, error)
	DeleteApplication(context.Context, *DeleteAppRequest) (*EmptyResponse, error)
	mustEmbedUnimplementedApplicationSvcServer()
}

// UnimplementedApplicationSvcServer must be embedded to have forward compatible implementations.
type UnimplementedApplicationSvcServer struct {
}

func (UnimplementedApplicationSvcServer) CreateApplication(context.Context, *CreateAppRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApplication not implemented")
}
func (UnimplementedApplicationSvcServer) ListApplications(context.Context, *ListAppRequest) (*ListAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApplications not implemented")
}
func (UnimplementedApplicationSvcServer) GetApplication(context.Context, *GetAppRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApplication not implemented")
}
func (UnimplementedApplicationSvcServer) DeleteApplication(context.Context, *DeleteAppRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApplication not implemented")
}
func (UnimplementedApplicationSvcServer) mustEmbedUnimplementedApplicationSvcServer() {}

// UnsafeApplicationSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApplicationSvcServer will
// result in compilation errors.
type UnsafeApplicationSvcServer interface {
	mustEmbedUnimplementedApplicationSvcServer()
}

func RegisterApplicationSvcServer(s grpc.ServiceRegistrar, srv ApplicationSvcServer) {
	s.RegisterService(&ApplicationSvc_ServiceDesc, srv)
}

func _ApplicationSvc_CreateApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationSvcServer).CreateApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApplicationSvc_CreateApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationSvcServer).CreateApplication(ctx, req.(*CreateAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationSvc_ListApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationSvcServer).ListApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApplicationSvc_ListApplications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationSvcServer).ListApplications(ctx, req.(*ListAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationSvc_GetApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationSvcServer).GetApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApplicationSvc_GetApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationSvcServer).GetApplication(ctx, req.(*GetAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationSvc_DeleteApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationSvcServer).DeleteApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApplicationSvc_DeleteApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationSvcServer).DeleteApplication(ctx, req.(*DeleteAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApplicationSvc_ServiceDesc is the grpc.ServiceDesc for ApplicationSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApplicationSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.application.v1.ApplicationSvc",
	HandlerType: (*ApplicationSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApplication",
			Handler:    _ApplicationSvc_CreateApplication_Handler,
		},
		{
			MethodName: "ListApplications",
			Handler:    _ApplicationSvc_ListApplications_Handler,
		},
		{
			MethodName: "GetApplication",
			Handler:    _ApplicationSvc_GetApplication_Handler,
		},
		{
			MethodName: "DeleteApplication",
			Handler:    _ApplicationSvc_DeleteApplication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "application/v1/application.proto",
}
