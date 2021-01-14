package handler

import (
	g "github.com/GramYang/gylog"
	"github.com/davyxu/cellnet"
	_ "github.com/davyxu/cellnet/codec/json"
)

func JsonHandler(ev cellnet.Event) {
	switch msg := ev.Message().(type) {
	case *cellnet.SessionAccepted:
		g.Debugln("server accepted")
	case *JsonEcho:
		g.Debugln("server recv", msg)
		ev.Session().Send(&JsonEcho{
			Msg: msg.Msg,
		})
	case *cellnet.SessionClosed:
		g.Debugln("session closed: ", ev.Session().ID())
	}
}
