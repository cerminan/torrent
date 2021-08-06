package transport

import (
	"context"

	"github.com/cerminan/torrent/endpoints"
	"github.com/cerminan/torrent/pb"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
  fetch gt.Handler
  files gt.Handler
  readat gt.Handler
  pb.UnimplementedTorrentServer
}

func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.TorrentServer {
  return &gRPCServer{
    files: gt.NewServer(
      endpoints.FilesEndpoint,
      decodeFilesRequest,
      encodeFilesRespones,
    ),
    readat: gt.NewServer(
      endpoints.ReadAtEndpoint,
      decodeReadAtRequest,
      endcodeReadAtResponse,
    ),
  }
}

func (s *gRPCServer) Files(ctx context.Context, req *pb.FilesReq) (*pb.FilesRes, error) {
  var err error
  var res interface{}
  _, res, err = s.files.ServeGRPC(ctx, req)
  if err != nil {
    return nil, err
  }

  return res.(*pb.FilesRes), nil
}

func decodeFilesRequest(_ context.Context, request interface{}) (interface{}, error) {
  var req *pb.FilesReq
  req = request.(*pb.FilesReq)
  return endpoints.FilesReq{Magnet: req.Magnet}, nil
}

func encodeFilesRespones(_ context.Context, response interface{}) (interface{}, error){
  var res endpoints.FilesRes
  res = response.(endpoints.FilesRes)

  var files []*pb.File
  for _, file := range res.Files {
    files = append(files, &pb.File{
      TorrentHash: file.TorrentHash,
      Name: file.Name,
      Length: file.Length,
    })
  }
  return &pb.FilesRes{Files: files}, nil
}

func (s *gRPCServer) ReadAt(ctx context.Context, req *pb.ReadAtReq) (*pb.ReadAtRes, error) {
  var err error
  var resp interface{}
  _, resp, err = s.readat.ServeGRPC(ctx, req)
  if err != nil {
    return nil, err
  }

  return resp.(*pb.ReadAtRes), nil
}

func decodeReadAtRequest(_ context.Context, request interface{}) (interface{}, error) {
  var req *pb.ReadAtReq
  req = request.(*pb.ReadAtReq)

  var reqFile endpoints.File
  reqFile = endpoints.File{
    TorrentHash: req.File.TorrentHash,
    Name: req.File.Name,
    Length: req.File.Length,
  }

  return endpoints.ReadAtReq{File: reqFile, Off: req.Off, Ln: req.Ln}, nil
}

func endcodeReadAtResponse(_ context.Context, response interface{}) (interface{}, error) {
  var res endpoints.ReadAtRes
  res = response.(endpoints.ReadAtRes)
  return &pb.ReadAtRes{Buffer: res.Buffer}, nil
}

