package main

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/proc"
	"iot_server/handler"
	"iot_server/log"

	_ "github.com/davyxu/cellnet/peer/tcp"
	_ "iot_server/proc/tcp"
)

const peerAddress = ":8086"

func main() {
	log.InitLog(0)
	queue := cellnet.NewEventQueue()
	peerIns := peer.NewGenericPeer("tcp.Acceptor", "server", peerAddress, queue)
	proc.BindProcessorHandler(peerIns, "tcp.iotltv", handler.JsonHandler)
	peerIns.Start()
	queue.StartLoop()
	queue.Wait()
}
