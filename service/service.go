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
  Files(ctx context.Context, url string) ([]File, error)
  ReadAt(ctx context.Context, file File, off int64, ln int64) ([]byte, error)
}

type File struct {
  TorrentHash string
  Name string
  Length int64
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

func (s *service) Files(ctx context.Context, url string) ([]File, error) {
  var err error

  var torrentInstance *anacrolixTorrent.Torrent
  torrentInstance, err = s.client.AddMagnet(url)
  if err != nil {
    return nil, err
  }
  
  <-torrentInstance.GotInfo()
  torrentInstance.DisallowDataUpload()

  var hash string
  hash = torrentInstance.InfoHash().HexString()

  var files []File
  for _, fileInstance := range torrentInstance.Files() {
    files = append(files, File{
      TorrentHash: hash,
      Name: fileInstance.DisplayPath(),
      Length: fileInstance.Length(),
    })
  }

  return files, nil
}

func (s *service) ReadAt(ctx context.Context, file File, off int64, ln int64) ([]byte, error) {
  var hash anacrolixMetainfo.Hash
  hash = anacrolixMetainfo.NewHashFromHex(file.TorrentHash)

  var torrentInstance *anacrolixTorrent.Torrent
  var exists bool
  torrentInstance, exists = s.client.Torrent(hash)
  if !exists {
    return nil, errors.New("Torrent not found.")
  }

  var fileInstances []*anacrolixTorrent.File
  fileInstances = torrentInstance.Files()

  var fileInstance *anacrolixTorrent.File
  for _, candidateFileInstance := range fileInstances {
    if candidateFileInstance.DisplayPath() == file.Name {
      fileInstance = candidateFileInstance
      break
    }
  }
  
  if fileInstance == nil {
    return nil, errors.New("File not found.")
  }

  var reader io.ReadSeekCloser
  reader = fileInstance.NewReader()
  defer reader.Close()

  var err error
  _, err = reader.Seek(off, 0)
  if err != nil {
    return nil, err
  }

  var buffer []byte
  buffer = make([]byte, ln)
  
  _, err = reader.Read(buffer)
  if err != nil {
    return nil, err
  }

  return buffer, nil
}
