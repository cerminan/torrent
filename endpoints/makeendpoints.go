package endpoints

import (
  "context"

  "github.com/cerminan/torrent/service"
  "github.com/go-kit/kit/endpoint"
)

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

type File struct {
  TorrentHash string
  Name string
  Length int64
}

type FilesRequst struct {
  Magnet string
}

type FilesResponse struct {
  Files []File
}

func makeFilesEndpoint(s service.Service) endpoint.Endpoint {
  return func(ctx context.Context, request interface{}) (response interface{}, err error) {
    var req FilesRequst
    req = request.(FilesRequst)
    
    var files []service.File
    files, err = s.Files(ctx, req.Magnet)
    if err != nil {
      return nil, err
    }
    
    var res FilesResponse
    res = FilesResponse{Files: make([]File, 0)}
    
    for _, file := range files {
      res.Files = append(res.Files, File(file))
    }

    return res, nil
  }
}

type ReadAtRequest struct {
  File File
  Off int64
  Ln int64
}

type ReadAtResponse struct {
  Buffer []byte
}

func makeReadAtEndpoint(s service.Service) endpoint.Endpoint {
  return func(ctx context.Context, request interface{}) (response interface{}, err error) {
    var req ReadAtRequest
    req = request.(ReadAtRequest)
 
    var res ReadAtResponse
    res = ReadAtResponse{
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
