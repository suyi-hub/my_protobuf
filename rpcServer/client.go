package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	Client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var pb string
	err = Client.Call("Suyi.Name", "阿毅", &pb)
	if err != nil {
		log.Fatal("Call suyi.Name:", err)

	}
	fmt.Printf("接收到对端传回的%v", pb)
}
