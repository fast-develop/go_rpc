package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	log "github.com/golang/glog"
	"google.golang.org/grpc/reflection"

	"recall_title2vec/config"
	"recall_title2vec/proto/recall_title2vec"
	"recall_title2vec/server"
)

var (
	ConfPath = flag.String("config", "../conf/config.xml", "config path")
)

func main() {
	flag.Parse()
	log.Infof("start recall_title2vec server")
    fmt.Println("start recall_title2vec server")

	//config init
	err := config.Conf.Init(ConfPath)
	if err != nil {
		log.Fatalf("config init error")
		return
	}
    fmt.Println("config init sucess")

	//server init
	srv, err := server.NewServerImpl(config.Conf)
	if err != nil {
		log.Errorf("server init error")
		return
	}
    fmt.Println("server init sucess")

	//register server
	grpcSrv := grpc.NewServer()
	recall_title2vec.RegisterRecallTitle2VecServer(grpcSrv, srv)

	// Register reflection service on gRPC server.
	reflection.Register(grpcSrv)
	if err := grpcSrv.Serve(srv.Listen); err != nil {
		log.Errorf("failed to serve: %v", err)
		return
	}

	log.Infof("server start sucess")
}
