package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/cerminan/torrent/endpoints"
	"github.com/cerminan/torrent/pb"
	"github.com/cerminan/torrent/service"
	"github.com/cerminan/torrent/transport"
	gklog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
)

func main() {
  var err error
  
  var logger gklog.Logger
  logger = gklog.NewLogfmtLogger(os.Stdout)
  logger = gklog.With(logger, "ts", gklog.DefaultTimestampUTC)
  logger = gklog.With(logger, "caller", gklog.DefaultCaller)

  var aservice service.Service
  aservice = service.NewService(logger)
  
  var grpcServer pb.TorrentServer
  grpcServer = transport.NewGRPCServer(endpoints.MakeEndpoints(aservice), logger)
    

  errs := make(chan error)
  go func() {
      c := make(chan os.Signal)
      signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
      errs <- fmt.Errorf("%s", <-c)
  }()

  grpcListener, err := net.Listen("tcp", ":50051")
  if err != nil {
      logger.Log("during", "Listen", "err", err)
      os.Exit(1)
  }

  go func() {
      baseServer := grpc.NewServer()
      pb.RegisterTorrentServer(baseServer, grpcServer)
      level.Info(logger).Log("msg", "Server started successfully")
      baseServer.Serve(grpcListener)
  }()

  level.Error(logger).Log("exit", <-errs)
} 

