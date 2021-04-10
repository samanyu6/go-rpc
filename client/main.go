package main

import (
	"context"
	"go-rpc/userpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9876", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client failed with error %v", err)
	}

	defer conn.Close()

	c := userpb.NewIdServiceClient(conn)

	singleUser, singleErr := c.GetUser(context.Background(), &userpb.GetId{Id: int32(1)})
	if singleErr != nil {
		log.Fatalf("Single user retrieve error %v", err)
	}
	multipleUser, mulErr := c.GetUserList(context.Background(), &userpb.GetIdList{Ids: []int32{1, 2, 3}})
	if mulErr != nil {
		log.Fatalf("Multiple user retrieve error %v", mulErr)
	}

	log.Printf("Response from single user %s", singleUser)
	log.Printf("Response from multiple user list %s", multipleUser)
}
