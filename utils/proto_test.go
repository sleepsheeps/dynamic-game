package utils

import (
	"dynamic-game/config"
	"dynamic-game/proto/helper"
	"dynamic-game/proto/test"
	"testing"
	"time"
)

func TestMarshal(t *testing.T) {
	config.Config.LoadConfig()
	StartMQ("0.0.0.0:4222")
	msg := &test.Student{}
	SendMsg(config.Config.ServerID, helper.MsgType_T_Student, msg)
	time.Sleep(2 * time.Second)
}
