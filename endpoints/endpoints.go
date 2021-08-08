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
  rawRes, err = e.FilesEndpoint(ctx, FilesRequst{Magnet: magnet})
  if err != nil {
    return nil, err
  }
  
  var res = rawRes.(FilesResponse)

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
  rawRes, err = e.ReadAtEndpoint(ctx, ReadAtRequest{File: File(file), Off: off, Ln: ln})
  if err != nil {
    return nil, err
  }

  var res ReadAtResponse
  res = rawRes.(ReadAtResponse)
  
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

