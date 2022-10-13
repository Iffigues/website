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

func cowGrpc(a chat.Cow) (response *chat.Responses, err error) {
	var conn *grpc.ClientConn
	println(a.Message)
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


func figletGrpc(a chat.Figlet) (response *chat.Responses, err error) {
	var conn *grpc.ClientConn
	conn, errs := grpc.Dial("gopiko.fr:9000", grpc.WithInsecure())
	if errs != nil {
		return nil, errs
	}
	defer conn.Close()
	c := chat.NewFigletServiceClient(conn)
	return c.GetFiglet(context.Background(), &a)
}

func figletFileGrpc() (response *chat.Responses, err error) {
	var conn *grpc.ClientConn
	e := &chat.Empty{}
	conn, errs := grpc.Dial("gopiko.fr:9000", grpc.WithInsecure())
	if errs != nil {
		return nil, errs
	}
	defer conn.Close()
	c := chat.NewFigletServiceClient(conn)
	return c.GetFigletFile(context.Background(), e)
}

func toiletGrpc(a chat.Toilet) (response *chat.Responses, err error) {
	var conn *grpc.ClientConn
	conn, errs := grpc.Dial("gopiko.fr:9000", grpc.WithInsecure())
	if errs != nil {
		return nil, errs
	}
	defer conn.Close()
	c := chat.NewToiletServiceClient(conn)
	return c.GetToilet(context.Background(), &a)
}

func toiletFFileGrpc() (response *chat.Responses, err error) {
	var conn *grpc.ClientConn
	e := &chat.Empty{}
	conn, errs := grpc.Dial("gopiko.fr:9000", grpc.WithInsecure())
	if errs != nil {
		return nil, errs
	}
	defer conn.Close()
	c := chat.NewToiletServiceClient(conn)
	return c.GetToiletFFile(context.Background(), e)
}

func toiletEFileGrpc() (response *chat.Responses, err error) {
	var conn *grpc.ClientConn
	e := &chat.Empty{}
	conn, errs := grpc.Dial("gopiko.fr:9000", grpc.WithInsecure())
	if errs != nil {
		return nil, errs
	}
	defer conn.Close()
	c := chat.NewToiletServiceClient(conn)
	return c.GetToiletEFile(context.Background(), e)
}


func toiletFFFileGrpc() (response *chat.Responses, err error) {
	var conn *grpc.ClientConn
	e := &chat.Empty{}
	conn, errs := grpc.Dial("gopiko.fr:9000", grpc.WithInsecure())
	if errs != nil {
		return nil, errs
	}
	defer conn.Close()
	c := chat.NewToiletServiceClient(conn)
	return c.GetToiletFFFile(context.Background(), e)
}

