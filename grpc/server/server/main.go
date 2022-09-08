package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	chat "github.com/Iffigues/website/grpc/protozoaire/linuxCommand"
	//"github.com/Iffigues/mywebsite/grpc/server/server/lw"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{}
	grpcServer := grpc.NewServer()
	chat.RegisterRigServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
