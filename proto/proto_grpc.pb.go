// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: proto/proto.proto

package proto

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

// TimeAskServiceClient is the client API for TimeAskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimeAskServiceClient interface {
	GetTime(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessageAck, error)
}

type timeAskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTimeAskServiceClient(cc grpc.ClientConnInterface) TimeAskServiceClient {
	return &timeAskServiceClient{cc}
}

func (c *timeAskServiceClient) GetTime(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessageAck, error) {
	out := new(MessageAck)
	err := c.cc.Invoke(ctx, "/proto.TimeAskService/GetTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeAskServiceServer is the server API for TimeAskService service.
// All implementations must embed UnimplementedTimeAskServiceServer
// for forward compatibility
type TimeAskServiceServer interface {
	GetTime(context.Context, *Message) (*MessageAck, error)
	mustEmbedUnimplementedTimeAskServiceServer()
}

// UnimplementedTimeAskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTimeAskServiceServer struct {
}

func (UnimplementedTimeAskServiceServer) GetTime(context.Context, *Message) (*MessageAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTime not implemented")
}
func (UnimplementedTimeAskServiceServer) mustEmbedUnimplementedTimeAskServiceServer() {}

// UnsafeTimeAskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimeAskServiceServer will
// result in compilation errors.
type UnsafeTimeAskServiceServer interface {
	mustEmbedUnimplementedTimeAskServiceServer()
}

func RegisterTimeAskServiceServer(s grpc.ServiceRegistrar, srv TimeAskServiceServer) {
	s.RegisterService(&TimeAskService_ServiceDesc, srv)
}

func _TimeAskService_GetTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeAskServiceServer).GetTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TimeAskService/GetTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeAskServiceServer).GetTime(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// TimeAskService_ServiceDesc is the grpc.ServiceDesc for TimeAskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimeAskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TimeAskService",
	HandlerType: (*TimeAskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTime",
			Handler:    _TimeAskService_GetTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/proto.proto",
}
