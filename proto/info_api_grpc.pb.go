// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: info_api.proto

package skyhook

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

// NodeInfoClient is the client API for NodeInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeInfoClient interface {
	GetNodeHardwareInfo(ctx context.Context, in *HardwareInfoRequest, opts ...grpc.CallOption) (*HardwareInfoResponse, error)
	GetFreePort(ctx context.Context, in *GetFreePortRequest, opts ...grpc.CallOption) (*GetFreePortResponse, error)
}

type nodeInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeInfoClient(cc grpc.ClientConnInterface) NodeInfoClient {
	return &nodeInfoClient{cc}
}

func (c *nodeInfoClient) GetNodeHardwareInfo(ctx context.Context, in *HardwareInfoRequest, opts ...grpc.CallOption) (*HardwareInfoResponse, error) {
	out := new(HardwareInfoResponse)
	err := c.cc.Invoke(ctx, "/NodeInfo/GetNodeHardwareInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeInfoClient) GetFreePort(ctx context.Context, in *GetFreePortRequest, opts ...grpc.CallOption) (*GetFreePortResponse, error) {
	out := new(GetFreePortResponse)
	err := c.cc.Invoke(ctx, "/NodeInfo/GetFreePort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeInfoServer is the server API for NodeInfo service.
// All implementations must embed UnimplementedNodeInfoServer
// for forward compatibility
type NodeInfoServer interface {
	GetNodeHardwareInfo(context.Context, *HardwareInfoRequest) (*HardwareInfoResponse, error)
	GetFreePort(context.Context, *GetFreePortRequest) (*GetFreePortResponse, error)
	mustEmbedUnimplementedNodeInfoServer()
}

// UnimplementedNodeInfoServer must be embedded to have forward compatible implementations.
type UnimplementedNodeInfoServer struct {
}

func (UnimplementedNodeInfoServer) GetNodeHardwareInfo(context.Context, *HardwareInfoRequest) (*HardwareInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeHardwareInfo not implemented")
}
func (UnimplementedNodeInfoServer) GetFreePort(context.Context, *GetFreePortRequest) (*GetFreePortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFreePort not implemented")
}
func (UnimplementedNodeInfoServer) mustEmbedUnimplementedNodeInfoServer() {}

// UnsafeNodeInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeInfoServer will
// result in compilation errors.
type UnsafeNodeInfoServer interface {
	mustEmbedUnimplementedNodeInfoServer()
}

func RegisterNodeInfoServer(s grpc.ServiceRegistrar, srv NodeInfoServer) {
	s.RegisterService(&NodeInfo_ServiceDesc, srv)
}

func _NodeInfo_GetNodeHardwareInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HardwareInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeInfoServer).GetNodeHardwareInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NodeInfo/GetNodeHardwareInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeInfoServer).GetNodeHardwareInfo(ctx, req.(*HardwareInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeInfo_GetFreePort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFreePortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeInfoServer).GetFreePort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NodeInfo/GetFreePort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeInfoServer).GetFreePort(ctx, req.(*GetFreePortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeInfo_ServiceDesc is the grpc.ServiceDesc for NodeInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "NodeInfo",
	HandlerType: (*NodeInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNodeHardwareInfo",
			Handler:    _NodeInfo_GetNodeHardwareInfo_Handler,
		},
		{
			MethodName: "GetFreePort",
			Handler:    _NodeInfo_GetFreePort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "info_api.proto",
}
