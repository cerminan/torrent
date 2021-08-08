package grpc

import (
	"context"

	"github.com/cerminan/torrent/endpoints"
	"github.com/cerminan/torrent/transport/grpc/pb"
	kitGRPC "github.com/go-kit/kit/transport/grpc"
)

type server struct {
  fetch kitGRPC.Handler
  files kitGRPC.Handler
  readat kitGRPC.Handler
  isMagnet kitGRPC.Handler
  pb.UnimplementedTorrentServer
}

func NewServer(endpoints endpoints.Endpoints) pb.TorrentServer {
  return &server{
    files: kitGRPC.NewServer(
      endpoints.FilesEndpoint,
      decodeFilesRequest,
      encodeFilesRespones,
    ),
    readat: kitGRPC.NewServer(
      endpoints.ReadAtEndpoint,
      decodeReadAtRequest,
      encodeReadAtResponse,
    ),
    isMagnet: kitGRPC.NewServer(
      endpoints.IsMagnetEndpoint,
      decodeIsMagnetRequest,
      encodeIsMagnetResponse,
    ),
  }
}

func (s *server) Files(ctx context.Context, req *pb.FilesRequest) (*pb.FilesResponse, error) {
  var err error
  var res interface{}
  _, res, err = s.files.ServeGRPC(ctx, req)
  if err != nil {
    return nil, err
  }

  return res.(*pb.FilesResponse), nil
}

func decodeFilesRequest(_ context.Context, request interface{}) (interface{}, error) {
  var req *pb.FilesRequest
  req = request.(*pb.FilesRequest)
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
  return &pb.FilesResponse{Files: files}, nil
}

func (s *server) ReadAt(ctx context.Context, req *pb.ReadAtRequest) (*pb.ReadAtResponse, error) {
  var err error
  var resp interface{}
  _, resp, err = s.readat.ServeGRPC(ctx, req)
  if err != nil {
    return nil, err
  }

  return resp.(*pb.ReadAtResponse), nil
}

func decodeReadAtRequest(_ context.Context, request interface{}) (interface{}, error) {
  var req *pb.ReadAtRequest
  req = request.(*pb.ReadAtRequest)

  var reqFile endpoints.File
  reqFile = endpoints.File{
    TorrentHash: req.File.TorrentHash,
    Name: req.File.Name,
    Length: req.File.Length,
  }

  return endpoints.ReadAtReq{File: reqFile, Off: req.Off, Ln: req.Ln}, nil
}

func encodeReadAtResponse(_ context.Context, response interface{}) (interface{}, error) {
  var res endpoints.ReadAtRes
  res = response.(endpoints.ReadAtRes)
  return &pb.ReadAtResponse{Buffer: res.Buffer}, nil
}

func (s *server) IsMagnet(ctx context.Context, req *pb.IsMagnetRequest) (*pb.IsMagnetResponse, error) {
  var err error
  var res interface{}
  _, res, err = s.isMagnet.ServeGRPC(ctx, req)
  if err != nil {
    return nil, err
  }

  return res.(*pb.IsMagnetResponse), nil
}

func decodeIsMagnetRequest(_ context.Context, request interface{}) (interface{}, error) {
  var req *pb.IsMagnetRequest
  req = request.(*pb.IsMagnetRequest)
  return endpoints.IsMagnetRequest{Magnet: req.Magnet}, nil
}

func encodeIsMagnetResponse(_ context.Context, response interface{}) (interface{}, error) {
  var res endpoints.IsMagnetResponse
  res = response.(endpoints.IsMagnetResponse)
  return &pb.IsMagnetResponse{Valid: res.Valid}, nil
}
