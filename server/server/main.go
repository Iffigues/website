package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	chat "github.com/Iffigues/website/proto/linuxCommand"
)

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serve := chat.Server{}
	grpcServer := grpc.NewServer()
	chat.RegisterRigServiceServer(grpcServer, &serve)
	chat.RegisterFortuneServiceServer(grpcServer, &serve)
	chat.RegisterCowServiceServer(grpcServer, &serve)
	chat.RegisterFigletServiceServer(grpcServer, &serve)
	chat.RegisterToiletServiceServer(grpcServer, &serve)
	hat.RegisterXkcdpassServiceServer(grpcServer, &serve)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
