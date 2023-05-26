// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: stanza/hub/v1/quota.proto

package hubv1

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
	QuotaService_GetToken_FullMethodName              = "/stanza.hub.v1.QuotaService/GetToken"
	QuotaService_GetTokenLease_FullMethodName         = "/stanza.hub.v1.QuotaService/GetTokenLease"
	QuotaService_SetTokenLeaseConsumed_FullMethodName = "/stanza.hub.v1.QuotaService/SetTokenLeaseConsumed"
	QuotaService_ValidateToken_FullMethodName         = "/stanza.hub.v1.QuotaService/ValidateToken"
)

// QuotaServiceClient is the client API for QuotaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuotaServiceClient interface {
	// Required for V0. Issue: https://github.com/StanzaSystems/stanza-hub/issues/25
	GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error)
	// All rpcs below this required for V1. Issue: https://github.com/StanzaSystems/stanza-hub/issues/120
	GetTokenLease(ctx context.Context, in *GetTokenLeaseRequest, opts ...grpc.CallOption) (*GetTokenLeaseResponse, error)
	SetTokenLeaseConsumed(ctx context.Context, in *SetTokenLeaseConsumedRequest, opts ...grpc.CallOption) (*SetTokenLeaseConsumedResponse, error)
	// Used by ingress decorators to validate Hub-generated tokens.
	ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error)
}

type quotaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuotaServiceClient(cc grpc.ClientConnInterface) QuotaServiceClient {
	return &quotaServiceClient{cc}
}

func (c *quotaServiceClient) GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error) {
	out := new(GetTokenResponse)
	err := c.cc.Invoke(ctx, QuotaService_GetToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotaServiceClient) GetTokenLease(ctx context.Context, in *GetTokenLeaseRequest, opts ...grpc.CallOption) (*GetTokenLeaseResponse, error) {
	out := new(GetTokenLeaseResponse)
	err := c.cc.Invoke(ctx, QuotaService_GetTokenLease_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotaServiceClient) SetTokenLeaseConsumed(ctx context.Context, in *SetTokenLeaseConsumedRequest, opts ...grpc.CallOption) (*SetTokenLeaseConsumedResponse, error) {
	out := new(SetTokenLeaseConsumedResponse)
	err := c.cc.Invoke(ctx, QuotaService_SetTokenLeaseConsumed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotaServiceClient) ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error) {
	out := new(ValidateTokenResponse)
	err := c.cc.Invoke(ctx, QuotaService_ValidateToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuotaServiceServer is the server API for QuotaService service.
// All implementations should embed UnimplementedQuotaServiceServer
// for forward compatibility
type QuotaServiceServer interface {
	// Required for V0. Issue: https://github.com/StanzaSystems/stanza-hub/issues/25
	GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error)
	// All rpcs below this required for V1. Issue: https://github.com/StanzaSystems/stanza-hub/issues/120
	GetTokenLease(context.Context, *GetTokenLeaseRequest) (*GetTokenLeaseResponse, error)
	SetTokenLeaseConsumed(context.Context, *SetTokenLeaseConsumedRequest) (*SetTokenLeaseConsumedResponse, error)
	// Used by ingress decorators to validate Hub-generated tokens.
	ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error)
}

// UnimplementedQuotaServiceServer should be embedded to have forward compatible implementations.
type UnimplementedQuotaServiceServer struct {
}

func (UnimplementedQuotaServiceServer) GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedQuotaServiceServer) GetTokenLease(context.Context, *GetTokenLeaseRequest) (*GetTokenLeaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenLease not implemented")
}
func (UnimplementedQuotaServiceServer) SetTokenLeaseConsumed(context.Context, *SetTokenLeaseConsumedRequest) (*SetTokenLeaseConsumedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTokenLeaseConsumed not implemented")
}
func (UnimplementedQuotaServiceServer) ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}

// UnsafeQuotaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuotaServiceServer will
// result in compilation errors.
type UnsafeQuotaServiceServer interface {
	mustEmbedUnimplementedQuotaServiceServer()
}

func RegisterQuotaServiceServer(s grpc.ServiceRegistrar, srv QuotaServiceServer) {
	s.RegisterService(&QuotaService_ServiceDesc, srv)
}

func _QuotaService_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotaServiceServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuotaService_GetToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotaServiceServer).GetToken(ctx, req.(*GetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuotaService_GetTokenLease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenLeaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotaServiceServer).GetTokenLease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuotaService_GetTokenLease_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotaServiceServer).GetTokenLease(ctx, req.(*GetTokenLeaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuotaService_SetTokenLeaseConsumed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTokenLeaseConsumedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotaServiceServer).SetTokenLeaseConsumed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuotaService_SetTokenLeaseConsumed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotaServiceServer).SetTokenLeaseConsumed(ctx, req.(*SetTokenLeaseConsumedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuotaService_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotaServiceServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuotaService_ValidateToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotaServiceServer).ValidateToken(ctx, req.(*ValidateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuotaService_ServiceDesc is the grpc.ServiceDesc for QuotaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuotaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stanza.hub.v1.QuotaService",
	HandlerType: (*QuotaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetToken",
			Handler:    _QuotaService_GetToken_Handler,
		},
		{
			MethodName: "GetTokenLease",
			Handler:    _QuotaService_GetTokenLease_Handler,
		},
		{
			MethodName: "SetTokenLeaseConsumed",
			Handler:    _QuotaService_SetTokenLeaseConsumed_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _QuotaService_ValidateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stanza/hub/v1/quota.proto",
}
