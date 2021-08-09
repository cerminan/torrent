package main

import (
	"net"
	"os"

	"github.com/cerminan/libs/exit"
	"github.com/cerminan/torrent/config"
	"github.com/cerminan/torrent/endpoints"
	"github.com/cerminan/torrent/service"
	transportGRPC "github.com/cerminan/torrent/transport/grpc"
	"github.com/cerminan/torrent/transport/grpc/pb"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
)

func main() {
  var cerr chan error
  cerr = exit.ExitSignal()

  var logger log.Logger
  {
    logger = log.NewLogfmtLogger(os.Stdout)
    logger = log.With(logger, "ts", log.DefaultTimestampUTC)
    logger = log.With(logger, "caller", log.DefaultCaller)
  }
  
  go func(){
    var err error

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
    grpcServer = transportGRPC.NewServer(ep)

    var listener net.Listener
    listener, err = net.Listen("tcp", cfg.Host)
    if err != nil {
        level.Error(logger).Log("listen", err)
        os.Exit(1)
    }

    var baseServer *grpc.Server
    baseServer = grpc.NewServer()
    pb.RegisterTorrentServer(baseServer, grpcServer)
    level.Info(logger).Log("msg", "Server started successfully")
    level.Info(logger).Log("listen", cfg.Host)
    err = baseServer.Serve(listener)
    if err != nil {
      level.Error(logger).Log("grpc", err)
    }
  }()

  level.Error(logger).Log("exit", <-cerr)
} 

