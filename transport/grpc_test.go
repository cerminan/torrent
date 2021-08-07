package transport

import (
	"context"
  "encoding/base64"
	"testing"
	"time"

	"github.com/cerminan/torrent/service"
	"google.golang.org/grpc"
)

const host string = ":50051"
const magnet string = "magnet:?xt=urn:btih:3WBFL3G4PSSV7MF37AJSHWDQMLNR63I4&dn=Big%20Buck%20Bunny&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969"

func newClient() (*grpc.ClientConn, error){
  var err error
  var conn *grpc.ClientConn
  conn, err = grpc.Dial(host, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
  if err != nil {
    return nil, err
  }

  return conn, nil
}

func TestFiles(t *testing.T) {
  var trueFiles []service.File 
  trueFiles = []service.File{
    {
      TorrentHash: "dd8255ecdc7ca55fb0bbf81323d87062db1f6d1c",
      Name: "Big Buck Bunny.en.srt",
      Length: 140,
    },
    {
      TorrentHash: "dd8255ecdc7ca55fb0bbf81323d87062db1f6d1c",
      Name: "Big Buck Bunny.mp4",
      Length: 276134947,
    },
    {
      TorrentHash: "dd8255ecdc7ca55fb0bbf81323d87062db1f6d1c",
      Name: "poster.jpg",
      Length: 310380,
    },
  }

  var err error

  var conn *grpc.ClientConn
  conn, err = newClient()
  if err != nil {
    t.Fatal(err)
  }
  defer conn.Close()

  var svc service.Service
  svc = NewGRPCClient(conn)

  var files []service.File
  files, err = svc.Files(context.Background(), magnet)
  if err != nil {
    t.Error(err)
  }

  for index := range files {
    if files[index] != trueFiles[index] {
      t.Error("Cannot fetch file list")
    }
  }
}

func TestReadAt(t *testing.T) {
  var file service.File
  file = service.File{
    TorrentHash: "dd8255ecdc7ca55fb0bbf81323d87062db1f6d1c",
    Name: "Big Buck Bunny.en.srt",
    Length: 140,
  }

  const off int64 = 10
  const ln int64 = 10

  var trueBuffer []byte
  trueBuffer = []byte{48, 50, 44, 48, 48, 48, 32, 45, 45, 62}

  var trueBufferStr string
  trueBufferStr = base64.StdEncoding.EncodeToString(trueBuffer)

  var err error

  var conn *grpc.ClientConn
  conn, err = newClient()
  if err != nil {
    t.Fatal(err)
  }
  defer conn.Close()

  var svc service.Service
  svc = NewGRPCClient(conn)

  var buffer []byte
  buffer, err = svc.ReadAt(context.Background(), file, off, ln)
  if err != nil {
    t.Error(err)
  }

  var bufferStr string
  bufferStr = base64.StdEncoding.EncodeToString(buffer)

  if bufferStr != trueBufferStr {
    t.Error("ReadAt does not match")
  }
}
