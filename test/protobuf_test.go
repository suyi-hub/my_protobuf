package test

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"studyProtobuf/github.com/suyi-hub/my_protobuf"
	"testing"
)

func TestProtobuf(t *testing.T) {
	//定义消息体
	mypanda := &my_protobuf.PandaRequest{
		Name:   "panda",
		High:   167,
		Weight: []int32{1, 2, 3},
	}
	//	proto序列化
	data, err := proto.Marshal(mypanda)
	if err != nil {
		fmt.Printf("编码失败 %v", err)
	}
	fmt.Println(data)
	//准备接收对象
	newmypanda := &my_protobuf.PandaRequest{}
	//将序列化的data传入
	err = proto.Unmarshal(data, newmypanda)
	if err != nil {
		fmt.Printf("解码失败 %v", err)
	}
	println(fmt.Println(newmypanda))

}
