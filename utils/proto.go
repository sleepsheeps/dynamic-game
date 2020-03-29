package utils

import (
	"dynamic-game/proto/helper"
	"github.com/golang/protobuf/proto"
)

func init() {
	if msgManager == nil {
		msgManager = new(MsgManager)
		msgManager.msgMap = make(map[helper.MsgType]RegisterFunc)
	}
}

type (
	RegisterFunc func(message proto.Message) error
)

type ProtoMsg struct {
	MsgType int32
	Msg     proto.Message
}

var (
	msgManager *MsgManager
)

func RegisterMsg(msgType helper.MsgType, f RegisterFunc) {
	msgManager.register(msgType, f)
}

func SendMsg(serverID string, msgType helper.MsgType, msg proto.Message) {
	m := &ProtoMsg{
		MsgType: int32(msgType),
		Msg:     msg,
	}
	mq.publish(serverID, m)
}

type MsgManager struct {
	msgMap map[helper.MsgType]RegisterFunc
}

func (m *MsgManager) register(msgType helper.MsgType, f RegisterFunc) {
	m.msgMap[msgType] = f
}

func (m *MsgManager) findFunc(msgType helper.MsgType) RegisterFunc {
	return m.msgMap[msgType]
}
