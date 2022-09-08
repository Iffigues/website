package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	chat "github.com/Iffigues/website/grpc/protozoaire/linuxCommand"
)

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serve := chat.Server{}
	grpcServer := grpc.NewServer()
	chat.RegisterRigServiceServer(grpcServer, &serve)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
