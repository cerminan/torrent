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
  pb.UnimplementedTorrentServiceServer
}

func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.TorrentServiceServer {
  return &gRPCServer{
    fetch: gt.NewServer(
      endpoints.Fetch,
      decodeFetchRequest,
      encodeFetchResponse,
    ),
    files: gt.NewServer(
      endpoints.Files,
      decodeFilesRequest,
      encodeFilesRespones,
    ),
    readat: gt.NewServer(
      endpoints.ReadAt,
      decodeReadAtRequest,
      endcodeReadAtResponse,
    ),
  }
}

func (s *gRPCServer) Fetch(ctx context.Context, req *pb.FetchReq) (*pb.FetchRes, error) {
  var err error
  var res interface{}
  _, res, err = s.fetch.ServeGRPC(ctx, req)
  if err != nil {
    return nil, err
  }

  return res.(*pb.FetchRes), nil
}

func decodeFetchRequest(_ context.Context, request interface{}) (interface{}, error) {
  var req *pb.FetchReq
  req = request.(*pb.FetchReq)
  return endpoints.FetchReq{Url: req.Url}, nil
}

func encodeFetchResponse(_ context.Context, response interface{}) (interface{}, error){
  var res endpoints.FetchRes
  res = response.(endpoints.FetchRes)
  return &pb.FetchRes{Data: res.Data}, nil
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
  return endpoints.FilesReq{Data: req.Data}, nil
}


func encodeFilesRespones(_ context.Context, response interface{}) (interface{}, error){
  var res endpoints.FilesRes
  res = response.(endpoints.FilesRes)

  var files []*pb.File
  for _, file := range res.Files {
    files = append(files, &pb.File{Name: file.Name, Length: file.Length})
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
  return endpoints.ReadAtReq{Data: req.Data, Index: req.Index, Off: req.Off, Ln: req.Ln}, nil
}

func endcodeReadAtResponse(_ context.Context, response interface{}) (interface{}, error) {
  var res endpoints.ReadAtRes
  res = response.(endpoints.ReadAtRes)
  return &pb.ReadAtRes{Buffer: res.Buffer}, nil
}

