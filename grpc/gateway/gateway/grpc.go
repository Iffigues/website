package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	chat "github.com/Iffigues/mywebsite/grpc/protozoaire/linuxCommand"
)

func test(a rig) (response *chat.RigResponses, err error) {
	var conn *grpc.ClientConn
	conn, errs := grpc.Dial(":9000", grpc.WithInsecure())
	if errs != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := chat.NewRigServiceClient(conn)
	response, err = c.GetRig(context.Background(), &chat.Rig{Nbr: a.Nbr, Man: a.Man, Woman: a.Woman})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	return 
}
