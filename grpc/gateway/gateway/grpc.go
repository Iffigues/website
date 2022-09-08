package main

import (
	"fmt"
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	chat "github.com/Iffigues/website/grpc/protozoaire/linuxCommand"
)

func rigGrpc(a rig) (response *chat.Responses, err error) {
	fmt.Println(a)
	var conn *grpc.ClientConn
	conn, errs := grpc.Dial("gopiko.fr:9000", grpc.WithInsecure())
	if errs != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := chat.NewRigServiceClient(conn)
	response, err = c.GetRig(context.Background(), &chat.Rig{Nbr: a.Nbr, Man: a.Man, Woman: a.Woman})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	fmt.Println(response)
	return 
}
