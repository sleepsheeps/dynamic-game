package utils

import (
	"dynamic-game/proto/helper"
	"dynamic-game/proto/test"
	"testing"
)

func TestMarshal(t *testing.T) {
	StartMQ("0.0.0.0:4222")
	msg := &test.Student{}
	SendMsg("test", helper.MsgType_T_Student, msg)
}
