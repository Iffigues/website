package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	chat "github.com/Iffigues/website/proto/linuxCommand"
)

func rigGrpc(a rig) (response *chat.Responses, err error) {
	var conn *grpc.ClientConn
	conn, errs := grpc.Dial("gopiko.fr:9000", grpc.WithInsecure())
	if errs != nil {
		return nil, errs
	}
	defer conn.Close()
	c := chat.NewRigServiceClient(conn)
	return c.GetRig(context.Background(), &chat.Rig{Nbr: a.Nbr, Man: a.Man, Woman: a.Woman})
}


func fortuneGrpc(a fortune) (response *chat.Responses, err error) {
	e := chat.Fortune{}
	var conn *grpc.ClientConn
	conn, errs := grpc.Dial("gopiko.fr:9000", grpc.WithInsecure())
	if errs != nil {
		return nil, errs
	}
	defer conn.Close()
	c := chat.NewFortuneServiceClient(conn)
	e.M = a.M
	return c.GetFortune(context.Background(), &e)
}

func fortuneFileGrpc() (response *chat.Responses, err error) {
	var conn *grpc.ClientConn
	conn, errs := grpc.Dial("gopiko.fr:9000", grpc.WithInsecure())
	if errs != nil {
		return nil, errs
	}
	defer conn.Close()
	c := chat.NewFortuneServiceClient(conn)
	return c.GetFortuneFile(context.Background(), nil)
}
