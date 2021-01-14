package main

import (
	"fmt"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/proc"
	"iot_server/handler"

	_ "github.com/davyxu/cellnet/peer/tcp"
	_ "github.com/davyxu/cellnet/proc/tcp"
)

//const peerAddress = "106.54.87.204:8086"

const peerAddress = "127.0.0.1:8086"

func main() {
	done := make(chan struct{})
	queue := cellnet.NewEventQueue()
	peerIns := peer.NewGenericPeer("tcp.Connector", "client", peerAddress, queue)
	proc.BindProcessorHandler(peerIns, "tcp.ltv", func(ev cellnet.Event) {
		switch msg := ev.Message().(type) {
		//连接成功后发送一个消息
		case *cellnet.SessionConnected:
			fmt.Println("client connected")
			ev.Session().Send(&handler.JsonEcho{
				Msg: "hello",
			})
		//收到响应后就关闭
		case *handler.JsonEcho:
			fmt.Printf("client recv %+v\n", msg)
			done <- struct{}{}
		case *cellnet.SessionClosed:
			fmt.Println("client closed")
		}
	})
	peerIns.Start()
	queue.StartLoop()
	<-done
}
