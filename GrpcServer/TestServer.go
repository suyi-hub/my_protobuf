package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	pb "studyProtobuf/github.com/suyi-hub/my_protobuf"
)

type server struct{}

// rpc 服务的函数名 （传入参数）返回（返回参数）
// 传入参数和返回参数都是消息
func (This *server) SayHello(c context.Context, w *pb.HelloRequest) (r *pb.HelloRespond, err error) {
	name := w.GetName() + ",你好呀我是服务器SayHello"
	r = &pb.HelloRespond{Msg: name}
	return r, nil
}
func (This *server) SayName(c context.Context, w *pb.NameRequest) (r *pb.NameRespond, err error) {
	name := w.GetName() + ",你好呀我是服务器SayName"
	r = &pb.NameRespond{Msg: name}
	return r, nil
	return r, err
}
func main() {
	//开启网络监听
	ln, err := net.Listen("tcp", "127.0.0.1:8081")

	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer ln.Close()
	//创建grpc服务
	grpcServer := grpc.NewServer()
	//注册服务
	pb.RegisterHelloServiceServer(grpcServer, &server{})
	//绑定网络
	fmt.Println("start grpc server success")

	err = grpcServer.Serve(ln)

	if err != nil {
		fmt.Println("grpcServer.Serve err:", err)
	}

}
