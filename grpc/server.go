package grpc

import (
	"net"
	"google.golang.org/grpc"
	"context"
	"fmt"
	"io"
	"time"
)

const (
	port = ":10086"
)

type data int

func openServer() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	var x data
	RegisterControllerServer(s, &x)
	if err = s.Serve(l); err != nil {
		panic(err)
	}
}

func (d *data) SayHello(ctx context.Context, req *HelloRequest) (res *HelloResponse, err error) {
	reqMsg := req.GetName()
	fmt.Println("server: ", reqMsg)
	res = &HelloResponse{
		Message: "hello " + reqMsg,
	}
	return res, err
}

func (d *data) SayStream(stream Controller_SayStreamServer) error {

	//开启接收
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			//fmt.Println(err)
			break
		}
		fmt.Println("reqeustid:", req.Requestid, "requestdata:", req.Data)
		time.Sleep(time.Second * 5) //假设处理5s
		//收到后返回消息
		res := &ResponseStreamObj{
			Responseid: req.Requestid,
			Data:       "response data",
		}
		stream.Send(res)
	}

	return nil
}
