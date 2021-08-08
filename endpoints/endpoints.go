package endpoints

import (
	"context"

	"github.com/cerminan/torrent/service"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
  FilesEndpoint endpoint.Endpoint
  ReadAtEndpoint endpoint.Endpoint
  IsMagnetEndpoint endpoint.Endpoint
}

func (e Endpoints) Files(ctx context.Context, magnet string) ([]service.File, error) {
  var err error

  var rawRes interface{}
  rawRes, err = e.FilesEndpoint(ctx, FilesReq{Magnet: magnet})
  if err != nil {
    return nil, err
  }
  
  var res = rawRes.(FilesRes)

  var files []service.File
  for _, file := range res.Files {
    files = append(files, service.File{
      TorrentHash: file.TorrentHash,
      Name: file.Name,
      Length: file.Length,
    })
  }

  return files, nil
}

func (e Endpoints) ReadAt(ctx context.Context, file service.File, off int64, ln int64) ([]byte, error) {
  var err error
  
  var rawRes interface{}
  rawRes, err = e.ReadAtEndpoint(ctx, ReadAtReq{File: File(file), Off: off, Ln: ln})
  if err != nil {
    return nil, err
  }

  var res ReadAtRes
  res = rawRes.(ReadAtRes)
  
  return res.Buffer, nil
}

func (e Endpoints) IsMagnet(ctx context.Context, magnet string) (bool, error) {
  var err error
  
  var rawRes interface{}
  rawRes, err = e.IsMagnetEndpoint(ctx, IsMagnetRequest{Magnet: magnet})
  if err != nil {
    return false, err
  }

  var res IsMagnetResponse
  res = rawRes.(IsMagnetResponse)

  return res.Valid, nil
}

type FilesReq struct {
  Magnet string
}

type FilesRes struct {
  Files []File
}

type File struct {
  TorrentHash string
  Name string
  Length int64
}

type ReadAtReq struct {
  File File
  Off int64
  Ln int64
}

type ReadAtRes struct {
  Buffer []byte
}

func MakeEndpoints(s service.Service) Endpoints {
  var filesEndpoint endpoint.Endpoint
  {
    filesEndpoint = makeFilesEndpoint(s)
  }

  var readAtEndpoint endpoint.Endpoint
  {
    readAtEndpoint = makeReadAtEndpoint(s)
  }

  var isMagnetEnpoint endpoint.Endpoint
  {
    isMagnetEnpoint = makeIsMagnetEndpoint(s)
  }
  
  return Endpoints{
    FilesEndpoint: filesEndpoint,
    ReadAtEndpoint: readAtEndpoint,
    IsMagnetEndpoint: isMagnetEnpoint,
  }
}

func makeFilesEndpoint(s service.Service) endpoint.Endpoint {
  return func(ctx context.Context, request interface{}) (response interface{}, err error) {
    var req FilesReq
    req = request.(FilesReq)
    
    var files []service.File
    files, err = s.Files(ctx, req.Magnet)
    if err != nil {
      return nil, err
    }
    
    var res FilesRes
    res = FilesRes{Files: make([]File, 0)}
    
    for _, file := range files {
      res.Files = append(res.Files, File(file))
    }

    return res, nil
  }
}

func makeReadAtEndpoint(s service.Service) endpoint.Endpoint {
  return func(ctx context.Context, request interface{}) (response interface{}, err error) {
    var req ReadAtReq
    req = request.(ReadAtReq)
 
    var res ReadAtRes
    res = ReadAtRes{
      Buffer: make([]byte, req.Ln),
    }

    res.Buffer, err = s.ReadAt(ctx, service.File(req.File), req.Off, req.Ln)
    if err != nil {
      return nil, err
    }

    return res, nil
  }
}

type IsMagnetRequest struct {
  Magnet string
}

type IsMagnetResponse struct {
  Valid bool
}

func makeIsMagnetEndpoint(s service.Service) endpoint.Endpoint {
  return func(ctx context.Context, request interface{}) (response interface{}, err error) {
    var req IsMagnetRequest
    req = request.(IsMagnetRequest)

    var res IsMagnetResponse
    res.Valid, err = s.IsMagnet(ctx, req.Magnet)
    if err != nil {
      return nil, err
    }

    return res, nil
  }
}
