package endpoints

import (
	"bytes"
	"context"
	"encoding/gob"
	"io"

	"github.com/cerminan/torrent/service"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
  Fetch endpoint.Endpoint
  Files endpoint.Endpoint
  ReadAt endpoint.Endpoint
}

type FetchReq struct {
  Url string
}

type FetchRes struct {
  Data []byte
}

type FilesReq struct {
  Data []byte 
}

type FilesRes struct {
  Files []service.File
}

type ReadAtReq struct {
  Data []byte
  Index int32
  Off int64
  Ln int64
}

type ReadAtRes struct {
  Buffer []byte
}

func MakeEndpoints(s service.Service) Endpoints {
  return Endpoints{
    Fetch: makeFetchEndpoint(s),
    Files: makeFilesEndpoint(s),
    ReadAt: makeReadAtEndpoint(s),
  }
}

func makeFetchEndpoint(s service.Service) endpoint.Endpoint {
  return func(ctx context.Context, request interface{}) (response interface{}, err error) {
    var req FetchReq
    req = request.(FetchReq)

    var buffer bytes.Buffer
    var encoder *gob.Encoder
    encoder = gob.NewEncoder(&buffer)

    var torrent service.Torrent
    torrent, err = s.Fetch(ctx, req.Url)
    if err != nil {
      return nil, err
    }

    err = encoder.Encode(torrent)
    if err != nil {
      return nil, err
    }

    return FetchRes{Data: buffer.Bytes()}, nil
  }
}

func makeFilesEndpoint(s service.Service) endpoint.Endpoint {
  return func(ctx context.Context, request interface{}) (response interface{}, err error) {
    var req FilesReq
    req = request.(FilesReq)
    
    var torrent service.Torrent
    torrent, err = bytes2torrent(req.Data)
    if err != nil {
      return nil, err
    }

    return FilesRes{Files: torrent.Files}, nil
  }
}

func makeReadAtEndpoint(s service.Service) endpoint.Endpoint {
  return func(ctx context.Context, request interface{}) (response interface{}, err error) {
    var req ReadAtReq
    req = request.(ReadAtReq)
 
    var torrent service.Torrent
    torrent, err = bytes2torrent(req.Data)
    if err != nil {
      return nil, err
    }

    var reader io.ReadSeekCloser
    reader, err = s.NewReader(ctx, torrent, int(req.Index))
    if err != nil {
      return nil, err
    }
    defer reader.Close()


    _, err = reader.Seek(req.Off, 0)
    if err != nil {
      return nil, err
    }
    
    var buffer []byte
    buffer = make([]byte, req.Ln)

    _, err = reader.Read(buffer)
    if err != nil {
      return nil, err
    }

    return ReadAtRes{Buffer: buffer}, nil
  }
}

func bytes2torrent(data []byte) (service.Torrent, error) {
  var err error

  var buffer *bytes.Buffer
  buffer = bytes.NewBuffer(data)

  var decoder *gob.Decoder
  decoder = gob.NewDecoder(buffer)
  
  var torrent service.Torrent
  err = decoder.Decode(&torrent)
  if err != nil {
    return service.Torrent{}, err
  }

  return torrent, nil
}
