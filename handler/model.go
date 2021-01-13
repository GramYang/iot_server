package handler

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	"reflect"
)

type JsonEcho struct {
	Msg   string
}

// 将消息注册到系统
func init() {
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("json"),
		Type:  reflect.TypeOf((*JsonEcho)(nil)).Elem(),
		ID:    1,
	})
}