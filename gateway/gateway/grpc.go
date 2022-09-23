package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	chat "github.com/Iffigues/website/proto/linuxCommand"
)

func rigGrpc(a chat.Rig) (response *chat.Responses, err error) {
	var conn *grpc.ClientConn
	conn, errs := grpc.Dial("gopiko.fr:9000", grpc.WithInsecure())
	if errs != nil {
		return nil, errs
	}
	defer conn.Close()
	c := chat.NewRigServiceClient(conn)
	return c.GetRig(context.Background(), &a)
}


func fortuneGrpc(a chat.Fortune) (response *chat.Responses, err error) {
	e := chat.Fortune{}
	var conn *grpc.ClientConn
	conn, errs := grpc.Dial("gopiko.fr:9000", grpc.WithInsecure())
	if errs != nil {
		return nil, errs
	}
	defer conn.Close()
	c := chat.NewFortuneServiceClient(conn)
	return c.GetFortune(context.Background(), &e)
}

func fortuneFileGrpc() (response *chat.Responses, err error) {
	var conn *grpc.ClientConn
	e := &chat.Empty{}
	conn, errs := grpc.Dial("gopiko.fr:9000", grpc.WithInsecure())
	if errs != nil {
		return nil, errs
	}
	defer conn.Close()
	c := chat.NewFortuneServiceClient(conn)
	return c.GetFortuneFile(context.Background(), e)
}
