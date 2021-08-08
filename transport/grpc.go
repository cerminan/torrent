package transport

import (
	"context"

	"github.com/cerminan/torrent/endpoints"
	"github.com/cerminan/torrent/transport/pb"
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
  isMagnet gt.Handler
  pb.UnimplementedTorrentServer
}

func NewGRPCClient(conn *grpc.ClientConn) service.Service {
  var files endpoint.Endpoint
  files = gt.NewClient(
    conn,
    "Torrent",
    "Files",
    encodeGRPCFilesRequest,
    decodeGRPCFilesResponse,
    pb.FilesRes{},
  ).Endpoint()

  var readAt endpoint.Endpoint
  readAt = gt.NewClient(
    conn,
    "Torrent",
    "ReadAt",
    encodeGRPCReadAtRequest,
    decodeGRPCReadAtResponse,
    pb.ReadAtRes{},
  ).Endpoint()

  var isMagnet endpoint.Endpoint
  isMagnet = gt.NewClient(
    conn,
    "Torrent",
    "IsMagnet",
    encodeGRPCIsMagnetRequest,
    decodeGRPCIsMagnetResponse,
    pb.IsMagnetResponse{},
  ).Endpoint()

  return endpoints.Endpoints{
    FilesEndpoint: files,
    ReadAtEndpoint: readAt, 
    IsMagnetEndpoint: isMagnet,
  }
}

func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.TorrentServer {
  return &gRPCServer{
    files: gt.NewServer(
      endpoints.FilesEndpoint,
      decodeGRPCFilesRequest,
      encodeGRPCFilesRespones,
    ),
    readat: gt.NewServer(
      endpoints.ReadAtEndpoint,
      decodeGRPCReadAtRequest,
      endcodeGRPCReadAtResponse,
    ),
    isMagnet: gt.NewServer(
      endpoints.IsMagnetEndpoint,
      decodeGRPCIsMagnetRequest,
      encodeGRPCIsMagnetResponse,
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

func encodeGRPCFilesRequest(_ context.Context, request interface{}) (interface{}, error){
  var req endpoints.FilesReq
  req = request.(endpoints.FilesReq)
  return &pb.FilesReq{Magnet: req.Magnet}, nil
}

func decodeGRPCFilesRequest(_ context.Context, request interface{}) (interface{}, error) {
  var req *pb.FilesReq
  req = request.(*pb.FilesReq)
  return endpoints.FilesReq{Magnet: req.Magnet}, nil
}

func encodeGRPCFilesRespones(_ context.Context, response interface{}) (interface{}, error){
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

func decodeGRPCFilesResponse(_ context.Context, response interface{}) (interface{}, error){
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

func encodeGRPCReadAtRequest(_ context.Context, request interface{}) (interface{}, error) {
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

func decodeGRPCReadAtRequest(_ context.Context, request interface{}) (interface{}, error) {
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

func endcodeGRPCReadAtResponse(_ context.Context, response interface{}) (interface{}, error) {
  var res endpoints.ReadAtRes
  res = response.(endpoints.ReadAtRes)
  return &pb.ReadAtRes{Buffer: res.Buffer}, nil
}

func decodeGRPCReadAtResponse(_ context.Context, response interface{}) (interface{}, error) {
  var res *pb.ReadAtRes
  res = response.(*pb.ReadAtRes)
  return endpoints.ReadAtRes{Buffer: res.Buffer}, nil
}

func (s *gRPCServer) IsMagnet(ctx context.Context, req *pb.IsMagnetRequest) (*pb.IsMagnetResponse, error) {
  var err error
  var res interface{}
  _, res, err = s.isMagnet.ServeGRPC(ctx, req)
  if err != nil {
    return nil, err
  }

  return res.(*pb.IsMagnetResponse), nil
}

func encodeGRPCIsMagnetRequest(_ context.Context, request interface{}) (interface{}, error) {
  var req endpoints.IsMagnetRequest
  req = request.(endpoints.IsMagnetRequest)
  return &pb.IsMagnetRequest{Magnet: req.Magnet}, nil
}

func decodeGRPCIsMagnetRequest(_ context.Context, request interface{}) (interface{}, error) {
  var req *pb.IsMagnetRequest
  req = request.(*pb.IsMagnetRequest)
  return endpoints.IsMagnetRequest{Magnet: req.Magnet}, nil
}

func encodeGRPCIsMagnetResponse(_ context.Context, response interface{}) (interface{}, error) {
  var res endpoints.IsMagnetResponse
  res = response.(endpoints.IsMagnetResponse)
  return &pb.IsMagnetResponse{Valid: res.Valid}, nil
}

func decodeGRPCIsMagnetResponse(_ context.Context, response interface{}) (interface{}, error) {
  var res *pb.IsMagnetResponse
  res = response.(*pb.IsMagnetResponse)
  return endpoints.IsMagnetResponse{Valid: res.Valid}, nil
}
