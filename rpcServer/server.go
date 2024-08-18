package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
)

func Panda(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}

type Suyi string

func (this *Suyi) Name(argType string, replyType *string) error {
	fmt.Printf("对端发送一个%v\n：", argType)
	//修改内容
	*replyType = argType

	return nil
}

func main() {
	//router
	//http.HandleFunc("/panda", Panda)

	yi := new(Suyi)         //类实例化对象
	err := rpc.Register(yi) //注册对象
	if err != nil {
		fmt.Println(err)
	}
	rpc.HandleHTTP()
	//listen net
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net err")
	}
	err = http.Serve(listen, nil)
	if err != nil {
		fmt.Println("http err")
	}

}
