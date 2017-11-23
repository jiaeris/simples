package grpc

import (
	"google.golang.org/grpc"
	"context"
	"fmt"
	"sync"
	"strconv"
)

const (
	address = "localhost:10086"
	times   = 10
)

func openClient() {

	//模拟100个客户端并发请求
	i := 0
	wg := sync.WaitGroup{}
	for i < times {
		wg.Add(1)
		go func(i int) {
			client, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				fmt.Println(err)
			}
			defer client.Close()
			cc := NewControllerClient(client)
			res, err := cc.SayHello(context.Background(), &HelloRequest{Name: "Yunga" + strconv.Itoa(i)}, grpc.FailFast(true))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("clent1: ", res.GetMessage())

			//开启接收线程
			streamClient, err := cc.SayStream(context.Background(), grpc.FailFast(true))
			resCh := make(chan string)
			go func() {
				for {
					res, err := streamClient.Recv()
					if err != nil {
						//fmt.Println(err)
						break
					}
					fmt.Println("responseid:", res.Responseid, "responsedata:", res.Data)
					resCh <- "ok"
				}
			}()
			//开启请求
			req := &RequestStreamObj{
				Requestid: int32(i),
				Data:      "request data",
			}
			streamClient.Send(req)
			<-resCh
			streamClient.CloseSend()
			wg.Done()
		}(i)
		i++
	}
	wg.Wait()
	fmt.Println("clients request finshed.")
}
