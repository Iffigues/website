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
	var conn *grpc.ClientConn
	conn, errs := grpc.Dial("gopiko.fr:9000", grpc.WithInsecure())
	if errs != nil {
		return nil, errs
	}
	defer conn.Close()
	c := chat.NewFortuneServiceClient(conn)
	return c.GetFortune(context.Background(), &a)
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

func CowGrpc(a chat.Cow) (response *chat.Responses, err error) {
	var conn *grpc.ClientConn
	conn, errs := grpc.Dial("gopiko.fr:9000", grpc.WithInsecure())
	if errs != nil {
		return nil, errs
	}
	defer conn.Close()
	c := chat.NewCowServiceClient(conn)
	return c.GetCow(context.Background(), &a)
}

func cowFileGrpc() (response *chat.Responses, err error) {
	var conn *grpc.ClientConn
	e := &chat.Empty{}
	conn, errs := grpc.Dial("gopiko.fr:9000", grpc.WithInsecure())
	if errs != nil {
		return nil, errs
	}
	defer conn.Close()
	c := chat.NewCowServiceClient(conn)
	return c.GetCowFile(context.Background(), e)
}
