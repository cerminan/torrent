package transport

import (
	"context"

	"github.com/cerminan/torrent/endpoints"
	"github.com/cerminan/torrent/pb"
	"github.com/cerminan/torrent/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

type gRPCServer struct {
  fetch gt.Handler
  files gt.Handler
  readat gt.Handler
  pb.UnimplementedTorrentServer
}

func NewGRPCClient(conn *grpc.ClientConn) service.Service {
  var files endpoint.Endpoint
  files = gt.NewClient(
    conn,
    "Torrent",
    "Files",
    encodeFilesRequest,
    decodeFilesResponse,
    pb.FilesRes{},
  ).Endpoint()

  var readAt endpoint.Endpoint
  readAt = gt.NewClient(
    conn,
    "Torrent",
    "ReadAt",
    encodeReadAtRequest,
    decodeReadAtResponse,
    pb.ReadAtRes{},
  ).Endpoint()

  return endpoints.Endpoints{
    FilesEndpoint: files,
    ReadAtEndpoint: readAt, 
  }
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

func encodeFilesRequest(_ context.Context, request interface{}) (interface{}, error){
  var req endpoints.FilesReq
  req = request.(endpoints.FilesReq)
  return &pb.FilesReq{Magnet: req.Magnet}, nil
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

func decodeFilesResponse(_ context.Context, response interface{}) (interface{}, error){
  var res *pb.FilesRes
  res = response.(*pb.FilesRes)

  var files []endpoints.File
  for _, file := range res.Files {
    files = append(files, endpoints.File{
      TorrentHash: file.TorrentHash,
      Name: file.Name,
      Length: file.Length,
    })
  }

  return endpoints.FilesRes{Files: files}, nil
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

func encodeReadAtRequest(_ context.Context, request interface{}) (interface{}, error) {
  var req endpoints.ReadAtReq
  req = request.(endpoints.ReadAtReq)
  
  var reqFile *pb.File
  reqFile = &pb.File{
    TorrentHash: req.File.TorrentHash,
    Name: req.File.Name,
    Length: req.File.Length,
  }

  return &pb.ReadAtReq{File: reqFile, Off: req.Off, Ln: req.Ln}, nil
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

func decodeReadAtResponse(_ context.Context, response interface{}) (interface{}, error) {
  var res *pb.ReadAtRes
  res = response.(*pb.ReadAtRes)
  return endpoints.ReadAtRes{Buffer: res.Buffer}, nil
}
