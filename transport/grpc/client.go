package grpc

import (
  "context"

  "github.com/cerminan/torrent/endpoints"
  "github.com/cerminan/torrent/service"
  "github.com/cerminan/torrent/transport/grpc/pb"
  "github.com/go-kit/kit/endpoint"
  kitGRPC "github.com/go-kit/kit/transport/grpc"
  "google.golang.org/grpc"
)

func NewClient(conn *grpc.ClientConn) service.Service {
  var files endpoint.Endpoint
  files = kitGRPC.NewClient(
    conn,
    "Torrent",
    "Files",
    encodeFilesRequest,
    decodeFilesResponse,
    pb.FilesResponse{},
  ).Endpoint()

  var readAt endpoint.Endpoint
  readAt = kitGRPC.NewClient(
    conn,
    "Torrent",
    "ReadAt",
    encodeReadAtRequest,
    decodeReadAtResponse,
    pb.ReadAtResponse{},
  ).Endpoint()

  var isMagnet endpoint.Endpoint
  isMagnet = kitGRPC.NewClient(
    conn,
    "Torrent",
    "IsMagnet",
    encodeIsMagnetRequest,
    decodeIsMagnetResponse,
    pb.IsMagnetResponse{},
  ).Endpoint()

  return endpoints.Endpoints{
    FilesEndpoint: files,
    ReadAtEndpoint: readAt, 
    IsMagnetEndpoint: isMagnet,
  }
}

func encodeFilesRequest(_ context.Context, request interface{}) (interface{}, error){
  var req endpoints.FilesRequst
  req = request.(endpoints.FilesRequst)
  return &pb.FilesRequest{Magnet: req.Magnet}, nil
}

func decodeFilesResponse(_ context.Context, response interface{}) (interface{}, error){
  var res *pb.FilesResponse
  res = response.(*pb.FilesResponse)

  var files []endpoints.File
  for _, file := range res.Files {
    files = append(files, endpoints.File{
      TorrentHash: file.TorrentHash,
      Name: file.Name,
      Length: file.Length,
    })
  }

  return endpoints.FilesResponse{Files: files}, nil
}

func encodeReadAtRequest(_ context.Context, request interface{}) (interface{}, error) {
  var req endpoints.ReadAtRequest
  req = request.(endpoints.ReadAtRequest)
  
  var reqFile *pb.File
  reqFile = &pb.File{
    TorrentHash: req.File.TorrentHash,
    Name: req.File.Name,
    Length: req.File.Length,
  }

  return &pb.ReadAtRequest{File: reqFile, Off: req.Off, Ln: req.Ln}, nil
}

func decodeReadAtResponse(_ context.Context, response interface{}) (interface{}, error) {
  var res *pb.ReadAtResponse
  res = response.(*pb.ReadAtResponse)
  return endpoints.ReadAtResponse{Buffer: res.Buffer}, nil
}

func encodeIsMagnetRequest(_ context.Context, request interface{}) (interface{}, error) {
  var req endpoints.IsMagnetRequest
  req = request.(endpoints.IsMagnetRequest)
  return &pb.IsMagnetRequest{Magnet: req.Magnet}, nil
}

func decodeIsMagnetResponse(_ context.Context, response interface{}) (interface{}, error) {
  var res *pb.IsMagnetResponse
  res = response.(*pb.IsMagnetResponse)
  return endpoints.IsMagnetResponse{Valid: res.Valid}, nil
}
