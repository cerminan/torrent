// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// TorrentClient is the client API for Torrent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TorrentClient interface {
	Files(ctx context.Context, in *FilesReq, opts ...grpc.CallOption) (*FilesRes, error)
	ReadAt(ctx context.Context, in *ReadAtReq, opts ...grpc.CallOption) (*ReadAtRes, error)
}

type torrentClient struct {
	cc grpc.ClientConnInterface
}

func NewTorrentClient(cc grpc.ClientConnInterface) TorrentClient {
	return &torrentClient{cc}
}

func (c *torrentClient) Files(ctx context.Context, in *FilesReq, opts ...grpc.CallOption) (*FilesRes, error) {
	out := new(FilesRes)
	err := c.cc.Invoke(ctx, "/Torrent/Files", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *torrentClient) ReadAt(ctx context.Context, in *ReadAtReq, opts ...grpc.CallOption) (*ReadAtRes, error) {
	out := new(ReadAtRes)
	err := c.cc.Invoke(ctx, "/Torrent/ReadAt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TorrentServer is the server API for Torrent service.
// All implementations must embed UnimplementedTorrentServer
// for forward compatibility
type TorrentServer interface {
	Files(context.Context, *FilesReq) (*FilesRes, error)
	ReadAt(context.Context, *ReadAtReq) (*ReadAtRes, error)
	mustEmbedUnimplementedTorrentServer()
}

// UnimplementedTorrentServer must be embedded to have forward compatible implementations.
type UnimplementedTorrentServer struct {
}

func (UnimplementedTorrentServer) Files(context.Context, *FilesReq) (*FilesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Files not implemented")
}
func (UnimplementedTorrentServer) ReadAt(context.Context, *ReadAtReq) (*ReadAtRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAt not implemented")
}
func (UnimplementedTorrentServer) mustEmbedUnimplementedTorrentServer() {}

// UnsafeTorrentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TorrentServer will
// result in compilation errors.
type UnsafeTorrentServer interface {
	mustEmbedUnimplementedTorrentServer()
}

func RegisterTorrentServer(s grpc.ServiceRegistrar, srv TorrentServer) {
	s.RegisterService(&Torrent_ServiceDesc, srv)
}

func _Torrent_Files_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TorrentServer).Files(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Torrent/Files",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TorrentServer).Files(ctx, req.(*FilesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Torrent_ReadAt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAtReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TorrentServer).ReadAt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Torrent/ReadAt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TorrentServer).ReadAt(ctx, req.(*ReadAtReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Torrent_ServiceDesc is the grpc.ServiceDesc for Torrent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Torrent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Torrent",
	HandlerType: (*TorrentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Files",
			Handler:    _Torrent_Files_Handler,
		},
		{
			MethodName: "ReadAt",
			Handler:    _Torrent_ReadAt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "torrent.proto",
}