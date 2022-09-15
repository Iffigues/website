package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	chat "github.com/Iffigues/website/grpc/protozoaire/linuxCommand"
)

func rigGrpc(a rig) (response *chat.Responses, err error) {
	var conn *grpc.ClientConn
	conn, errs := grpc.Dial("gopiko.fr:9000", grpc.WithInsecure())
	if errs != nil {
		return nil, errs
	}
	defer conn.Close()
	c := chat.NewRigServiceClient(conn)
	print("nbr", a.Nbr)
	return c.GetRig(context.Background(), &chat.Rig{Nbr: a.Nbr, Man: a.Man, Woman: a.Woman})
}
