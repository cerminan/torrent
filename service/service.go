package service

import (
	"context"
	"errors"
	"io"
	"time"

	anacrolixLog "github.com/anacrolix/log"
	anacrolixTorrent "github.com/anacrolix/torrent"
	anacrolixMetainfo "github.com/anacrolix/torrent/metainfo"
	anacrolixStorage "github.com/anacrolix/torrent/storage"
	"github.com/cerminan/torrent/service/storage"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
  logger log.Logger
  client *anacrolixTorrent.Client
}

type Service interface {
  Fetch(ctx context.Context, url string) (Torrent, error)
  NewReader(ctx context.Context, torrent Torrent, index int) (io.ReadSeekCloser, error)
}

type Torrent struct {
  Hash string
  Files []File
}


func NewService(logger log.Logger) (Service) {
  var err error

  var basedir string
  basedir = "pieces"
  var lifetime time.Duration
  lifetime = 1 * time.Minute

  var config *anacrolixTorrent.ClientConfig
  config = anacrolixTorrent.NewDefaultClientConfig()
  config.Logger = anacrolixLog.Discard
  config.DefaultStorage = anacrolixStorage.NewResourcePieces(storage.NewProvider(basedir, lifetime))
  
  var client *anacrolixTorrent.Client
  client, err = anacrolixTorrent.NewClient(config)
  if err != nil {
    level.Error(logger).Log("init", err.Error())
  }

  return &service{
    logger: logger,
    client: client,
  }
}

type File struct {
  Name string
  Length int64
}

func (s *service) Fetch(ctx context.Context, url string) (Torrent, error) {
  var err error
  var torrent Torrent

  var torrentInstance *anacrolixTorrent.Torrent

  torrentInstance, err = s.client.AddMagnet(url)
  if err != nil {
    return torrent, err
  }
  
  <-torrentInstance.GotInfo()
  torrentInstance.DisallowDataUpload()

  torrent = Torrent{
    Hash: torrentInstance.InfoHash().HexString(),
  }

  for _, fileInstance := range torrentInstance.Files() {
    torrent.Files = append(torrent.Files, File{
      Name: fileInstance.DisplayPath(),
      Length: fileInstance.Length(),
    })
  }

  return torrent, nil
}

func (s *service) NewReader(ctx context.Context, torrent Torrent, index int) (io.ReadSeekCloser, error) {
  var hash anacrolixMetainfo.Hash
  hash = anacrolixMetainfo.NewHashFromHex(torrent.Hash)
  var torrentInstance *anacrolixTorrent.Torrent
  var exists bool
  torrentInstance, exists = s.client.Torrent(hash)
  if !exists {
    return nil, errors.New("Torrent not found.")
  }
  var fileInstances []*anacrolixTorrent.File
  fileInstances = torrentInstance.Files()
  if len(fileInstances) < index {
    return nil, errors.New("File not found.")
  }

  var fileInstance *anacrolixTorrent.File
  fileInstance = fileInstances[index]

  return fileInstance.NewReader(), nil
}
