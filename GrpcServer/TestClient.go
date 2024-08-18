package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "studyProtobuf/github.com/suyi-hub/my_protobuf"
)

func main() {
	conn, err := grpc.NewClient("127.0.0.1:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//conn, err := grpc.Dial("127.0.0.1:8099", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("NewClient error:%v\n", err)
	}
	defer conn.Close()
	client := pb.NewHelloServiceClient(conn)
	res, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "阿毅"})
	if err != nil {
		fmt.Printf("SayHello error:%v\n", err)
	}
	fmt.Printf("SayHello result:%v\n", res.GetMsg())
}
