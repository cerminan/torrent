package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/cerminan/torrent/config"
	"github.com/cerminan/torrent/endpoints"
	"github.com/cerminan/torrent/service"
	"github.com/cerminan/torrent/transport"
	"github.com/cerminan/torrent/transport/pb"
	gklog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
)

func main() {
  var err error

  var logger gklog.Logger
  logger = newLogger()

  var cfg config.Config
  cfg, err = config.DefaultConfig()
  if err != nil {
    level.Error(logger).Log("config", err)
  }

  err = cfg.LoadEnvar()
  if err != nil {
    level.Error(logger).Log("config", err)
  }
  
  var svc service.Service
  svc = service.NewService(logger)

  var ep endpoints.Endpoints
  ep = endpoints.MakeEndpoints(svc)
  
  var grpcServer pb.TorrentServer
  grpcServer = transport.NewGRPCServer(ep, logger)
    
  errs := make(chan error)
  go func() {
      c := make(chan os.Signal)
      signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
      errs <- fmt.Errorf("%s", <-c)
  }()

  grpcListener, err := net.Listen("tcp", cfg.Host)
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

