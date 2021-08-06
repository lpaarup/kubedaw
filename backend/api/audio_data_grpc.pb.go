// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package kubedaw

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

// AudioDataClient is the client API for AudioData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AudioDataClient interface {
	Request(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*Data, error)
}

type audioDataClient struct {
	cc grpc.ClientConnInterface
}

func NewAudioDataClient(cc grpc.ClientConnInterface) AudioDataClient {
	return &audioDataClient{cc}
}

func (c *audioDataClient) Request(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/kubedaw.AudioData/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AudioDataServer is the server API for AudioData service.
// All implementations must embed UnimplementedAudioDataServer
// for forward compatibility
type AudioDataServer interface {
	Request(context.Context, *DataRequest) (*Data, error)
	mustEmbedUnimplementedAudioDataServer()
}

// UnimplementedAudioDataServer must be embedded to have forward compatible implementations.
type UnimplementedAudioDataServer struct {
}

func (UnimplementedAudioDataServer) Request(context.Context, *DataRequest) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}
func (UnimplementedAudioDataServer) mustEmbedUnimplementedAudioDataServer() {}

// UnsafeAudioDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AudioDataServer will
// result in compilation errors.
type UnsafeAudioDataServer interface {
	mustEmbedUnimplementedAudioDataServer()
}

func RegisterAudioDataServer(s grpc.ServiceRegistrar, srv AudioDataServer) {
	s.RegisterService(&AudioData_ServiceDesc, srv)
}

func _AudioData_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudioDataServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubedaw.AudioData/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudioDataServer).Request(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AudioData_ServiceDesc is the grpc.ServiceDesc for AudioData service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AudioData_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kubedaw.AudioData",
	HandlerType: (*AudioDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _AudioData_Request_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "audio_data.proto",
}
